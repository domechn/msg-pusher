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
	"time"

	"github.com/domgoer/msg-pusher/pkg/pb/meta"
	"github.com/domgoer/msg-pusher/storer/cache"
	"github.com/domgoer/msg-pusher/storer/db"
)

type msgServiceImpl struct {
}

// NewMsgServiceImpl 初始化消息service
func NewMsgServiceImpl() msgServiceImpl {
	return msgServiceImpl{}
}

// Produce 接收消息，并将验证通过的消息存入mq和缓存
func (s msgServiceImpl) Produce(ctx context.Context, m *meta.MsgProducer) (string, error) {
	var templ string
	var args map[string]string
	var err error
	to := m.GetSendTo()
	sendTimeStr := m.GetSendTime()
	ti, err := time.Parse(meta.ISO8601Layout, sendTimeStr)
	if err != nil {
		return "", err
	}
	if err := checkSendRate(ctx, to, ti); err != nil {
		return "", err
	}
	if templ, args, err = checkTemplateAndArguments(ctx, m.GetTemplate(), m.GetArguments()); err != nil {
		return "", err
	}
	content := getContent(args, templ)
	err = s.produce(ctx, m, content)
	return m.GetId(), err
}

func (s msgServiceImpl) produce(ctx context.Context, p *meta.MsgProducer, content string) error {
	dbMsg := &meta.DbMsg{
		Id:        p.Id,
		SubId:     p.SubId,
		Content:   content,
		SendTo:    p.SendTo,
		Reserved:  p.Reserved,
		Template:  p.Template,
		Arguments: p.Arguments,
		SendTime:  p.SendTime,
		Server:    meta.Server(p.Server),
		Type:      meta.Type(p.Type),
	}
	return produce(ctx,
		p,
		dbMsg,
	)
}

// Detail 根据id获取消息的详情
func (s msgServiceImpl) Detail(ctx context.Context, id string) (Marshaler, error) {
	return s.detail(ctx, id)
}

// DetailByPhonePage 直接数据库中取,不走缓存
func (s msgServiceImpl) DetailByToAndPage(ctx context.Context, to string, page int) ([]*meta.DbMsg, error) {
	return db.DetailByToAndPage(ctx, to, page)
}

// DetailByPlat 根据key值分页查询，直接查询数据库
func (s msgServiceImpl) DetailByKeyAndPage(ctx context.Context, key string, page int) ([]*meta.DbMsg, error) {
	return db.DetailByKey(ctx, key, page)
}

func (msgServiceImpl) detail(ctx context.Context, id string) (Marshaler, error) {
	res := &meta.DbMsg{}
	return res, detail(ctx, id, res)
}

func (s msgServiceImpl) Cancel(ctx context.Context, id string) error {
	return s.cancel(ctx, id)
}

func (s msgServiceImpl) cancel(ctx context.Context, id string) error {
	ds := &meta.DbMsg{}
	if err := cancel(ctx,
		id,
		ds); err != nil {
		return err
	}
	go func() {
		sendTimeStr := ds.GetSendTime()
		ti, err := time.Parse(meta.ISO8601Layout, sendTimeStr)
		if err != nil {
			return
		}
		// 删除限速的限制
		cache.RemoveLimit(context.Background(), ds.GetSendTo(), ti)
	}()
	return nil
}

func (s msgServiceImpl) Edit(ctx context.Context, m Meta) error {
	dbParam := &meta.DbMsg{}
	return edit(ctx,
		m,
		dbParam,
	)
}

// WaitMsgIdByPlat 按平台号获取待发送的消息
func (s msgServiceImpl) WaitingMsgIdByPlat(ctx context.Context, key string) ([]string, error) {
	return db.WaitingMsgByKey(ctx, key)
}

// CancelBatch 批量取消,return 取消发送失败的消息的id
func (s msgServiceImpl) CancelBatch(ctx context.Context, ids []string) []string {
	var fail []string
	for _, id := range ids {
		if err := s.cancel(ctx, id); err != nil {
			fail = append(fail, id)
		}
	}
	return fail
}

// 监测发送速率
func checkSendRate(ctx context.Context, mobile string, sendTime time.Time) error {
	return cache.RateLimit(ctx, mobile, sendTime)
}
