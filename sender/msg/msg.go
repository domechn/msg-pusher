/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : msg.go
#   Created       : 2019/1/28 11:46
#   Last Modified : 2019/1/28 11:46
#   Describe      :
#
# ====================================================*/
package msg

import (
	"context"
	"fmt"
	"github.com/domgoer/msg-pusher/pkg/pb/meta"
	"github.com/domgoer/msg-pusher/pkg/send"
	"github.com/domgoer/msg-pusher/pkg/send/email"
	"github.com/domgoer/msg-pusher/pkg/send/sms"
	"github.com/domgoer/msg-pusher/pkg/send/wechat"
	"github.com/domgoer/msg-pusher/sender/pub"
	"github.com/sirupsen/logrus"
	"time"
)

type Receiver struct {
	queueName string
	routerKey string
}

func NewReceiver() *Receiver {
	return &Receiver{
		queueName: "msg",
		routerKey: "msg",
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
	dm := &meta.DbMsg{}
	if err := r.check(data, dm); err != nil {
		return
	}
	if err := pub.Send(dm.Id, pub.SendRetryFunc(dm, r.send)); err != nil {
		logrus.WithFields(logrus.Fields{
			"type":  r.queueName,
			"id":    dm.Id,
			"data":  dm,
			"error": err,
		}).Error("发送失败")
	}
	return
}

func (r *Receiver) doList(c pub.Cache, b []byte) error {
	return c.RPush(context.Background(), b)
}

// check 验证data是否符合要求，如果符合要求会返回nil，并且按照data转化的id将数据赋值给msg
func (r *Receiver) check(data []byte, msg pub.Messager) (err error) {
	id := string(data)
	logrus.WithField("type", r.queueName).Info("开始验证消息的有效性")
	err = pub.Check(id, msg)
	logrus.WithField("type", r.queueName).Infof("消息验证结束,err: %v", err)
	return
}

func (r *Receiver) send(msg pub.Messager) error {
	m := msg.(*meta.DbMsg)
	var client send.Sender
	var sendMsg send.Message
	sendTo := m.GetSendTo()
	switch m.Type {
	case meta.Sms:
		client = pub.SmsClient(m.Server)
		sendMsg = sms.NewRequest(sendTo, m.Reserved, m.Template, m.Arguments, "12345")
	case meta.WeChat:
		client = pub.WeChatClient()
		sendMsg = wechat.NewRequest(sendTo, m.Template, m.Reserved, []byte(m.Content))
	case meta.Email:
		client = pub.EmailClient(m.Server)
		sendMsg = email.NewMessage(sendTo, m.Reserved, m.Content)
	}

	if client == nil || sendMsg == nil {
		return fmt.Errorf("send: cannot find corresponding sending client, msg server not support")
	}

	return client.Send(sendMsg, nil)
}
