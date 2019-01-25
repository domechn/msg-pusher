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
package sms

import (
	"context"

	"github.com/domgoer/msgpusher/pkg/pb/meta"
	"github.com/domgoer/msgpusher/pkg/send/sms"
	"github.com/domgoer/msgpusher/sender/pub"
	"github.com/sirupsen/logrus"
)

// check 验证data是否符合要求，如果符合要求会返回nil，并且按照data转化的id将数据赋值给msg
func (r *Receiver) check(data []byte, msg pub.Messager) (err error) {
	id := string(data)
	logrus.WithField("type", r.queueName).Info("开始验证消息的有效性")
	err = pub.Check(id, msg)
	logrus.WithField("type", r.queueName).Infof("消息验证结束,err: %v", err)
	return
}

func (r *Receiver) send(msg pub.Messager) error {
	smsMsg := msg.(*meta.DbSms)
	return pub.SmsClient.Send(sms.NewRequest(
		smsMsg.Mobile,
		"UUabc",
		smsMsg.Template,
		smsMsg.Arguments,
		"12345",
	), nil)
}

func (r *Receiver) doList(c pub.Cache, b []byte) error {
	return c.RPushSms(context.Background(), b)
}
