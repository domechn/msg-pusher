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
	"strconv"

	"github.com/domgoer/msg-pusher/pkg/errors"
	"github.com/domgoer/msg-pusher/pkg/pb/meta"
	"github.com/domgoer/msg-pusher/pkg/utils"
	"github.com/domgoer/msg-pusher/receiver/model"
	"github.com/domgoer/msg-pusher/receiver/service"
)

var smsService = service.NewSmsServiceImpl()

// @router(POST,"/sms")
// SmsProducer 接收用户提交的json，并将json转化成消息插入到sms消息队列
func SmsProducer(ctx context.Context, body []byte) (res []byte, err error) {
	p := &meta.SmsProducer{}
	if err = json.Unmarshal(body, p); err != nil {
		err = errors.ErrParam
		return
	}
	var id string
	if id, err = processData(ctx, smsService, p); err != nil {
		return
	}

	res = model.NewResponseDataKey("id", id).MustMarshal()
	return
}

// @router(POST,"/smss")
// SmsProducers 批量将用户的消息插入sms队列
func SmsProducers(ctx context.Context, body []byte) (res []byte, err error) {
	p := &meta.SmsProducers{}
	if err = json.Unmarshal(body, p); err != nil {
		err = errors.ErrParam
		return
	}
	// 检验参数
	for _, d := range p.Data {
		d.Platform = p.Platform
		if err = d.Validated(); err != nil {
			return
		}
		d.Transfer(true)
	}
	data, sErr := smsService.ProduceBatch(ctx, p.Data)
	if sErr != nil {
		err = sErr
		return
	}
	res = model.NewResponseData(data).MustMarshal()
	return
}

// @router(GET,"/sms/{id}")
func SmsIDDetail(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		err = errors.ErrIDIsInvalid
		return
	}
	data, err := smsService.Detail(ctx, id)
	if err != nil {
		return nil, err
	}
	res = model.NewResponseDataKey("detail", data).MustMarshal()

	return
}

// @router(GET,"/sms/plat/{plat}/key/{key}")
func SmsDetailByPlat(ctx context.Context, d map[string]string) (res []byte, err error) {
	plat := d["plat"]
	platform, err := strconv.Atoi(plat)
	if err != nil {
		return nil, errors.ErrParam
	}
	key := d["key"]
	m, err := smsService.DetailByPlat(ctx, int32(platform), key)
	res = model.NewResponseDataKey("detail", m).MustMarshal()
	return res, err
}

// @router(GET,"/sms/mobile/{mobile}/page/{p}")
func SmsMobileDetail(ctx context.Context, d map[string]string) (res []byte, err error) {
	mobile := d["mobile"]
	p := d["p"]
	if err = checkMobileDetail(mobile, p); err != nil {
		return
	}
	pg, _ := strconv.Atoi(p)
	data, err := smsService.DetailByPhonePage(ctx, mobile, pg)
	if err != nil {
		return
	}
	res = model.NewResponseDataKey("detail", data).MustMarshal()
	return
}

// @router(PATCH,"/sms")
// SmsEdit 修改短信发送消息
func SmsEdit(ctx context.Context, body []byte) (res []byte, err error) {
	p := &meta.SmsProducer{}
	if err = json.Unmarshal(body, p); err != nil {
		err = errors.ErrParam
		return
	}
	if err = checkEdit(p); err != nil {
		return
	}
	if err = smsService.Edit(ctx, p); err != nil {
		return
	}
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
	if err = smsService.Cancel(ctx, id); err != nil {
		return
	}
	res = successResp
	return
}

// @router(DELETE,"/sms/plat/{plat}/key/{key}")
func SmsCancelByPlat(ctx context.Context, d map[string]string) (res []byte, err error) {
	plat := d["plat"]
	p, err := strconv.Atoi(plat)

	if err != nil {
		return nil, errors.ErrParam
	}
	ids, err := smsService.WaitSmsIdByPlat(ctx, int32(p), d["key"])
	if err != nil {
		return nil, err
	}
	fails := smsService.CancelBatch(ctx, ids)
	if len(fails) == 0 {
		res = successResp
	} else {
		res = model.NewResponseDataKey("fail", fails).MustMarshal()
	}
	return
}
