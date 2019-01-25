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

	"github.com/domgoer/msgpusher/pkg/pb/meta"
	"github.com/domgoer/msgpusher/sender/pub"
	"github.com/sirupsen/logrus"
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
	}).Fatal("初始化消费队列失败")
}

func (r *Receiver) OnReceive(data []byte) (res bool) {
	// 防止立即发送的数据还没有存入缓存中
	time.Sleep(time.Millisecond * 300)
	res = true
	ds := &meta.DbSms{}
	if err := r.check(data, ds); err != nil {
		return
	}
	if err := pub.Send(ds.Id, pub.SendRetryFunc(ds, r.send, r.doList)); err != nil {
		logrus.WithFields(logrus.Fields{
			"type":  r.queueName,
			"id":    ds.Id,
			"data":  ds,
			"error": err,
		}).Error("发送失败")
	}
	return
}
