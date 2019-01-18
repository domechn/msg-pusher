/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : receive.go
#   Created       : 2019/1/16 15:43
#   Last Modified : 2019/1/16 15:43
#   Describe      :	从mq中接收信息
#
# ====================================================*/
package wechat

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
		queueName: "wechat",
		routerKey: "wechat",
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
	dw := &meta.DbWeChat{}
	if err := r.check(data, dw); err != nil {
		return
	}
	if err := pub.Send(dw.Id, pub.SendRetryFunc(dw, r.send, r.doDB)); err != nil {
		logrus.WithFields(logrus.Fields{
			"type":  r.queueName,
			"id":    dw.Id,
			"data":  dw,
			"error": err,
		}).Error("发送失败")
	}
	return
}
