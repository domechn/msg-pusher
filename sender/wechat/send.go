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
package wechat

import (
	"context"

	"github.com/domgoer/msgpusher/pkg/pb/meta"
	"github.com/domgoer/msgpusher/pkg/send/wechat"
	"github.com/domgoer/msgpusher/sender/pub"
	"github.com/sirupsen/logrus"
)

func (r *Receiver) check(data []byte, msg pub.Messager) (err error) {
	id := string(data)
	logrus.WithField("type", r.queueName).Info("开始验证消息的有效性")
	err = pub.Check(id, msg)
	logrus.WithField("type", r.queueName).Infof("消息验证结束,err: %v", err)
	return
}

func (r *Receiver) send(msg pub.Messager) error {
	weChatMsg := msg.(*meta.DbWeChat)
	return pub.WeChatClient.Send(wechat.NewRequest(
		weChatMsg.Touser,
		weChatMsg.Template,
		weChatMsg.Url,
		[]byte(weChatMsg.Content),
	), nil)
}

func (r *Receiver) doList(c pub.Cache, b []byte) error {
	return c.RPushWeChat(context.Background(), b)
}
