/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : sms.go
#   Created       : 2019/1/11 12:01
#   Last Modified : 2019/1/11 12:01
#   Describe      :
#
# ====================================================*/
package handler

import (
	"context"
	"encoding/json"
	"uuabc.com/sendmsg/api/service"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/pkg/utils"
)

// @router(POST,"/sms")
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

// @router(POST,"/smss")
// SmsProducers 批量将用户的消息插入sms队列
func SmsProducers(ctx context.Context, body []byte) (res []byte, err error) {
	return
}

// @router(GET,"/sms/{id}")
func SmsIDDetail(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		return
	}
	data, err := service.DetailerImpl.IDDetail(ctx, "sms", id)
	if err != nil {
		return nil, err
	}
	res, err = json.Marshal(data)
	return
}

// @router(GET,"/sms/key/{key}")
func SmsKeyDetail(ctx context.Context, body []byte) (res []byte, err error) {
	return
}

// @router(GET,"/sms/mobile/{mobile}/page/{p}")
func SmsMobileDetail(ctx context.Context, d map[string]string) (res []byte, err error) {
	return
}

// @router(PATCH,"/sms")
func SmsEdit(ctx context.Context, body []byte) (res []byte, err error) {
	return
}

// @router(DELETE,"/sms/{id}")
func SmsCancel(ctx context.Context, body []byte) (res []byte, err error) {
	// q := &model.CancelReq{}
	// if err = json.Unmarshal(body, q); err != nil {
	// 	return
	// }
	id := string(body)
	if err = service.Canceler.Cancel(id); err != nil {
		return
	}
	res = []byte{}
	return
}

// @router(DELETE,"/sms")
func SmsKeyCancel(ctx context.Context, body []byte) (res []byte, err error) {
	return
}
