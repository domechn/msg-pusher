/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : sms.go
#   Created       : 2019/1/16 10:27
#   Last Modified : 2019/1/16 10:27
#   Describe      :
#
# ====================================================*/
package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/storer/db"
	"uuabc.com/sendmsg/storer/mq"
)

type smsServiceImpl struct {
}

func NewSmsServiceImpl() smsServiceImpl {
	return smsServiceImpl{}
}

func (s smsServiceImpl) Produce(ctx context.Context, m Meta) (string, error) {
	var templ string
	var args map[string]string
	var err error
	if templ, args, err = checkTemplateAndArguments(m.GetTemplate(), m.GetArguments()); err != nil {
		return "", err
	}
	content := getContent(args, templ)
	ttl := m.Delay()
	err = s.produce(ctx, m.(*meta.SmsProducer), content, ttl)
	return m.GetId(), err
}

func (smsServiceImpl) produce(ctx context.Context, p *meta.SmsProducer, content string, ttl int64) error {
	dbSms := &meta.DbSms{
		Id:          p.Id,
		Platform:    p.Platform,
		PlatformKey: p.PlatformKey,
		Content:     content,
		Mobile:      p.Mobile,
		Template:    p.Template,
		Arguments:   p.Arguments,
		SendTime:    p.SendTime,
		Server:      p.Server,
		Type:        p.Type,
	}
	tx, err := db.SmsInsert(ctx, dbSms)
	if err != nil {
		return err
	}
	id := dbSms.Id
	err = mq.SmsProduce(ctx, []byte(id), ttl)
	if err != nil {
		logrus.WithField("type", "Sms").Errorf("消息 %s 插入消息队列失败，正在回滚。。。，error: %v\n", id, err)
		db.RollBack(tx)
		return err
	}
	logrus.WithField("type", "Sms").Infof("消息 %s 插入消息队列成功,正在等待发送,开始提交到数据库", id)
	err = db.Commit(tx)
	if err != nil {
		return err
	}
	go updateCache(context.Background(), id, dbSms)
	logrus.WithField("type", "Sms").Infof("消息 %s 插入数据库成功", id)
	return nil
}

func (s smsServiceImpl) Detail(ctx context.Context, id string) (Marshaler, error) {
	return s.detail(ctx, id)
}

// DetailByPhonePage 直接数据库中取,不走缓存
func (s smsServiceImpl) DetailByPhonePage(ctx context.Context, mobile string, page int) ([]byte, error) {
	return nil, nil
}

func (smsServiceImpl) detail(ctx context.Context, id string) (Marshaler, error) {
	res := &meta.DbSms{}
	return res, detail(ctx, id, res, func(ctx2 context.Context, id string) (Marshaler, error) {
		return db.SmsDetailByID(ctx2, id)
	})
}

func (s smsServiceImpl) Cancel(ctx context.Context, id string) error {
	return s.cancel(ctx, id)
}

func (smsServiceImpl) cancel(ctx context.Context, id string) error {
	return cancel(ctx, id, func(i context.Context, s string) (*sqlx.Tx, error) {
		return db.SmsCancelMsgByID(i, s)
	}, &meta.DbSms{})
}

func (s smsServiceImpl) Edit(ctx context.Context, m Meta) error {
	m.Transfer(false)
	v := m.(*meta.SmsProducer)
	return s.edit(
		ctx,
		m,
		&meta.DbSms{
			Id:        v.Id,
			Arguments: v.Arguments,
			SendTime:  v.SendTime,
			Mobile:    v.Mobile,
		})
}

func (s smsServiceImpl) edit(ctx context.Context, m Meta, e *meta.DbSms) error {
	em := &meta.DbSms{}
	if err := checkStatus(m.GetId(), em); err != nil {
		return err
	}

	// 修改数据
	em.Arguments = e.Arguments
	em.SendTime = e.SendTime
	if e.Mobile != "" {
		em.Mobile = e.Mobile
	}

	tx, err := db.SmsEdit(ctx, e)
	if err != nil {
		db.RollBack(tx)
		logrus.WithField("type", "Sms").Errorf("edit修改数据库失败,error: %v", err)
		return err
	}

	err = edit(ctx, em, m, mq.SmsProduce)
	if err != nil {
		db.RollBack(tx)
		logrus.WithField("type", "Sms").Errorf("edit更新mq失败，正在回滚,error: %v", err)
		return err
	}

	return db.Commit(tx)
}
