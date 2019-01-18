/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : receive.go
#   Created       : 2019/1/16 16:00
#   Last Modified : 2019/1/16 16:00
#   Describe      :
#
# ====================================================*/
package email

import (
	"time"

	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/sender/pub"
)

type Receiver struct {
	queueName string
	routerKey string
}

func NewReceiver() *Receiver {
	return &Receiver{
		queueName: "email",
		routerKey: "email",
	}
}

func (r *Receiver) QueueName() string {
	return r.queueName
}

func (r *Receiver) RouterKey() string {
	return r.routerKey
}

func (r *Receiver) OnError(err error) {
	logrus.WithFields(logrus.Fields{
		"Queue": r.queueName,
		"error": err,
	}).Error("初始化消费队列失败")
}

func (r *Receiver) OnReceive(data []byte) (res bool) {
	// 防止立即发送的数据还没有存入缓存中
	time.Sleep(time.Millisecond * 300)
	res = true
	de := &meta.DbEmail{}
	if err := r.check(data, de); err != nil {
		return
	}
	if err := pub.Send(de.Id, pub.SendRetryFunc(de, r.send, r.doDB)); err != nil {
		logrus.WithFields(logrus.Fields{
			"type":  r.queueName,
			"id":    de.Id,
			"data":  de,
			"error": err,
		}).Error("发送失败")
	}
	return
}
