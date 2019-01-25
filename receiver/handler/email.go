/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : putter.go
#   Created       : 2019/1/11 12:01
#   Last Modified : 2019/1/11 12:01
#   Describe      :
#
# ====================================================*/
package handler

import (
	"context"

	"github.com/domgoer/msg-pusher/pkg/errors"
	"github.com/domgoer/msg-pusher/pkg/pb/meta"
	"github.com/domgoer/msg-pusher/pkg/utils"
	"github.com/domgoer/msg-pusher/receiver/model"
	"github.com/domgoer/msg-pusher/receiver/service"
)

var emailService = service.NewEmailSeriveImpl()

// @router(POST,"/version/email")
// EmailProducer 接收用户提交的json，并将json转化成消息插入到email消息队列
func EmailProducer(ctx context.Context, body []byte) (res []byte, err error) {
	p := &meta.EmailProducer{}
	if err = json.Unmarshal(body, p); err != nil {
		err = errors.ErrParam
		return
	}
	var id string
	if id, err = processData(ctx, emailService, p); err != nil {
		return
	}

	res = model.NewResponseDataKey("id", id).MustMarshal()
	return
}

// @router(PATCH,"/version/email")
// EmailEdit 修改email发送的信息
func EmailEdit(ctx context.Context, body []byte) (res []byte, err error) {
	p := &meta.EmailProducer{}
	if err = json.Unmarshal(body, p); err != nil {
		err = errors.ErrParam
		return
	}
	if err = checkEdit(p); err != nil {
		return
	}
	if err = emailService.Edit(ctx, p); err != nil {
		return
	}
	res = successResp
	return
}

// @route(GET,"/version/email/{id}")
// EmailIDDetail 根据email消息的id返回msg的具体信息
func EmailIDDetail(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		err = errors.ErrIDIsInvalid
		return
	}
	data, err := emailService.Detail(ctx, id)
	if err != nil {
		return nil, err
	}
	res = model.NewResponseDataKey("detail", data).MustMarshal()
	return
}

// @route(DELETE,"/version/email/{id}")
// EmailCancel 取消发送email
func EmailCancel(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		err = errors.ErrIDIsInvalid
		return
	}
	if err = emailService.Cancel(ctx, id); err != nil {
		return
	}
	res = successResp
	return
}
