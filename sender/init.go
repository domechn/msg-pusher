/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : init.go
#   Created       : 2019/1/16 15:44
#   Last Modified : 2019/1/16 15:44
#   Describe      :
#
# ====================================================*/
package sender

import (
	"github.com/streadway/amqp"
	"uuabc.com/sendmsg/config"
	"uuabc.com/sendmsg/sender/email"
	"uuabc.com/sendmsg/sender/pub"
	"uuabc.com/sendmsg/sender/sms"
	"uuabc.com/sendmsg/sender/wechat"
	"uuabc.com/sendmsg/storer"
)

func Init() error {
	if err := storer.Init(); err != nil {
		return err
	}
	pub.Init()
	return nil
}

func Start() error {
	conn, err := storer.MqCli.RabbitMQ(config.MQConf().ExChangeName, "x-delayed-message", amqp.Table{
		"x-delayed-type": "direct",
	})
	if err != nil {
		return err
	}
	conn.RegisterReceiver(wechat.NewReceiver())
	conn.RegisterReceiver(email.NewReceiver())
	conn.RegisterReceiver(sms.NewReceiver())
	conn.Start(storer.MqCli.Connection())
	return nil
}

func Close() error {
	return storer.Close()
}
