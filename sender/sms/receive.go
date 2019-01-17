/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : receive.go
#   Created       : 2019/1/16 15:58
#   Last Modified : 2019/1/16 15:58
#   Describe      :
#
# ====================================================*/
package sms

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
		queueName: "sms",
		routerKey: "sms",
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
	ds := &meta.DbSms{}
	if err := r.check(data, ds); err != nil {
		return
	}
	pub.Send(r.send(ds))
	return
}
