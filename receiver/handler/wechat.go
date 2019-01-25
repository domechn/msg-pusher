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
	"github.com/domgoer/msgpusher/pkg/errors"
	"github.com/domgoer/msgpusher/pkg/pb/meta"
	"github.com/domgoer/msgpusher/pkg/utils"
	"github.com/domgoer/msgpusher/receiver/model"
	"github.com/domgoer/msgpusher/receiver/service"
)

var weChatService = service.NewWeChatServiceImpl()

// @router(POST,"/version/wechat")
// WeChatProducer 接收用户提交的json，并将json转化成消息插入到wechat消息队列
func WeChatProducer(ctx context.Context, body []byte) (res []byte, err error) {
	p := &meta.WeChatProducer{}
	if err = json.Unmarshal(body, p); err != nil {
		err = errors.ErrParam
		return
	}
	var id string
	if id, err = processData(ctx, weChatService, p); err != nil {
		return
	}

	res = model.NewResponseDataKey("id", id).MustMarshal()
	return
}

// @router(PATCH,"version/wechat")
// WeChatEdit 修改微信发送信息
func WeChatEdit(ctx context.Context, body []byte) (res []byte, err error) {
	p := &meta.WeChatProducer{}
	if err = json.Unmarshal(body, p); err != nil {
		return
	}
	if err = checkEdit(p); err != nil {
		return
	}
	if err = weChatService.Edit(ctx, p); err != nil {
		return
	}
	res = successResp
	return
}

// WeChatIDDetail 通过id获取具体信息的内容
func WeChatIDDetail(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		err = errors.ErrIDIsInvalid
		return
	}
	data, err := weChatService.Detail(ctx, id)
	if err != nil {
		return nil, err
	}
	res = model.NewResponseDataKey("detail", data).MustMarshal()

	return

}

// @router(DELETE,"version/wechat/{id}")
// WeChatCancel 取消发送微信
func WeChatCancel(ctx context.Context, d map[string]string) (res []byte, err error) {
	id := d["id"]
	if err = utils.ValidateUUIDV4(id); err != nil {
		err = errors.ErrIDIsInvalid
		return
	}
	if err = weChatService.Cancel(ctx, id); err != nil {
		return
	}
	res = successResp
	return
}
