/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : sms.go
#   Created       : 2019/1/11 12:01
#   Last Modified : 2019/1/11 12:01
#   Describe      : 处理和短信操作有关的数据
#
# ====================================================*/
package handler

import (
	"context"
	"encoding/json"
	"strconv"
	"uuabc.com/sendmsg/api/model"
	"uuabc.com/sendmsg/api/service"
	"uuabc.com/sendmsg/pkg/errors"
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
	var id string
	if id, err = processData(ctx, p); err != nil {
		return
	}

	res = model.NewResponseDataKey("id", id).MustMarshal()
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
		err = errors.ErrIDIsInvalid
		return
	}
	data, err := service.DetailerImpl.Detail(ctx, "sms", id)
	if err != nil {
		return nil, err
	}
	res = model.NewResponseDataKey("detail", data).MustMarshal()

	return
}

// @router(GET,"/sms/key/{key}")
func SmsKeyDetail(ctx context.Context, body []byte) (res []byte, err error) {
	return
}

// @router(GET,"/sms/mobile/{mobile}/page/{p}")
func SmsMobileDetail(ctx context.Context, d map[string]string) (res []byte, err error) {
	mobile := d["mobile"]
	p := d["p"]
	if err = checkMobileDetail(mobile, p); err != nil {
		return
	}
	pg, _ := strconv.Atoi(p)
	data, err := service.DetailerImpl.DetailByPhonePage(ctx, mobile, pg)
	if err != nil {
		return
	}
	res = model.NewResponseDataKey("detail", data).MustMarshal()
	return
}

func checkMobileDetail(mobile, p string) error {
	if !utils.ValidatePhone(mobile) {
		return errors.ErrPhoneNumber
	}
	pg, err := strconv.Atoi(p)
	if err != nil || pg < 1 || pg > 10 {
		return errors.ErrPageInvalidate
	}
	return nil
}

// @router(PATCH,"/sms")
func SmsEdit(ctx context.Context, body []byte) (res []byte, err error) {
	res = successResp
	return
}

// @router(DELETE,"/sms/{id}")
// SmsCancel 取消发送短信
func SmsCancel(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		return
	}
	if err = service.Canceler.Cancel(ctx, "sms", id); err != nil {
		return
	}
	res = successResp
	return
}

// @router(DELETE,"/sms/key/{key}")
func SmsKeyCancel(ctx context.Context, d map[string]string) (res []byte, err error) {
	res = successResp
	return
}
