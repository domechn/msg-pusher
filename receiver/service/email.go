/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : email.go
#   Created       : 2019/1/16 10:40
#   Last Modified : 2019/1/16 10:40
#   Describe      :
#
# ====================================================*/
package service

import (
	"context"

	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/storer/mq"
)

type emailServiceImpl struct {
}

func NewEmailSeriveImpl() emailServiceImpl {
	return emailServiceImpl{}
}

// Produce 接收要发送的email信息，并保存
func (s emailServiceImpl) Produce(ctx context.Context, m Meta) (string, error) {
	var templ string
	var args map[string]string
	var err error
	if templ, args, err = checkTemplateAndArguments(ctx, m.GetTemplate(), m.GetArguments()); err != nil {
		return "", err
	}
	content := getContent(args, templ)
	err = s.produce(ctx, m.(*meta.EmailProducer), content)
	return m.GetId(), err
}

func (e emailServiceImpl) produce(ctx context.Context, p *meta.EmailProducer, content string) error {
	dbEmail := &meta.DbEmail{
		Id:          p.Id,
		Platform:    p.Platform,
		PlatformKey: p.PlatformKey,
		Title:       p.Title,
		Content:     content,
		Destination: p.Destination,
		Type:        p.Type,
		Template:    p.Template,
		Arguments:   p.Arguments,
		Server:      p.Server,
		SendTime:    p.SendTime,
	}
	return produce(ctx,
		p,
		dbEmail,
		e.rPush,
		mq.EmailProduce)
}

func (emailServiceImpl) rPush(ctx context.Context, c Cache, b []byte) error {
	return c.RPushEmail(ctx, b)
}

// Detail 返回要发送的email的具体信息
func (s emailServiceImpl) Detail(ctx context.Context, id string) (Marshaler, error) {
	return s.detail(ctx, id)
}

func (s emailServiceImpl) detail(ctx context.Context, id string) (Marshaler, error) {
	res := &meta.DbEmail{}
	return res, detail(ctx, id, res)
}

func (s emailServiceImpl) Cancel(ctx context.Context, id string) error {
	return s.cancel(ctx, id)
}

func (e emailServiceImpl) cancel(ctx context.Context, id string) error {
	return cancel(ctx,
		id,
		e.rPush,
		&meta.DbEmail{})
}

// 修改信息
func (s emailServiceImpl) Edit(ctx context.Context, m Meta) error {
	dbParam := &meta.DbEmail{}
	return edit(ctx,
		m,
		dbParam,
		s.rPush,
		mq.EmailProduce,
	)
}
