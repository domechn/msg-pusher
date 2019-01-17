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
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/storer/db"
	"uuabc.com/sendmsg/storer/mq"
)

type emailServiceImpl struct {
}

func NewEmailSeriveImpl() emailServiceImpl {
	return emailServiceImpl{}
}

// Produce 接收要发送的email信息，并保存
func (s emailServiceImpl) Produce(ctx context.Context, m Meta) (string, error) {
	if err := checkTemplateAndArguments(m.GetTemplate(), m.GetArguments()); err != nil {
		return "", err
	}
	ttl := m.Delay()
	err := s.produce(ctx, m.(*meta.EmailProducer), ttl)
	return m.GetId(), err
}

func (emailServiceImpl) produce(ctx context.Context, p *meta.EmailProducer, ttl int64) error {
	dbEmail := &meta.DbEmail{
		Id:          p.Id,
		Platform:    p.Platform,
		PlatformKey: p.PlatformKey,
		Title:       p.Title,
		Content:     "",
		Destination: p.Destination,
		Type:        p.Type,
		Template:    p.Template,
		Arguments:   p.Arguments,
		Server:      p.Server,
		SendTime:    p.SendTime,
	}
	id := dbEmail.Id
	tx, err := db.EmailInsert(ctx, dbEmail)
	if err != nil {
		return err
	}
	err = mq.EmailProduce(ctx, []byte(id), ttl)
	if err != nil {
		db.RollBack(tx)
		logrus.WithField("type", "Email").Errorf("消息 %s 插入消息队列失败，正在回滚。。。，error: %v\n", id, err)
		return err
	}
	logrus.WithField("type", "Email").Infof("消息 %s 插入消息队列成功,正在等待发送,开始提交到数据库", id)
	err = db.Commit(tx)
	if err != nil {
		return err
	}
	go updateCache(context.Background(), id, dbEmail)
	logrus.WithField("type", "Email").Infof("消息 %s 插入数据库成功", id)
	return nil
}

// Detail 返回要发送的email的具体信息
func (s emailServiceImpl) Detail(ctx context.Context, id string) (Marshaler, error) {
	return s.detail(ctx, id)
}

func (s emailServiceImpl) detail(ctx context.Context, id string) (Marshaler, error) {
	res := &meta.DbEmail{}
	return res, detail(ctx, id, res, func(ctx2 context.Context, id string) (Marshaler, error) {
		return db.EmailDetailByID(ctx2, id)
	})
}

func (s emailServiceImpl) Cancel(ctx context.Context, id string) error {
	return s.cancel(ctx, id)
}

func (emailServiceImpl) cancel(ctx context.Context, id string) error {
	return cancel(ctx, id, func(i context.Context, s string) (*sqlx.Tx, error) {
		return db.EmailCancelMsgByID(i, s)
	}, &meta.DbSms{})
}

func (s emailServiceImpl) Edit(ctx context.Context, m Meta) error {
	m.Transfer(false)
	v := m.(*meta.EmailProducer)
	return s.edit(
		ctx,
		m,
		&meta.DbEmail{
			Id:          v.Id,
			Arguments:   v.Arguments,
			SendTime:    v.SendTime,
			Destination: v.Destination,
		})
}

func (s emailServiceImpl) edit(ctx context.Context, m Meta, e *meta.DbEmail) error {
	// 用于更新缓存
	em := &meta.DbEmail{}
	if err := checkStatus(m.GetId(), em); err != nil {
		return err
	}

	// 修改数据
	em.Arguments = e.Arguments
	em.SendTime = e.SendTime
	if e.Destination != "" {
		em.Destination = e.Destination
	}

	tx, err := db.EmailEdit(ctx, e)
	if err != nil {
		db.RollBack(tx)
		logrus.WithField("type", "Email").Errorf("edit修改数据库失败,error: %v", err)
		return err
	}

	err = edit(ctx, em, m, mq.EmailProduce)
	if err != nil {
		db.RollBack(tx)
		logrus.WithField("type", "Email").Errorf("edit更新mq失败，正在回滚,error: %v", err)
		return err
	}

	return db.Commit(tx)
}
