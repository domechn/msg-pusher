/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : public.go
#   Created       : 2019/1/16 17:20
#   Last Modified : 2019/1/16 17:20
#   Describe      :
#
# ====================================================*/
package pub

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/pkg/utils"
	"uuabc.com/sendmsg/storer/cache"
)

var (
	// 消息已经被处理
	ErrMsgHasDealed = errors.New("msg has been dealed")
	// 重试多次后仍然失败
	ErrTooManyTimes = errors.New("retry for many times, still failed")
	// 消息已过期
	ErrMsgIsExpiration = errors.New("msg is expiration")
	// 消息发送时间未到
	ErrMsgDeliveryNotArrived = errors.New("the msg delivery time has not arrived")

	ErrMsgSendFailed = errors.New("msg send failed")

	DefaultExpirationTime = time.Minute * 15
)

const (
	TryNum = 3
)

func Check(id string, msg Messager) error {
	if err := checkId(id); err != nil {
		return err
	}
	if err := checkStatus(id, msg); err != nil {
		return err
	}
	if err := checkStatus(id, msg); err != nil {
		return err
	}
	return nil
}

func checkId(id string) error {
	if err := utils.ValidateUUIDV4(id); err != nil {
		logrus.WithFields(logrus.Fields{
			"data": id,
		}).Error("从mq中接收到的数据不符合要求")
		return err
	}
	return nil
}

// CheckSendTime 检查发送时间，若时间未到返回ErrMsgDeliveryNotArrived（消息未到发送时间）
// 如果当前时间超过发送时间15分钟返回ErrMsgIsExpiration（消息过期）
func checkSendTime(msg Messager) error {
	now := time.Now().UTC()
	sendT, err := time.Parse("2006-01-02T15:04:05Z", msg.GetSendTime())
	if err != nil {
		return err
	}
	if sendT.Add(DefaultExpirationTime).Before(now) {
		logrus.WithFields(logrus.Fields{
			"now":       now.String(),
			"send_time": sendT.String(),
		}).Error("消息已过期")
		return ErrMsgIsExpiration
	}
	if sendT.After(now) {
		logrus.WithFields(logrus.Fields{
			"now":       now.String(),
			"send_time": sendT.String(),
		}).Error("消息未到发送时间")
		return ErrMsgDeliveryNotArrived
	}
	return nil
}

func checkStatus(id string, msg Messager) error {
	b, err := cache.BaseDetail(context.Background(), id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"data":  id,
			"error": err,
		}).Error("从cache中获取数据失败")
		return err
	}
	if err := msg.Unmarshal(b); err != nil {
		logrus.WithFields(logrus.Fields{
			"data":  string(b),
			"error": err,
		}).Error("数据转码异常")
		return err
	}
	status := msg.GetStatus()
	if status != int32(meta.Status_Wait) {
		logrus.WithFields(logrus.Fields{
			"data":   string(b),
			"status": meta.Status_name[status],
			"error":  err,
		}).Error("数据状态不是待发送状态，可能已被处理")
		return ErrMsgHasDealed
	}
	return nil
}
