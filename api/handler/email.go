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
	if err = processData(ctx, p); err != nil {
		return
	}
	return
}

func EmailEdit(ctx context.Context, body []byte) (res []byte, err error) {
	return
}

// @route(GET,"/version/email/{id}")
// EmailIDDetail 根据email消息的id返回msg的具体信息
func EmailIDDetail(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		return
	}
	data, err := service.DetailerImpl.IDDetail(ctx, "email", id)
	if err != nil {
		return nil, err
	}
	res, err = json.Marshal(data)
	return
}

func EmailCancel(ctx context.Context, body []byte) (res []byte, err error) {
	return
}
