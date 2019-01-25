/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : wechat.go
#   Created       : 2019/1/11 16:58
#   Last Modified : 2019/1/11 16:58
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"

	"github.com/domgoer/msg-pusher/pkg/pb/meta"
	"github.com/jmoiron/sqlx"
)

// WeChatCancelMsgByID 将wechat信息的发送状态设置为取消
func WeChatCancelMsgByID(ctx context.Context, id string) (*sqlx.Tx, error) {
	return update(ctx, "WeChatCancelMsgByID", `UPDATE wechats SET status=2,result_status=2 WHERE id = ?`, id)
}

// WeChatDetailByID 按照id查询wechat所有字段信息，如果未找到返回error
func WeChatDetailByID(ctx context.Context, id string) (*meta.DbWeChat, error) {
	res := &meta.DbWeChat{}
	err := query(ctx, res, "WeChatDetailbyID", `SELECT * FROM wechats WHERE id = ? LIMIT 1`, id)
	return res, err
}

// InsertWechats 将数据插入wechats表
func WeChatInsert(ctx context.Context, wechat *meta.DbWeChat) (*sqlx.Tx, error) {
	sendT := changeSendTime(wechat.SendTime)

	return insert(ctx,
		"WeChatInsert",
		`INSERT INTO wechats (id,platform,platform_key,touser,type,template,url,content,arguments,send_time) VALUES (?,?,?,?,?,?,?,?,?,?)`,
		wechat.Id,
		wechat.Platform,
		wechat.PlatformKey,
		wechat.Touser,
		wechat.Type,
		wechat.Template,
		wechat.Url,
		wechat.Content,
		wechat.Arguments,
		sendT)
}

func WeChatEdit(ctx context.Context, w *meta.DbWeChat) (*sqlx.Tx, error) {
	sendT := changeSendTime(w.SendTime)

	return update(ctx,
		"WeChatEdit",
		`UPDATE wechats SET content=?,arguments=?,send_time=?,touser=?,template=? WHERE id=? AND status=1`,
		w.Content,
		w.Arguments,
		sendT,
		w.Touser,
		w.Template,
		w.Id)
}

// WeChatUpdateSendResult 修改微信发送结果
func WeChatUpdateSendResult(ctx context.Context, w *meta.DbWeChat) (*sqlx.Tx, error) {
	return update(ctx,
		"WeChatUpdateSendResult",
		`UPDATE wechats SET try_num=?,status=?,result_status=?,reason=? WHERE id=?`,
		w.TryNum,
		w.Status,
		w.ResultStatus,
		w.Reason,
		w.Id)
}

// WeChatUpdateAndInsertBatch 批量执行修改,如果不存在就插入
func WeChatUpdateAndInsertBatch(ctx context.Context, dw []*meta.DbWeChat) error {
	var args []interface{}
	for _, w := range dw {
		args = append(args,
			w.Id,
			w.Platform,
			w.PlatformKey,
			w.Touser,
			w.Type,
			w.Template,
			w.Url,
			w.Content,
			w.Arguments,
			w.SendTime,
			w.TryNum,
			w.Status,
			w.ResultStatus,
			w.Reason,
			w.Version)
	}
	err := batch(ctx,
		"wechats",
		[]string{"id", "platform", "platform_key", "touser", "type", "template", "url", "content", "arguments", "send_time", "try_num", "status", "result_status", "reason", "version"},
		args...,
	)
	return err
}
