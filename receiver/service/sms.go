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
	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/storer/cache"
	"uuabc.com/sendmsg/storer/db"
	"uuabc.com/sendmsg/storer/mq"
)

type smsServiceImpl struct {
}

// NewSmsServiceImpl 初始化消息service
func NewSmsServiceImpl() smsServiceImpl {
	return smsServiceImpl{}
}

func (s smsServiceImpl) Produce(ctx context.Context, m Meta) (string, error) {
	var templ string
	var args map[string]string
	var err error
	mobile := m.GetSendTo()
	if err := s.checkSendRate(ctx, mobile); err != nil {
		return "", err
	}
	if templ, args, err = checkTemplateAndArguments(ctx, m.GetTemplate(), m.GetArguments()); err != nil {
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
	return produce(ctx,
		p,
		dbSms,
		func(i context.Context, messager Messager) (*sqlx.Tx, error) {
			return db.SmsInsert(i, messager.(*meta.DbSms))
		},
		mq.SmsProduce)
}

func (s smsServiceImpl) Detail(ctx context.Context, id string) (Marshaler, error) {
	return s.detail(ctx, id)
}

// DetailByPhonePage 直接数据库中取,不走缓存
func (s smsServiceImpl) DetailByPhonePage(ctx context.Context, mobile string, page int) ([]*meta.DbSms, error) {
	return db.SmsDetailByPhoneAndPage(ctx, mobile, page)
}

func (s smsServiceImpl) DetailByPlat(ctx context.Context, plat int32, key string) ([]*meta.DbSms, error) {
	return db.SmsDetailByPlat(ctx, plat, key)
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
		return db.SmsCancelByID(i, s)
	}, &meta.DbSms{})
}

func (s smsServiceImpl) Edit(ctx context.Context, m Meta) error {
	dbParam := &meta.DbSms{}
	return edit(ctx,
		m,
		dbParam,
		func(i context.Context, messager Messager) (*sqlx.Tx, error) {
			return db.SmsEdit(i, messager.(*meta.DbSms))
		},
		mq.SmsProduce,
	)
}

func (s smsServiceImpl) checkSendRate(ctx context.Context, mobile string) error {
	m1, err := cache.MobileCache1Min(ctx, mobile)
	if err != nil {
		return err
	}
	if m1 > 1 {
		return errors.ErrMsg1MinuteLimit
	}
	m2, err := cache.MobileCache1Hour(ctx, mobile)
	if err != nil {
		return err
	}
	if m2 > 5 {
		return errors.ErrMsg1HourLimit
	}
	m3, err := cache.MobileCache1Day(ctx, mobile)
	if err != nil {
		return err
	}
	if m3 > 10 {
		return errors.ErrMsg1DayLimit
	}
	return nil
}
