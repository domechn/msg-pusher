/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : email.go
#   Created       : 2019/1/11 12:01
#   Last Modified : 2019/1/11 12:01
#   Describe      :
#
# ====================================================*/
package handler

import (
	"context"
	"encoding/json"
	"github.com/hellofresh/janus/pkg/errors"
	"uuabc.com/sendmsg/api/model"
	"uuabc.com/sendmsg/api/service"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/pkg/utils"
)

// EmailProducer 接收用户提交的json，并将json转化成消息插入到email消息队列
func EmailProducer(ctx context.Context, body []byte) (res []byte, err error) {
	p := &meta.EmailProducer{}
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

func EmailEdit(ctx context.Context, body []byte) (res []byte, err error) {

	res = successResp
	return
}

// @route(GET,"/version/email/{id}")
// EmailIDDetail 根据email消息的id返回msg的具体信息
func EmailIDDetail(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		err = errors.ErrInvalidID
		return
	}
	data, err := service.DetailerImpl.Detail(ctx, "email", id)
	if err != nil {
		return nil, err
	}
	res = data.MustMarshal()
	return
}

// @route(DELETE,"/version/email/{id}")
// EmailCancel 取消发送email
func EmailCancel(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		err = errors.ErrInvalidID
		return
	}
	if err = service.Canceler.Cancel(ctx, "email", id); err != nil {
		return
	}
	res = successResp
	return
}
