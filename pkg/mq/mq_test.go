/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : mq_test.go
#   Created       : 2019/1/9 11:14
#   Last Modified : 2019/1/9 11:14
#   Describe      :
#
# ====================================================*/
package mq

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"testing"
	"time"
)

func TestRabbitMQ_Start(t *testing.T) {
	c, err := New("amqp://guest:guest@127.0.0.1:5672/test")
	if err != nil {
		t.Error(err)
	}
	r, err := c.RabbitMQ("exchange_test", "direct", amqp.Table{
		"x-delayed-type": "direct",
	})
	if err != nil {
		t.Error(err)
	}
	r.ReConnect(c.conn)
	time.Sleep(time.Second * 20)
	if err := r.Close(); err != nil {
		t.Error(err)
	}
	c.Close()
	fmt.Println(r)
}

func TestRabbitMQ_Publish(t *testing.T) {
	c, err := New("amqp://test:test@localhost:5672/test")
	if err != nil {
		t.Error(err)
		return
	}
	r, err := c.RabbitMQ("exchange_test", "x-delayed-message", amqp.Table{
		"x-delayed-type": "direct",
	})
	if err != nil {
		t.Error(err)
		return
	}
	p := amqp.Publishing{
		Headers: amqp.Table{
			"x-delay": int64(0),
		},
		ContentType: "text/plain",
		Body:        []byte(`{"code":200,"msg":"hello"}`),
	}
	err = r.Publish(context.Background(), "email", p)
	if err != nil {
		t.Error(err)
		return
	}
	if err := r.Close(); err != nil {
		t.Error(err)
		return
	}
	c.Close()
}
