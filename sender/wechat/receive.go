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
	"fmt"

	"github.com/sirupsen/logrus"
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

func (r *Receiver) OnReceive(data []byte) bool {
	fmt.Printf("%s", string(data))
	return true
}
