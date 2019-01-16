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
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/pkg/retry/backoff"
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

type RetryFunc func(count int) error

// Send 发送消息到客户端，有重试机制，重试次数根据retryFunc里判断
func Send(sendFunc RetryFunc) error {
	bk := backoff.NewServiceBackOff()
	var count int
	doFunc := func() error {
		count++
		if err := sendFunc(count); err != nil {
			if err == ErrTooManyTimes {
				logrus.WithField("error", err).Error("发送消息失败")
			}
			return err
		}
		return nil
	}
	return backoff.Retry(doFunc, bk)
}
