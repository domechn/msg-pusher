/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : client.go
#   Created       : 2019/1/16 19:19
#   Last Modified : 2019/1/16 19:19
#   Describe      :
#
# ====================================================*/
package pub

import (
	"github.com/domgoer/msg-pusher/config"
	"github.com/domgoer/msg-pusher/pkg/pb/meta"
	"github.com/domgoer/msg-pusher/pkg/send"
	"github.com/domgoer/msg-pusher/pkg/send/email"
	"github.com/domgoer/msg-pusher/pkg/send/sms"
	"github.com/domgoer/msg-pusher/pkg/send/wechat"
	"github.com/domgoer/msg-pusher/storer"
)

var (
	clientMap = make(map[string]send.Sender)
)

func Init() {
	smsConf := config.SmsConf()
	for k, v := range smsConf {
		clientMap["sms-"+k] = sms.NewClient(sms.Config{
			AccessKeyId:  v.AccessKeyId,
			AccessSecret: v.AccessSecret,
			GatewayURL:   v.GatewayURL,
		})
	}

	weChatConf := config.WeChatConf()
	if weChatConf != nil {
		clientMap["wechat"] = wechat.NewClient(
			wechat.Config{
				APPId:     weChatConf.AppId,
				APPSecret: weChatConf.AppSecret,
			},
			storer.Cache)
	}

	emailConf := config.EmailConf()
	for k, v := range emailConf {
		clientMap["email-"+k] = email.NewClient(email.Config{
			ServerAddr: v.Addr,
			Username:   v.Username,
			Password:   v.Password,
			Host:       v.Host,
			TLS:        v.TLS,
		})
	}
}

// WeChatClient 获取微信客户端
func WeChatClient() send.Sender {
	return clientMap["wechat"]
}

// SmsClient 根据服务商获取sms客户端发送消息
func SmsClient(s meta.Server) send.Sender {
	return clientMap["sms-"+meta.Server_name[int32(s)]]
}

// EmailClient 根据服务商获取email客户端发送邮件
func EmailClient(s meta.Server) send.Sender {
	return clientMap["email-"+meta.Server_name[int32(s)]]
}
