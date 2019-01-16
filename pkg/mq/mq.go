/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : mq.go
#   Created       : 2019/1/8 19:20
#   Last Modified : 2019/1/8 19:20
#   Describe      :
#
# ====================================================*/
package mq

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitConn struct {
	addr string
	conn *amqp.Connection
}

type RabbitMQ struct {
	wg sync.WaitGroup

	channel      *amqp.Channel
	exchangeName string     // exchange的名称
	exchangeType string     // exchange的类型
	exchangeArgs amqp.Table // exchange的额外参数
	receivers    []Receiver
}

// New 按地址新建连接到mq
func New(addr string) (*RabbitConn, error) {
	r := new(RabbitConn)
	var err error
	r.conn, err = amqp.Dial(addr)
	if err != nil {
		return nil, err
	}
	r.addr = addr
	return r, nil
}

// New 初始化实例
func (c *RabbitConn) RabbitMQ(exName, exType string, exArgs amqp.Table) (*RabbitMQ, error) {
	var err error
	r := new(RabbitMQ)
	r.channel, err = c.conn.Channel()
	if err != nil {
		return r, err
	}
	r.exchangeName = exName
	r.exchangeType = exType
	r.exchangeArgs = exArgs
	return r, nil
}

func (c *RabbitConn) Connection() *amqp.Connection {
	return c.conn
}

// Close 关闭连接
func (c *RabbitConn) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// RegisterReceiver 注册一个用于接收指定队列指定路由的数据接收者
func (r *RabbitMQ) RegisterReceiver(receiver Receiver) {
	r.receivers = append(r.receivers, receiver)
}

func (r *RabbitMQ) Publish(ctx context.Context, routeKey string, publishing amqp.Publishing) error {
	return r.channel.Publish(
		r.exchangeName,
		routeKey,
		false,
		false,
		publishing,
	)
}

// prepareExchange 准备rabbitmq的Exchange
func (r *RabbitMQ) prepareExchange() error {
	return r.channel.ExchangeDeclare(
		r.exchangeName, // exchange
		r.exchangeType, // type
		true,           // durable
		false,          // autoDelete
		false,          // internal
		false,          // nowait
		r.exchangeArgs, // args
	)
}

// ReConnect 重新连接channel
func (r *RabbitMQ) ReConnect(conn *amqp.Connection) error {
	var err error
	r.Close()
	r.channel, err = conn.Channel()
	if err != nil {
		return err
	}
	return nil
}

// Close 关闭conn和channel
func (r *RabbitMQ) Close() (err error) {
	if r.channel != nil {
		err = r.channel.Close()
	}
	return err
}

func (r *RabbitMQ) listen(receiver Receiver) {
	defer r.wg.Done()
	// 这里获取每个接收者需要监听的队列和路由
	queueName := receiver.QueueName()
	routerKey := receiver.RouterKey()
	// 申明Queue
	_, err := r.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive(排他性队列)
		false,     // no-wait
		nil,       // arguments
	)
	if nil != err {
		// 当队列初始化失败的时候，需要告诉这个接收者相应的错误
		receiver.OnError(fmt.Errorf("初始化队列 %s 失败: %s", queueName, err.Error()))
	}

	// 将Queue绑定到Exchange上去
	err = r.channel.QueueBind(
		queueName,      // queue name
		routerKey,      // routing key
		r.exchangeName, // exchange
		false,          // no-wait
		nil,
	)
	if nil != err {
		receiver.OnError(fmt.Errorf("绑定队列 [%s - %s] 到交换机失败: %s", queueName, routerKey, err.Error()))
	}

	// 获取消费通道
	r.channel.Qos(
		10, // prefetchCount mq不要同时给一个消费者推送多余N个消息，即一旦有N个消息没有ack，该consumer会block掉，直到有消息ack
		0,
		true, // global 如果为true，prefetchCount和prefetchSize的限制级别为channel，否则为consumer
	) // 确保rabbitmq会一个一个发消息
	msgs, err := r.channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if nil != err {
		receiver.OnError(fmt.Errorf("获取队列 %s 的消费通道失败: %s", queueName, err.Error()))
	}

	// 使用callback消费数据
	for msg := range msgs {
		go func(mg amqp.Delivery) {
			// 当接收者消息处理失败的时候，
			// 比如网络问题导致的数据库连接失败，redis连接失败等等这种
			// 通过重试可以成功的操作，那么这个时候是需要重试的
			// 直到数据处理成功后再返回，然后才会回复rabbitmq ack
			for !receiver.OnReceive(mg.Body) {
				time.Sleep(1 * time.Second)
			}
			// 确认收到本条消息, multiple必须为false
			mg.Ack(false)
		}(msg)
	}
}

// run 各个消费者监听
func (r *RabbitMQ) run() error {
	if len(r.receivers) == 0 {
		return nil
	}
	for _, rec := range r.receivers {
		r.wg.Add(1)
		go r.listen(rec)
	}
	r.wg.Wait()

	logrus.Error("所有处理queue的任务都退出了")
	return fmt.Errorf("all queues quit accidently")
}

// Start 启动消费者
func (r *RabbitMQ) Start(conn *amqp.Connection) {
	for {
		err := r.run()
		// 非异常退出
		if err == nil {
			time.Sleep(time.Second * 3)
			continue
		}
		// 断开连接
		r.Close()
		time.Sleep(time.Second * 3)
		// 重新连接
		if err := r.ReConnect(conn); err != nil {
			logrus.Errorf("mq连接已关闭，无法重新建立连接")
			return
		}
	}
}
