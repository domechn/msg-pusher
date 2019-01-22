/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : wechat.go
#   Created       : 2019/1/16 10:34
#   Last Modified : 2019/1/16 10:34
#   Describe      :
#
# ====================================================*/
package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/storer/db"
	"uuabc.com/sendmsg/storer/mq"
)

type weChatServiceImpl struct {
}

func NewWeChatServiceImpl() weChatServiceImpl {
	return weChatServiceImpl{}
}

func (s weChatServiceImpl) Produce(ctx context.Context, m Meta) (string, error) {
	var templ string
	var args map[string]string
	var err error
	if templ, args, err = checkTemplateAndArguments(ctx, m.GetTemplate(), m.GetArguments()); err != nil {
		return "", err
	}
	content := getContent(args, templ)
	ttl := m.Delay()
	err = s.produce(ctx, m.(*meta.WeChatProducer), content, ttl)
	return m.GetId(), err
}

func (weChatServiceImpl) produce(ctx context.Context, p *meta.WeChatProducer, content string, ttl int64) error {
	dbWeChat := &meta.DbWeChat{
		Id:          p.Id,
		Platform:    p.Platform,
		PlatformKey: p.PlatformKey,
		Touser:      p.Touser,
		Type:        p.Type,
		Content:     content,
		Template:    p.Template,
		Url:         p.Url,
		Arguments:   p.Arguments,
		SendTime:    p.SendTime,
	}
	return produce(ctx,
		p,
		dbWeChat,
		func(i context.Context, messager Messager) (*sqlx.Tx, error) {
			return db.WeChatInsert(i, messager.(*meta.DbWeChat))
		},
		mq.WeChatProduce)
}

func (s weChatServiceImpl) Detail(ctx context.Context, id string) (Marshaler, error) {
	return s.detail(ctx, id)
}

func (s weChatServiceImpl) detail(ctx context.Context, id string) (Marshaler, error) {
	res := &meta.DbWeChat{}
	return res, detail(ctx, id, res, func(ctx2 context.Context, id string) (Marshaler, error) {
		return db.WeChatDetailByID(ctx2, id)
	})
}

// Cancel 取消微信发送
func (s weChatServiceImpl) Cancel(ctx context.Context, id string) error {
	return s.cancel(ctx, id)
}

func (weChatServiceImpl) cancel(ctx context.Context, id string) error {
	return cancel(ctx, id, func(i context.Context, s string) (*sqlx.Tx, error) {
		return db.WeChatCancelMsgByID(i, s)
	}, &meta.DbSms{})
}

func (s weChatServiceImpl) Edit(ctx context.Context, m Meta) error {
	dbParam := &meta.DbWeChat{}
	return edit(ctx,
		m,
		dbParam,
		func(i context.Context, messager Messager) (*sqlx.Tx, error) {
			return db.WeChatEdit(i, messager.(*meta.DbWeChat))
		},
		mq.WeChatProduce,
	)
}
