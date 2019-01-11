/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : wechat.go
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

func WeChatEdit(ctx context.Context, body []byte) (res []byte, err error) {
	return
}

func WeChatIDDetail(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		return
	}
	data, err := service.DetailerImpl.IDDetail(ctx, "wechat", id)
	if err != nil {
		return nil, err
	}
	res, err = json.Marshal(data)
	return

}

func WeChatCancel(ctx context.Context, body []byte) (res []byte, err error) {
	return
}
