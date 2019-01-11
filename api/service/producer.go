/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : producer.go
#   Created       : 2019/1/9 15:49
#   Last Modified : 2019/1/9 15:49
#   Describe      :
#
# ====================================================*/
package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
	"uuabc.com/sendmsg/api/storer/db"
	"uuabc.com/sendmsg/api/storer/mq"
	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/pkg/retry/backoff"
)

var ProducerImpl = producerImpl{}

type producerImpl struct {
}

func (p producerImpl) Produce(ctx context.Context, meta Meta) error {
	b, err := meta.Marshal()
	if err != nil {
		return err
	}
	err = p.sendToMq(ctx, meta.TypeName(), b, meta.Delay())
	if err != nil {
		logrus.Errorf("消息 %s 插入消息队列失败，error: %v\n", string(b), err)
		return err
	}
	logrus.Infof("消息 %s 插入消息队列成功,正在等待发送,开始准备插入数据库", string(b))

	// 异步将数据插入数据库
	go func() {
		var err error
		var count int
		// 重试机制，最多试两次
		retryFunc := func() error {
			if count > 2 {
				logrus.WithFields(logrus.Fields{
					"type": meta.TypeName(),
					"id":   meta.GetId(),
					"data": string(b),
				}).Errorf("数据插入数据库失败，请手动尝试,error: %v", err)
				return nil
			}
			count++
			err = p.saveToDB(context.Background(), meta)
			if err == nil {
				logrus.Infof("消息 %s 插入数据库成功", string(b))
			}
			return err
		}
		back := backoff.NewExponentialBackOff()
		back.InitialInterval = time.Millisecond * 100
		back.Multiplier = 1.2
		back.MaxInterval = time.Millisecond * 300
		back.MaxElapsedTime = time.Second * 5
		err = backoff.Retry(retryFunc, back)
		if err != nil {
			logrus.Errorf("unexpected error: %s", err.Error())
		}
	}()

	return nil
}

func (producerImpl) sendToMq(ctx context.Context, n string, b []byte, d int64) (err error) {
	switch n {
	case weixin:
		err = mq.ProduceWeChat(ctx, b, d)
	case sms:
		err = mq.ProduceSms(ctx, b, d)
	case email:
		err = mq.ProduceEmail(ctx, b, d)
	default:
		err = errors.ErrMsgTypeNotFound
	}
	return
}

func (producerImpl) saveToDB(ctx context.Context, m Meta) (err error) {
	switch m.(type) {
	case *meta.SmsProducer:
		v := m.(*meta.SmsProducer)
		err = db.InsertSmss(ctx, v)
	case *meta.WeChatProducer:
		v := m.(*meta.WeChatProducer)
		err = db.InsertWechats(ctx, v)
	case *meta.EmailProducer:
		v := m.(*meta.EmailProducer)
		err = db.InsertEmails(ctx, v)
	default:
		err = errors.ErrMsgTypeNotFound
	}
	return
}
