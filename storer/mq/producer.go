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
	"uuabc.com/sendmsg/storer"
)

func SmsProduce(ctx context.Context, msg []byte, delay int64) error {
	return produce(ctx, "sms", msg, delay)
}

func EmailProduce(ctx context.Context, msg []byte, delay int64) error {
	return produce(ctx, "email", msg, delay)
}

func WeChatProduce(ctx context.Context, msg []byte, delay int64) error {
	return produce(ctx, "wechat", msg, delay)
}

// TODO 修改err类型
func produce(ctx context.Context, typeName string, msg []byte, delay int64) error {
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
