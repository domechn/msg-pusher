/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : producer.go
#   Created       : 2019/1/9 17:21
#   Last Modified : 2019/1/9 17:21
#   Describe      :
#
# ====================================================*/
package mq

import (
	"context"
	"github.com/streadway/amqp"
	"uuabc.com/sendmsg/api/storer"
)

func ProduceSms(ctx context.Context, msg []byte, delay int32) error {
	return produce(ctx, "sms", msg, delay)
}

func ProduceEmail(ctx context.Context, msg []byte, delay int32) error {
	return produce(ctx, "email", msg, delay)
}

func ProduceWeChat(ctx context.Context, msg []byte, delay int32) error {
	return produce(ctx, "wechat", msg, delay)
}

func produce(ctx context.Context, typeName string, msg []byte, delay int32) error {
	channel, err := storer.MqCli.RabbitMQ(storer.ExChangeName, "x-delayed-message", amqp.Table{
		"x-delayed-type": "direct",
	})
	if err != nil {
		return err
	}
	defer channel.Close()
	return channel.Publish(ctx, typeName, amqp.Publishing{
		ContentType: "application/json",
		Body:        msg,
		Headers: amqp.Table{
			"x-delay": delay,
		},
	})
}
