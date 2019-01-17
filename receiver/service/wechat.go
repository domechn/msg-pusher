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
	"github.com/sirupsen/logrus"
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
	if err := checkTemplateAndArguments(m.GetTemplate(), m.GetArguments()); err != nil {
		return "", err
	}
	ttl := m.Delay()
	err := s.produce(ctx, m.(*meta.WeChatProducer), ttl)
	return m.GetId(), err
}

func (weChatServiceImpl) produce(ctx context.Context, p *meta.WeChatProducer, ttl int64) error {
	dbWeChat := &meta.DbWeChat{
		Id:        p.Id,
		Platform:  p.Platform,
		Touser:    p.Touser,
		Type:      p.Type,
		Template:  p.Template,
		Url:       p.Url,
		Arguments: p.Arguments,
		SendTime:  p.SendTime,
	}
	tx, err := db.WeChatInsert(ctx, dbWeChat)
	if err != nil {
		return err
	}
	id := dbWeChat.Id
	err = mq.WeChatProduce(ctx, []byte(id), ttl)
	if err != nil {
		db.RollBack(tx)
		logrus.WithField("type", "WeChat").Errorf("消息 %s 插入消息队列失败，正在回滚。。。，error: %v\n", id, err)
		return err
	}
	logrus.WithField("type", "WeChat").Infof("消息 %s 插入消息队列成功,正在等待发送,开始提交到数据库", id)
	err = db.Commit(tx)
	if err != nil {
		return err
	}

	go updateCache(context.Background(), id, dbWeChat)
	logrus.WithField("type", "WeChat").Infof("消息 %s 插入数据库成功", id)
	return nil
}

func (s weChatServiceImpl) Detail(ctx context.Context, id string) (Marshaler, error) {
	return s.detail(ctx, id)
}

func (s weChatServiceImpl) detail(ctx context.Context, id string) (Marshaler, error) {
	res := &meta.DbEmail{}
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
	m.Transfer(false)
	v := m.(*meta.WeChatProducer)
	return s.edit(
		ctx,
		m,
		&meta.DbWeChat{
			Id:        v.Id,
			Arguments: v.Arguments,
			SendTime:  v.SendTime,
			Touser:    v.Touser,
		})
}

// todo 验证参数
func (weChatServiceImpl) edit(ctx context.Context, m Meta, e *meta.DbWeChat) error {
	em := &meta.DbWeChat{}
	if err := checkStatus(m.GetId(), em); err != nil {
		return err
	}

	// 修改数据
	em.Arguments = e.Arguments
	em.SendTime = e.SendTime
	if e.Touser != "" {
		em.Touser = e.Touser
	}

	tx, err := db.WeChatEdit(ctx, e)
	if err != nil {
		db.RollBack(tx)
		logrus.WithField("type", "WeChat").Errorf("edit修改数据库失败,error: %v", err)
		return err
	}

	err = edit(ctx, em, m, mq.WeChatProduce)
	if err != nil {
		db.RollBack(tx)
		logrus.WithField("type", "WeChat").Errorf("edit更新mq失败，正在回滚,error: %v", err)
		return err
	}

	return db.Commit(tx)
}
