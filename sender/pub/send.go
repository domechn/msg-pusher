/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : send.go
#   Created       : 2019/1/18 14:20
#   Last Modified : 2019/1/18 14:20
#   Describe      :
#
# ====================================================*/
package pub

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/pkg/retry/backoff"
	"uuabc.com/sendmsg/pkg/utils"
	"uuabc.com/sendmsg/storer/cache"
)

type RetryFunc func(count int) error

// Send 发送消息到客户端，有重试机制，重试次数根据retryFunc里判断
func Send(id string, sendFunc RetryFunc) error {
	bk := backoff.NewServiceBackOff()
	var count int

	doFunc := func() error {
		// 分布式锁，防止资源竞争
		err := cache.LockId(context.Background(), id)
		if err != nil {
			logrus.WithField("id", id).Error("获取分布式锁失败，消息可能正在被其他线程在处理")
			// 等待一段时间，再获取
			time.Sleep(time.Millisecond * 300)
			return err
		}
		logrus.WithField("id", id).Info("获取分布式锁成功，正在发送消息")
		// 释放分布式锁
		defer cache.UnlockId(context.Background(), id)

		count++
		if err := sendFunc(count); err != nil {
			if err == ErrTooManyTimes {
				logrus.WithField("error", err).Error("发送消息失败")
				return nil
			}
			return err
		}
		return nil
	}
	return backoff.Retry(doFunc, bk)
}

// SendRetryFunc 返回一个可以用于重试发送的方法
func SendRetryFunc(msg Messager, send func(Messager) error, doList func(Cache, []byte) error) RetryFunc {
	var reason error
	return func(count int) error {
		// 发送之前检查状态,如果已发送就直接返回成功
		if bl, _ := cache.SendResult(context.Background(), msg.GetId()); bl {
			logrus.Info("消息已处理，无需重复发送")
			return nil
		}

		if count > TryNum {
			msg.SetStatus(int32(meta.Status_Final))
			msg.SetResult(int32(meta.Result_Fail))
			msg.SetTryNum(TryNum)
			if reason != nil {
				msg.SetReason(reason.Error())
			}
			updateCache(msg, doList)
			return ErrTooManyTimes
		}
		err := send(msg)
		defer func() { reason = err }()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"id":    msg.GetId(),
				"error": err,
			}).Errorf("消息发送失败，正在重新尝试第%d次", count)
			return ErrMsgSendFailed
		}

		msg.SetStatus(int32(meta.Status_Final))
		msg.SetResult(int32(meta.Result_Success))
		msg.SetTryNum(int32(count))
		// 更新数据库和缓存，如果出错打印日志，不做错误处理
		err = updateCache(msg, doList)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"id":    msg.GetId(),
				"data":  msg,
				"error": err,
			}).Errorf("消息发送成功但是更新缓存和数据库时发生错误，请手动修改")
		}

		logrus.WithField("id", msg.GetId()).Info("消息成功发送")
		return nil
	}
}

func updateCache(msg Messager, doList func(Cache, []byte) error) error {
	// 消息已发送，后续不会再修改，所以这里版本号直接更新一个大数字
	newVersion := msg.GetVersion() + 99
	msg.SetVersion(newVersion)
	// 更新修改时间
	msg.SetUpdatedAt(utils.NowTimeStampStr())
	b, _ := msg.Marshal()
	t := cache.NewTransaction()
	defer t.Close()
	doList(t, b)
	// 强制更新缓存
	t.PutBaseCache(context.Background(), msg.GetId(), b)
	t.PutSendSuccess(context.Background(), msg.GetId())
	return t.Commit()
}
