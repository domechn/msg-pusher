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

var msgService = service.NewMsgServiceImpl()

// @router(POST,"/msg")
// MsgProducer 接收用户提交的json，并将json转化成消息插入到消息队列
func MsgProducer(ctx context.Context, body []byte) (res []byte, err error) {
	p := &meta.MsgProducer{}
	if err = json.Unmarshal(body, p); err != nil {
		err = errors.ErrParam
		return
	}
	var id string

	if err = p.Validated(); err != nil {
		return
	}
	p.Transfer(true)

	if id, err = msgService.Produce(ctx, p); err != nil {
		return
	}

	res = model.NewResponseDataKey("id", id).MustMarshal()
	return
}

// MsgIDDetail 根据id或取消消息的具体信息
// @router(GET,"/msg/{id}")
func MsgIDDetail(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		err = errors.ErrIDIsInvalid
		return
	}
	data, err := msgService.Detail(ctx, id)
	if err != nil {
		return nil, err
	}
	res = model.NewResponseDataKey("detail", data).MustMarshal()

	return
}

// @router(GET,"/msg/key/{key}/page/{p}")
func MsgDetailByKey(ctx context.Context, d map[string]string) (res []byte, err error) {
	p := d["p"]
	key := d["key"]
	page, err := strconv.Atoi(p)
	if err != nil || key == "" {
		return nil, errors.ErrParam
	}
	m, err := msgService.DetailByKeyAndPage(ctx, key, page)
	res = model.NewResponseDataKey("detail", m).MustMarshal()
	return res, err
}

// @router(GET,"/msg/to/{to}/page/{p}")
func MsgDetailByTo(ctx context.Context, d map[string]string) (res []byte, err error) {
	mobile := d["to"]
	p := d["p"]
	if err = checkMobileDetail(mobile, p); err != nil {
		return
	}
	pg, _ := strconv.Atoi(p)
	data, err := msgService.DetailByToAndPage(ctx, mobile, pg)
	if err != nil {
		return
	}
	res = model.NewResponseDataKey("detail", data).MustMarshal()
	return
}

// @router(PATCH,"/msg")
// SmsEdit 修改短信发送消息
func MsgEdit(ctx context.Context, body []byte) (res []byte, err error) {
	p := &meta.MsgProducer{}
	if err = json.Unmarshal(body, p); err != nil {
		err = errors.ErrParam
		return
	}
	if err = checkEdit(p); err != nil {
		return
	}
	if err = msgService.Edit(ctx, p); err != nil {
		return
	}
	res = successResp
	return
}

// @router(DELETE,"/msg/{id}")
// MsgCancel 取消发送短信
func MsgCancel(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		return
	}
	if err = msgService.Cancel(ctx, id); err != nil {
		return
	}
	res = successResp
	return
}

// @router(DELETE,"/msg/key/{key}")
func MsgCancelByKey(ctx context.Context, d map[string]string) (res []byte, err error) {
	key := d["key"]
	if key == "" {
		return
	}
	ids, err := msgService.WaitingMsgIdByPlat(ctx, key)
	if err != nil {
		return nil, err
	}
	fails := msgService.CancelBatch(ctx, ids)
	if len(fails) == 0 {
		res = successResp
	} else {
		res = model.NewResponseDataKey("fail", fails).MustMarshal()
	}
	return
}
