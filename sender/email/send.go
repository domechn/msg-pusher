/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : send.go
#   Created       : 2019/1/16 16:44
#   Last Modified : 2019/1/16 16:44
#   Describe      :
#
# ====================================================*/
package email

import (
	"context"

	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/pkg/send/email"
	"uuabc.com/sendmsg/sender/pub"
	"uuabc.com/sendmsg/storer/cache"
	"uuabc.com/sendmsg/storer/db"
)

// check 验证data是否符合要求，如果符合要求会返回nil，并且按照data转化的id将数据赋值给msg
func (r *Receiver) check(data []byte, msg pub.Messager) (err error) {
	id := string(data)
	logrus.WithField("type", r.queueName).Info("开始验证消息的有效性")
	defer logrus.WithField("type", r.queueName).Infof("消息验证结束,err: %v", err)
	err = pub.Check(id, msg)
	return
}

func (r *Receiver) send(msg pub.Messager) pub.RetryFunc {
	// 失败原因
	var reason error
	return func(count int) error {
		lockKey := "send_" + msg.GetId()
		emailMsg := msg.(*meta.DbEmail)
		// 发送之前检查状态,如果已发送就直接返回成功
		if bl, _ := cache.SendResult(context.Background(), msg.GetId()); bl {
			logrus.Info("消息已处理，无需重复发送")
			return nil
		}
		// 分布式锁，防止资源竞争
		err := cache.LockID5s(context.Background(), lockKey)
		defer func() { reason = err }()
		if err != nil {
			logrus.WithField("id", emailMsg.Id).Error("获取分布式锁失败，消息可能正在被其他线程在处理")
			return nil
		}
		logrus.WithField("id", emailMsg.Id).Info("获取分布式锁成功，正在发送消息")
		// 释放分布式锁
		defer cache.ReleaseLock(context.Background(), lockKey)
		if count > pub.TryNum {
			emailMsg.SetStatus(int32(meta.Status_Final))
			emailMsg.SetResult(int32(meta.Result_Fail))
			emailMsg.SetTryNum(pub.TryNum)
			if reason != nil {
				emailMsg.Reason = reason.Error()
			}
			updateDbAndCache(emailMsg)
			return pub.ErrTooManyTimes
		}

		err = pub.EmailClient.Send(email.NewMessage(
			emailMsg.Destination,
			emailMsg.Title,
			emailMsg.Content,
		), nil)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("邮件发送失败，error: %v", err)
			return pub.ErrMsgSendFailed
		}

		emailMsg.SetStatus(int32(meta.Status_Final))
		emailMsg.SetResult(int32(meta.Result_Success))
		emailMsg.SetTryNum(int32(count))
		// 更新数据库和缓存，如果出错打印日志，不做错误处理
		err = updateDbAndCache(emailMsg)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"id":    emailMsg.Id,
				"data":  emailMsg,
				"error": err,
			}).Errorf("消息发送成功但是更新缓存和数据库时发生错误，请手动修改")
		}

		logrus.WithField("id", emailMsg.Id).Info("消息成功发送")
		return nil
	}
}

// 更新数据库和缓存
func updateDbAndCache(msg *meta.DbEmail) error {
	var err error
	tx, err := db.EmailUpdateSendResult(context.Background(), msg)
	if err != nil {
		db.RollBack(tx)
		return err
	}
	b, _ := msg.Marshal()
	// 强制更新缓存
	err = cache.PutBaseCacheForce(context.Background(), msg.Id, b)
	cache.PutLastestCache(context.Background(), msg.Id, b)
	cache.PutSendSuccess(context.Background(), msg.Id)
	if err != nil {
		db.RollBack(tx)
		return err
	}

	return db.Commit(tx)
}
