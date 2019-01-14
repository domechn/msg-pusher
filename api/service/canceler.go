/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : canceler.go
#   Created       : 2019/1/10 19:58
#   Last Modified : 2019/1/10 19:58
#   Describe      :
#
# ====================================================*/
package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/api/model"
	"uuabc.com/sendmsg/api/storer/cache"
	"uuabc.com/sendmsg/api/storer/db"
	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/pkg/pb/meta"
)

var Canceler = newCancelerImpl()

type getFunc func(context.Context, string) (Messager, error)
type updateFunc func(i context.Context, s string) (*sqlx.Tx, error)

type cancelerImpl struct {
	w cancelWeChatImpl
	e cancelEmailImpl
	s cancelSmsImpl
}

type cancelWeChatImpl struct {
	g getFunc
	u updateFunc
}

type cancelEmailImpl struct {
	g getFunc
	u updateFunc
}

type cancelSmsImpl struct {
	g getFunc
	u updateFunc
}

func newCancelerImpl() cancelerImpl {
	wi := cancelWeChatImpl{
		g: func(i context.Context, s string) (Messager, error) {
			return db.WeChatDetailByID(i, s)
		},
		u: func(i context.Context, s string) (*sqlx.Tx, error) {
			return db.WeChatCancelMsgByID(i, s)
		},
	}
	ei := cancelEmailImpl{
		g: func(i context.Context, s string) (Messager, error) {
			return db.EmailDetailByID(i, s)
		},
		u: func(i context.Context, s string) (*sqlx.Tx, error) {
			return db.EmailCancelMsgByID(i, s)
		},
	}
	si := cancelSmsImpl{
		g: func(i context.Context, s string) (Messager, error) {
			return db.SmsDetailByID(i, s)
		},
		u: func(i context.Context, s string) (*sqlx.Tx, error) {
			return db.SmsCancelMsgByID(i, s)
		},
	}
	return cancelerImpl{
		w: wi,
		e: ei,
		s: si,
	}
}

// Cancel 根据id取消发送消息
func (c cancelerImpl) Cancel(ctx context.Context, typeN, id string) error {
	err := c.cancel(ctx, typeN, id)
	if err == nil {
		return nil
	}
	if _, ok := err.(*errors.Error); !ok {
		err = errors.NewError(
			10000000,
			err.Error(),
		)
	}
	return err
}

func (c cancelerImpl) cancel(ctx context.Context, typeN, id string) (err error) {
	switch typeN {
	case wechat:
		err = cancel(ctx, id, c.w.g, c.w.u, &model.DbWeChat{})
	case sms:
		err = cancel(ctx, id, c.s.g, c.s.u, &model.DbSms{})
	case email:
		err = cancel(ctx, id, c.e.g, c.e.u, &model.DbEmail{})
	default:
		return errors.ErrMsgTypeNotFound
	}

	return
}

func cancel(ctx context.Context, id string, g getFunc, u updateFunc, m Messager) error {
	data, err := cache.BaseDetail(ctx, id)
	if err != nil {
		logrus.Errorf("从缓存中获取要取消的数据失败，id: %s,error: %v", id, err)
		return errors.ErrMsgNotFound
	}

	logrus.Debug("从缓存中取出数据检查状态")
	// 查看缓存中的数据的状态
	err = m.Unmarshal(data)
	if err != nil {
		return err
	}
	// 如果已取消
	if st := m.GetStatus(); st == meta.Status_Cancel {
		return errors.ErrMsgHasCancelled
	}
	// 如果已发送
	if st := m.GetStatus(); st == meta.Status_Final {
		return errors.ErrMsgCantEdit
	}

	logrus.Debug("缓存中的数据状态为可取消状态")

	// 更新数据库中的status
	tx, err := u(ctx, id)
	if err != nil {
		rollback(tx)
		// 如果没有行被修改说明有别的线程已经修改了该条数据的状态
		if err == db.ErrNoRowsEffected {
			return errors.ErrMsgHasCancelled
		}
		return err
	}

	// 获取数据库中的值，并更新到缓存,必须同步，因为发送的时候需要检查缓存中信息的状态
	// 如果异步操作失败，会导致已取消的信息发送
	msg, err := g(ctx, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"method": "getFunc",
			"id":     id,
			"error":  err,
		}).Error("从数据库中获取数据失败")
		rollback(tx)
		return err
	}
	b, err := msg.Marshal()
	if err != nil {
		rollback(tx)
		return err
	}

	// 该方法并发安全
	if err := cache.PutBaseCache(id, b); err != nil {
		rollback(tx)
		logrus.WithFields(logrus.Fields{
			"method": "putBaseCache",
			"id":     id,
			"error":  err,
		}).Error("在取消发送后更新baseCache时出现错误")
		return err
	}
	if err := cache.PutLastestCache(id, b); err != nil {
		// 更新最新状态失败，无需回滚
		// rollback(tx)
		logrus.WithFields(logrus.Fields{
			"method": "putLastestCache",
			"id":     id,
			"error":  err,
		}).Error("在取消发送后更新lastestCache时出现错误")
		return err
	}

	return commit(tx)
}
