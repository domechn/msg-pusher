/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : receiver.go
#   Created       : 2019/1/8 19:20
#   Last Modified : 2019/1/8 19:20
#   Describe      :
#
# ====================================================*/
package mq

type Receiver interface {
	QueueName() string     // 获取接收者需要监听的队列
	RouterKey() string     // 这个队列绑定的路由
	OnError(error)         // 处理遇到的错误，当RabbitMQ对象发生了错误，他需要告诉接收者处理错误
	OnReceive([]byte) bool // 处理收到的消息, 这里需要告知RabbitMQ对象消息是否处理成功
}
