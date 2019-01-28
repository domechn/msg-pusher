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
	"github.com/hiruok/msg-pusher/storer"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/streadway/amqp"
)

func MsgProduce(ctx context.Context, msg []byte, delay int64) error {
	return produce(ctx, "msg", msg, delay)
}

func produce(ctx context.Context, typeName string, msg []byte, delay int64) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan("MqProduce", opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "rabbit-mq")
		span.SetTag("mq.queuename", typeName)
		span.SetTag("mq.msg", string(msg))
		span.SetTag("mq.delay", delay)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

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
