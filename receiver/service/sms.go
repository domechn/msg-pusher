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

	"github.com/sirupsen/logrus"
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
	err = s.produce(ctx, m.(*meta.SmsProducer), content)
	return m.GetId(), err
}

func (s smsServiceImpl) produce(ctx context.Context, p *meta.SmsProducer, content string) error {
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
		s.rPush,
		mq.SmsProduce)
}

func (s smsServiceImpl) ProduceBatch(ctx context.Context, ms []*meta.SmsProducer) ([]string, error) {
	var res []string
	var byts [][]byte
	var metas []*meta.DbSms
	for _, m := range ms {
		var templ string
		var args map[string]string
		var err error
		mobile := m.GetSendTo()
		if err := s.checkSendRate(ctx, mobile); err != nil {
			return nil, err
		}
		if templ, args, err = checkTemplateAndArguments(ctx, m.GetTemplate(), m.GetArguments()); err != nil {
			return nil, err
		}
		dbSms := &meta.DbSms{
			Id:          m.Id,
			Platform:    m.Platform,
			PlatformKey: m.PlatformKey,
			Content:     getContent(args, templ),
			Mobile:      m.Mobile,
			Template:    m.Template,
			Arguments:   m.Arguments,
			SendTime:    m.SendTime,
			Server:      m.Server,
			Status:      int32(meta.Status_Wait),
			Type:        m.Type,
		}
		res = append(res, m.Id)
		initMsg(dbSms)
		b, _ := dbSms.Marshal()
		byts = append(byts, b)
		metas = append(metas, dbSms)
	}
	return res, s.produceBatch(ctx, byts, metas, ms)
}

func (s smsServiceImpl) produceBatch(ctx context.Context, byts [][]byte, metas []*meta.DbSms, ms []*meta.SmsProducer) error {
	// 开启redis事务
	t := cache.NewTransaction()
	defer t.Close()

	if len(byts) != len(metas) || len(byts) != len(ms) {
		return errors.NewError(10000000, "未知错误")
	}

	for idx, byt := range byts {
		if err := produceStore(ctx, metas[idx].Id, byt, ms[idx].Delay(), t, mq.SmsProduce, s.rPush); err != nil {
			logrus.WithFields(logrus.Fields{
				"method": "produceSmsBatch",
				"error":  err.Error(),
			}).Error("批量发送短信失败")
			t.Rollback()
			return err
		}
	}
	return t.Commit()
}

func (smsServiceImpl) rPush(ctx context.Context, c Cache, b []byte) error {
	return c.RPushSms(ctx, b)
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
	return res, detail(ctx, id, res)
}

func (s smsServiceImpl) Cancel(ctx context.Context, id string) error {
	return s.cancel(ctx, id)
}

func (s smsServiceImpl) cancel(ctx context.Context, id string) error {
	return cancel(ctx,
		id,
		s.rPush,
		&meta.DbSms{})
}

func (s smsServiceImpl) Edit(ctx context.Context, m Meta) error {
	dbParam := &meta.DbSms{}
	return edit(ctx,
		m,
		dbParam,
		s.rPush,
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
