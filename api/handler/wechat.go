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
	"github.com/hellofresh/janus/pkg/errors"
	"uuabc.com/sendmsg/api/model"
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
	var id string
	if id, err = processData(ctx, p); err != nil {
		return
	}

	res = model.NewResponseDataKey("id", id).MustMarshal()
	return
}

func WeChatEdit(ctx context.Context, body []byte) (res []byte, err error) {
	return
}

// WeChatIDDetail 通过id获取具体信息的内容
func WeChatIDDetail(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		err = errors.ErrInvalidID
		return
	}
	data, err := service.DetailerImpl.Detail(ctx, "wechat", id)
	if err != nil {
		return nil, err
	}
	res = data.MustMarshal()
	return

}

// @router(DELETE,"/sms/{id}")
// WeChatCancel 取消发送微信
func WeChatCancel(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		err = errors.ErrInvalidID
		return
	}
	if err = service.Canceler.Cancel(ctx, "wechat", id); err != nil {
		return
	}
	res = successResp
	return
}
