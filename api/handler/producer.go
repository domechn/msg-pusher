/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : producer.go
#   Created       : 2019/1/8 16:32
#   Last Modified : 2019/1/8 16:32
#   Describe      :
#
# ====================================================*/
package handler

import (
	"context"
	"encoding/json"
	"uuabc.com/sendmsg/api/service"
	"uuabc.com/sendmsg/pkg/pb/meta"
)

// SmsProducer 接收用户提交的json，并将json转化成消息插入到sms消息队列
func SmsProducer(ctx context.Context, body []byte) (res []byte, err error) {
	p := &meta.SmsProducer{}
	if err = json.Unmarshal(body, p); err != nil {
		return
	}
	if err = processData(ctx, p); err != nil {
		return
	}

	return
}

// SmsProducers 批量将用户的消息插入sms队列
func SmsProducers(ctx context.Context, body []byte) (res []byte, err error) {
	return
}

// WeChatProducer 接收用户提交的json，并将json转化成消息插入到wechat消息队列
func WeChatProducer(ctx context.Context, body []byte) (res []byte, err error) {
	p := &meta.WeChatProducer{}
	if err = json.Unmarshal(body, p); err != nil {
		return
	}
	if err = processData(ctx, p); err != nil {
		return
	}
	return
}

// EmailProducer 接收用户提交的json，并将json转化成消息插入到email消息队列
func EmailProducer(ctx context.Context, body []byte) (res []byte, err error) {
	p := &meta.EmailProducer{}
	if err = json.Unmarshal(body, p); err != nil {
		return
	}
	if err = processData(ctx, p); err != nil {
		return
	}
	return
}

func processData(ctx context.Context, p service.Meta) (err error) {
	if err = p.Validated(); err != nil {
		return
	}
	p.Transfer()
	err = service.ProducerImpl.Produce(ctx, p)
	return
}
