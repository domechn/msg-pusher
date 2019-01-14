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
	"uuabc.com/sendmsg/api/model"
	"uuabc.com/sendmsg/api/storer/cache"
	"uuabc.com/sendmsg/api/storer/db"
	"uuabc.com/sendmsg/api/storer/mq"
	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/pkg/pb/meta"
)

var ProducerImpl = producerImpl{}

type producerImpl struct{}

func (p producerImpl) Produce(ctx context.Context, m Meta) (string, error) {
	err := p.produce(ctx, m)
	if err == nil {
		return m.GetId(), nil
	}
	// 不是项目用的error
	if _, ok := err.(*errors.Error); !ok {
		err = errors.NewError(
			10000000,
			err.Error(),
		)
	}
	return "", err
}

func (p producerImpl) produce(ctx context.Context, m Meta) error {
	ttl := m.Delay()
	switch m.(type) {
	case *meta.WeChatProducer:
		return p.produceWechat(
			ctx,
			m.(*meta.WeChatProducer),
			func(p *meta.WeChatProducer) *model.DbWeChat {
				return &model.DbWeChat{
					ID:       p.Id,
					Platform: p.Platform,
					Touser:   p.Touser,
					Type:     p.Type,
					Template: p.TemplateID,
					URL:      p.Url,
					Content:  p.Data,
					SendTime: p.SendTime,
				}
			},
			ttl,
		)
	case *meta.EmailProducer:
		return p.produceEmail(
			ctx,
			m.(*meta.EmailProducer),
			func(p *meta.EmailProducer) *model.DbEmail {
				return &model.DbEmail{
					ID:          p.Id,
					Platform:    p.Platform,
					PlatformKey: p.PlatformKey,
					Title:       p.Content,
					Destination: p.Destination,
					Type:        p.Type,
					Template:    p.Template,
					Arguments:   p.Arguments,
					Server:      p.Server,
					SendTime:    p.SendTime,
				}
			},
			ttl,
		)
	case *meta.SmsProducer:
		return p.produceSms(
			ctx,
			m.(*meta.SmsProducer),
			func(p *meta.SmsProducer) *model.DbSms {
				return &model.DbSms{
					ID:        p.Id,
					Platform:  p.Platform,
					Content:   p.Content,
					Mobile:    p.Mobile,
					Template:  p.Template,
					Arguments: p.Arguments,
					SendTime:  p.SendTime,
					Server:    p.Server,
					Type:      p.Type,
				}
			},
			ttl)
	}
	return errors.ErrMsgTypeNotFound
}

func (producerImpl) produceSms(ctx context.Context, sms *meta.SmsProducer, change func(*meta.SmsProducer) *model.DbSms, ttl int64) error {
	b, err := sms.Marshal()
	if err != nil {
		return err
	}
	dbSms := change(sms)
	tx, err := db.SmsInsert(ctx, dbSms)
	if err != nil {
		return err
	}
	err = mq.ProduceSms(ctx, b, ttl)
	if err != nil {
		logrus.Errorf("消息 %s 插入消息队列失败，正在回滚。。。，error: %v\n", string(b), err)
		rollback(tx)
		return err
	}
	logrus.Infof("消息 %s 插入消息队列成功,正在等待发送,开始提交到数据库", string(b))
	err = commit(tx)
	if err != nil {
		return err
	}
	go func() {
		byt, gErr := dbSms.Marshal()
		if gErr != nil {
			logrus.Errorf("set cache go func() error: %v", gErr)
			return
		}
		// 插入redis
		cache.PutBaseCache(sms.GetId(), byt)
	}()
	logrus.Infof("消息 %s 插入数据库成功", string(b))
	return nil
}

func (producerImpl) produceEmail(ctx context.Context, email *meta.EmailProducer, change func(*meta.EmailProducer) *model.DbEmail, ttl int64) error {
	b, err := email.Marshal()
	if err != nil {
		return err
	}
	dbEmail := change(email)
	tx, err := db.InsertEmails(ctx, dbEmail)
	if err != nil {
		return err
	}
	err = mq.ProduceEmail(ctx, b, ttl)
	if err != nil {
		rollback(tx)
		return err
	}
	err = commit(tx)
	if err != nil {
		return err
	}
	go func() {
		byt, gErr := dbEmail.Marshal()
		if gErr != nil {
			logrus.Errorf("set cache go func() error: %v", gErr)
			return
		}
		// 插入redis
		cache.PutBaseCache(email.GetId(), byt)
	}()
	return nil
}

func (producerImpl) produceWechat(ctx context.Context, wechat *meta.WeChatProducer, change func(*meta.WeChatProducer) *model.DbWeChat, ttl int64) error {
	b, err := wechat.Marshal()
	if err != nil {
		return err
	}
	dbWechat := change(wechat)
	tx, err := db.WeChatInsert(ctx, dbWechat)
	if err != nil {
		return err
	}
	err = mq.ProduceWeChat(ctx, b, ttl)
	if err != nil {
		rollback(tx)
		return err
	}
	err = commit(tx)
	if err != nil {
		return err
	}
	go func() {
		byt, gErr := dbWechat.Marshal()
		if gErr != nil {
			logrus.Errorf("set cache go func() error: %v", gErr)
			return
		}
		// 插入redis
		cache.PutBaseCache(wechat.GetId(), byt)
	}()
	return nil
}
