/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : service.go
#   Created       : 2019/1/16 10:24
#   Last Modified : 2019/1/16 10:24
#   Describe      :
#
# ====================================================*/
package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
	cache2 "uuabc.com/sendmsg/pkg/cache"
	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/storer/cache"
	"uuabc.com/sendmsg/storer/db"
)

type updateFunc func(i context.Context, s string) (*sqlx.Tx, error)

const (
	timeLayout = "2006-01-02T15:03:04Z"
)

type MsgService interface {
	Produce(ctx context.Context, m Meta) (string, error)
	Detail(ctx context.Context, id string) (Marshaler, error)
	Cancel(ctx context.Context, id string) error
	Edit(ctx context.Context, m Meta) error
}

// updateCache 更新缓存
func updateCache(ctx context.Context, id string, msg Messager) {
	// 统一和数据库中的信息
	now := time.Now().UTC().Format(timeLayout)
	msg.SetStatus(int32(meta.Status_Wait))
	msg.SetCreatedAt(now)
	msg.SetUpdatedAt(now)

	byt, gErr := msg.Marshal()
	if gErr != nil {
		logrus.Errorf("set cache go func() error: %v", gErr)
		return
	}
	// 插入redis
	cache.PutBaseCache(ctx, id, byt)
}

func edit(ctx context.Context, msg Messager, m Meta, mqFunc func(context.Context, []byte, int64) error) error {
	now := time.Now().UTC().Format(timeLayout)
	msg.SetUpdatedAt(now)
	b, err := msg.Marshal()
	if err != nil {
		return err
	}
	// 如果mq发送失败就回滚返回错误，即使redis中更新失败了了，
	// mq中的数据无法回滚也不会有影响，因为在发送时会去获取缓存中的数据，
	// 所以只要保证缓存中的数据和数据库中的数据一致并有效就行
	err = mqFunc(ctx, []byte(m.GetId()), m.Delay())
	if err != nil {
		return err
	}
	err = cache.PutBaseCache(ctx, m.GetId(), b)
	if err != nil {
		return err
	}
	cache.PutLastestCache(ctx, m.GetId(), b)

	return nil
}

func updateDetailCache(ctx context.Context, id string, getDbData func(ctx2 context.Context, id string) (Marshaler, error)) {
	// 5秒内更新一次
	gErr := cache.LockID5s(context.Background(), id)
	if gErr != nil {
		logrus.WithField("id", id).Errorf("频繁请求更新不存在的key。")
		return
	}
	dt, err := getDbData(ctx, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"id":    id,
			"error": err,
		}).Errorf("从数据库中获取数据失败，无法更新缓存")
		return
	}
	// 更新缓存
	dbRes, dbErr := dt.Marshal()
	if dbErr != nil {
		logrus.Errorf("后台通过数据库更新cache失败，key:%s,error: %v", id, dbErr)
		return
	}
	cache.PutBaseCache(context.Background(), id, dbRes)
	cache.PutLastestCache(context.Background(), id, dbRes)
	logrus.WithField("id", id).Errorf("后台通过数据库添加cache成功")
}

func detail(ctx context.Context, id string, res Marshaler, getDbData func(ctx2 context.Context, id string) (Marshaler, error)) error {
	var data []byte
	var err error
	data, e1 := cache.LastestDetail(ctx, id)
	if e1 == nil {
		logrus.WithFields(logrus.Fields{
			"data": string(data),
			"id":   id,
		}).Infof("在lastest缓存中找到数据，直接返回结果")
	} else {
		data, err = cache.BaseDetail(ctx, id)
	}
	// 如果最新缓存不存在，则更新最新缓存和基础缓存
	if e1 == cache2.ErrCacheMiss {
		go updateDetailCache(
			context.Background(),
			id,
			getDbData)
	}
	if err != nil {
		return err
	}
	err = res.Unmarshal(data)
	return err
}

func cancel(ctx context.Context, id string, u updateFunc, m Messager) error {
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
	if st := m.GetStatus(); st == int32(meta.Status_Cancel) {
		return errors.ErrMsgHasCancelled
	}
	// 如果已发送
	if st := m.GetStatus(); st == int32(meta.Status_Final) {
		return errors.ErrMsgCantEdit
	}

	logrus.Debug("缓存中的数据状态为可取消状态")

	// 更新数据库中的status
	tx, err := u(ctx, id)
	if err != nil {
		db.RollBack(tx)
		// 如果没有行被修改说明有别的线程已经修改了该条数据的状态
		if err == db.ErrNoRowsEffected {
			return errors.ErrMsgHasCancelled
		}
		return err
	}

	// 获取数据库中的值，并更新到缓存,必须同步，因为发送的时候需要检查缓存中信息的状态
	// 如果异步操作失败，会导致已取消的信息发送
	var b []byte
	m.SetStatus(int32(meta.Status_Cancel))
	m.SetResult(int32(meta.Result_Fail))
	b, _ = m.Marshal()

	// 该方法并发安全
	if err := cache.PutBaseCache(ctx, id, b); err != nil {
		db.RollBack(tx)
		logrus.WithFields(logrus.Fields{
			"method": "putBaseCache",
			"id":     id,
			"error":  err,
		}).Error("在取消发送后更新baseCache时出现错误")
		return err
	}
	if err := cache.PutLastestCache(ctx, id, b); err != nil {
		// 更新最新状态失败，无需回滚
		// rollback(tx)
		logrus.WithFields(logrus.Fields{
			"method": "putLastestCache",
			"id":     id,
			"error":  err,
		}).Error("在取消发送后更新lastestCache时出现错误")
		return err
	}

	return db.Commit(tx)
}

func checkStatus(id string, msg Messager) error {
	b, err := cache.BaseDetail(context.Background(), id)
	// ttl := m.Delay()
	if err != nil {
		return err
	}
	err = msg.Unmarshal(b)
	if err != nil {
		return err
	}
	st := msg.GetStatus()
	if st == int32(meta.Status_Cancel) {
		return errors.ErrMsgHasCancelled
	}
	if st == int32(meta.Status_Final) {
		return errors.ErrMsgCantEdit
	}
	return nil
}
