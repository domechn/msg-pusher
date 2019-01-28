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

	"github.com/hiruok/msg-pusher/pkg/errors"
	"github.com/hiruok/msg-pusher/pkg/pb/meta"
	"github.com/hiruok/msg-pusher/pkg/utils"
	"github.com/hiruok/msg-pusher/storer/cache"
	"github.com/hiruok/msg-pusher/storer/mq"
	"github.com/sirupsen/logrus"
)

type MqFunc func(context.Context, []byte, int64) error

// MsgService 用于消息的增删改查
type MsgService interface {
	Produce(ctx context.Context, m Meta) (string, error)
	Detail(ctx context.Context, id string) (Marshaler, error)
	Cancel(ctx context.Context, id string) error
	Edit(ctx context.Context, m Meta) error
}

// detail 根据id查询消息
func detail(ctx context.Context, id string, msg Messager) error {
	b, err := cache.BaseDetail(ctx, id)
	if err != nil {
		return err
	}
	err = msg.Unmarshal(b)
	if err != nil {
		return errors.ErrMisMatch
	}
	return nil
}

// produce 将消息插入数据库 mq 和缓存
func produce(ctx context.Context, m Meta, em Messager) error {
	logrus.Info("开始添加消息...,data: ", em)
	id := em.GetId()

	em.SetStatus(meta.Wait)
	initMsg(em)
	byt, _ := em.Marshal()
	// 开启redis事务
	t := cache.NewTransaction()
	defer t.Close()
	if err := produceStore(ctx, id, byt, m.Delay(), t); err != nil {
		t.Rollback(ctx)
		return err
	}
	logrus.Infof("消息添加成功,id: %s", id)
	return t.Commit(ctx)
}

// 初始化信息，设置信息的创建修改时间和版本号
func initMsg(em Messager) {
	timeStamp := utils.NowTimeStampStr()
	em.SetCreatedAt(timeStamp)
	em.SetUpdatedAt(timeStamp)
	em.SetVersion(1)
}

// produceStore 向cache和mq中提交数据
func produceStore(ctx context.Context,
	id string,
	b []byte,
	ttl int64,
	t *cache.Transaction,
) error {
	// 插入redis
	err := t.PutBaseCache(ctx, id, b)
	if err != nil {
		logrus.Errorf("消息插入缓存时失败,errors: %v", err)
		return err
	}
	// 将数据添加到list，用于入库
	if err := t.RPushMsg(ctx, b); err != nil {
		logrus.Errorf("消息插入入库队列时失败,errors: %v", err)
		return err
	}
	err = mq.MsgProduce(ctx, []byte(id), ttl)
	if err != nil {
		logrus.Errorf("消息 %s 插入消息队列失败: %v\n", id, err)
		return err
	}
	return nil
}

// todo 限速
func edit(ctx context.Context, m Meta, em Messager) error {
	// 先根据id从缓存中获取数据的具体内容
	if err := detail(ctx, m.GetId(), em); err != nil {
		return err
	}
	// 判断消息的发送状态
	if err := checkStatus(em); err != nil {
		return err
	}
	var needEdit bool
	var templateArgumentsChanged bool
	// 是否需要重新发送
	var reSend bool
	var ttl int64 = -1
	// 按照传来的字段，将修改的字段修改到元字段上
	if m.GetSendTo() != "" && m.GetSendTo() != em.GetSendTo() {
		needEdit = true
		em.SetSendTo(m.GetSendTo())
	}
	if m.GetArguments() != "" && m.GetArguments() != em.GetArguments() {
		needEdit = true
		templateArgumentsChanged = true
		em.SetArguments(m.GetArguments())
	}
	if m.GetTemplate() != "" && m.GetTemplate() != em.GetTemplate() {
		needEdit = true
		templateArgumentsChanged = true
		em.SetTemplate(m.GetTemplate())
	}
	if m.GetSendTime() != "" {
		m.Transfer(false)
		if m.GetSendTime() != em.GetSendTime() {
			needEdit = true
			reSend = true
			ttl = m.Delay()
			em.SetSendTime(m.GetSendTime())
		}
	}
	if !needEdit {
		return errors.ErrMsgIsSameBefore
	}
	// 参数验证
	// 如果参数和模板没有发生变化就不检查
	if templateArgumentsChanged {
		var templ string
		var args map[string]string
		var err error
		if templ, args, err = checkTemplateAndArguments(ctx, em.GetTemplate(), em.GetArguments()); err != nil {
			return err
		}
		content := getContent(args, templ)
		em.SetContent(content)
	}
	// redis中设置分布式锁
	if err := cache.LockId(ctx, m.GetId()); err != nil {
		logrus.WithField("id", m.GetId()).Error("获取分布式锁失败，消息可能正在被其他线程在处理")
		return errors.ErrMsgBusy
	}
	logrus.WithField("id", m.GetId()).Info("获取分布式锁成功，正在修改消息")
	defer cache.UnlockId(ctx, m.GetId())
	var mqFunc MqFunc
	// 判断sendtime是否改变，如果改变就向mq中重新发送id
	if reSend {
		mqFunc = mq.MsgProduce
	}
	// 修改修改时间
	changeUpdateAndVersion(em)
	b, _ := em.Marshal()
	t := cache.NewTransaction()
	defer t.Close()

	// 修改cache和mq中信息
	if err := editStore(ctx, em.GetId(), b, ttl, t, mqFunc); err != nil {
		t.Rollback(ctx)
		logrus.WithFields(logrus.Fields{
			"method": "updateStore",
			"id":     em.GetId(),
			"error":  err.Error(),
		}).Errorf("edit修改list失败,error: %v", err)
		return err
	}

	// 提交事务
	err := t.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func editStore(ctx context.Context, id string, b []byte, ttl int64, t *cache.Transaction, mqf MqFunc) error {
	// 异步修改db中的数据
	err := t.RPushMsg(ctx, b)
	if err != nil {
		return err
	}
	err = t.PutBaseCache(ctx, id, b)
	if err != nil {
		return err
	}

	if mqf != nil {
		err = mqf(ctx, []byte(id), ttl)
		if err != nil {
			return err
		}
	}
	return nil
}

// 内部方法，在detail()中被调用,不启用
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
	// 获取锁
	if err = cache.LockId(ctx, id); err != nil {
		return
	}
	defer cache.UnlockId(ctx, id)
	cache.PutBaseCache(context.Background(), id, dbRes)
	logrus.WithField("id", id).Errorf("后台通过数据库添加cache成功")
}

// func detail(ctx context.Context, id string, res Marshaler, getDbData func(ctx2 context.Context, id string) (Marshaler, error)) error {
// 	var data []byte
// 	var err error
// 	data, e1 := cache.LastestDetail(ctx, id)
// 	if e1 == nil {
// 		logrus.WithFields(logrus.Fields{
// 			"data": string(data),
// 			"id":   id,
// 		}).Infof("在lastest缓存中找到数据，直接返回结果")
// 	} else {
// 		data, err = cache.BaseDetail(ctx, id)
// 	}
// 	// 如果最新缓存不存在，则更新最新缓存和基础缓存
// 	if e1 == cache2.ErrCacheMiss {
// 		go updateDetailCache(
// 			context.Background(),
// 			id,
// 			getDbData)
// 	}
// 	if err != nil {
// 		return err
// 	}
// 	err = res.Unmarshal(data)
// 	return err
// }

func cancel(ctx context.Context, id string, m Messager) error {
	err := detail(ctx, id, m)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"method": "cancel",
			"id":     id,
			"error":  err.Error(),
		}).Errorf("从缓存中获取要取消的数据失败")
		return err
	}
	if err = checkStatus(m); err != nil {
		logrus.WithFields(logrus.Fields{
			"method": "cancel",
			"id":     id,
			"error":  err.Error(),
		}).Error("状态检查失败")
		return err
	}
	if err := cache.LockId(ctx, id); err != nil {
		logrus.WithField("id", id).Error("获取分布式锁失败，消息可能正在被其他线程在处理")
		return errors.ErrMsgBusy
	}
	logrus.WithField("id", id).Info("获取分布式锁成功，正在取消消息")
	defer cache.UnlockId(ctx, id)
	logrus.Debug("缓存中的数据状态为可取消状态")

	// 获取数据库中的值，并更新到缓存,必须同步，因为发送的时候需要检查缓存中信息的状态
	// 如果异步操作失败，会导致已取消的信息发送
	m.SetStatus(meta.Cancel)
	m.SetResult(meta.Fail)
	changeUpdateAndVersion(m)
	b, _ := m.Marshal()

	t := cache.NewTransaction()
	defer t.Close()

	if err = cancelStore(ctx, id, b, t); err != nil {
		t.Rollback(ctx)
		logrus.WithFields(logrus.Fields{
			"method": "cancelStore",
			"id":     id,
			"error":  err,
		}).Error("在取消发送后更新Cache时出现错误")
		return err
	}

	return t.Commit(ctx)

}

func cancelStore(ctx context.Context, id string, b []byte, t *cache.Transaction) error {
	// 更新数据库中的status
	err := t.RPushMsg(ctx, b)
	if err != nil {
		return err
	}
	return t.PutBaseCache(ctx, id, b)
}

func checkStatus(msg Messager) error {
	st := meta.Status(msg.GetStatus())
	if st == meta.Cancel {
		return errors.ErrMsgHasCancelled
	}
	if st == meta.Final {
		return errors.ErrMsgCantEdit
	}
	return nil
}

func changeUpdateAndVersion(msg Messager) {
	timeStamp := utils.NowTimeStampStr()
	newVersion := msg.GetVersion() + 1
	msg.SetVersion(newVersion)
	msg.SetUpdatedAt(timeStamp)
}
