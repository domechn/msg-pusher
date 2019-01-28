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
	"github.com/hiruok/msg-pusher/config"
	"github.com/hiruok/msg-pusher/pkg/mq"
	"github.com/hiruok/msg-pusher/sender/msg"
	"github.com/hiruok/msg-pusher/sender/pub"
	"github.com/hiruok/msg-pusher/storer"
	"github.com/streadway/amqp"
)

func Init() error {
	if err := storer.Init(); err != nil {
		return err
	}
	pub.Init()
	return nil
}

func Start() error {
	stopC := make(chan struct{})
	start(msg.NewReceiver())
	<-stopC
	return nil
}

func Close() error {
	return storer.Close()
}

func start(r mq.Receiver) error {
	conn, err := storer.MqCli.RabbitMQ(config.MQConf().ExChangeName, "x-delayed-message", amqp.Table{
		"x-delayed-type": "direct",
	})
	if err != nil {
		return err
	}
	conn.RegisterReceiver(r)
	go conn.Start(storer.MqCli.Connection())
	return nil
}
