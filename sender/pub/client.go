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
	"github.com/domgoer/msgpusher/config"
	"github.com/domgoer/msgpusher/pkg/send/email"
	"github.com/domgoer/msgpusher/pkg/send/sms"
	"github.com/domgoer/msgpusher/pkg/send/wechat"
	"github.com/domgoer/msgpusher/storer"
)

var (
	SmsClient    *sms.Client
	WeChatClient *wechat.Client
	EmailClient  *email.Client
)

func Init() {
	aliyunConf := config.AliyunConf()
	SmsClient = sms.NewClient(sms.Config{
		AccessKeyId:  aliyunConf.AccessKeyId,
		AccessSecret: aliyunConf.AccessSecret,
		GatewayURL:   aliyunConf.GatewayURL,
	})

	weChatConf := config.WeChatConf()
	WeChatClient = wechat.NewClient(
		wechat.Config{
			APPId:     weChatConf.AppId,
			APPSecret: weChatConf.AppSecret,
		},
		storer.Cache)

	emailConf := config.EmailConf()
	EmailClient = email.NewClient(email.Config{
		ServerAddr: emailConf.ServerAddr,
		Username:   emailConf.Username,
		Password:   emailConf.Password,
		Host:       emailConf.Host,
		TLS:        emailConf.TLS,
	})

	config.MQConf()
}
