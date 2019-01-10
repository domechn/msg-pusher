/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : send.go
#   Created       : 2019-01-07 14:20:59
#   Last Modified : 2019-01-07 14:20:59
#   Describe      :
#
# ====================================================*/
package mail

import (
	"fmt"
	"net/smtp"
	"strings"
	"uuabc.com/sendmsg/send"
)

type Client struct {
	cfg  *Config
	auth smtp.Auth
}

func NewClient(cfg *Config) send.Sender {
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	return &Client{
		cfg:  cfg,
		auth: auth,
	}
}

// Send 发送邮件，msg用mail.NewRequest(...)生成
// do参数不做处理
func (c *Client) Send(msg send.Message, do send.DoRes) error {
	if c.cfg == nil {
		return fmt.Errorf("sender initializes the exception, use the NewSender() method to initialize it")
	}
	return smtp.SendMail(c.cfg.ServerAddr, c.auth, c.cfg.Username, strings.Split(msg.To(), ";"), msg.Content())
}
