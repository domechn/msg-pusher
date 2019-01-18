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
	"database/sql"

	"github.com/jmoiron/sqlx"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/storer"
)

// WeChatCancelMsgByID 将wechat信息的发送状态设置为取消
func WeChatCancelMsgByID(ctx context.Context, id string) (*sqlx.Tx, error) {
	tx, err := storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	stmt, err := storer.DB.PrepareContext(ctx, "UPDATE wechats SET status=2,result_status=2 WHERE id = ?")
	if err != nil {
		return tx, err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return tx, err
	}
	if i, _ := res.RowsAffected(); i == 0 {
		return tx, ErrNoRowsEffected
	}
	return tx, nil
}

// WeChatDetailByID 按照id查询wechat所有字段信息，如果未找到返回error
func WeChatDetailByID(ctx context.Context, id string) (*meta.DbWeChat, error) {
	res := &meta.DbWeChat{}
	err := storer.DB.GetContext(ctx, res, "SELECT * FROM wechats WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// InsertWechats 将数据插入wechats表
func WeChatInsert(ctx context.Context, wechat *meta.DbWeChat) (*sqlx.Tx, error) {
	tx, err := storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO wechats (id,platform,touser,type,template,url,content,arguments,send_time) VALUES (?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return tx, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(
		ctx,
		wechat.Id,
		wechat.Platform,
		wechat.Touser,
		wechat.Type,
		wechat.Template,
		wechat.Url,
		wechat.Content,
		wechat.Arguments,
		changeSendTime(wechat.SendTime),
	)
	if err != nil {
		return tx, err
	}
	return tx, nil
}

func WeChatEdit(ctx context.Context, w *meta.DbWeChat) (*sqlx.Tx, error) {
	tx, err := storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	query := "UPDATE wechats SET arguments=?,send_time=?,touser=?,template=? WHERE id=? AND status=1"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return tx, err
	}
	defer stmt.Close()
	var res sql.Result
	sendT := changeSendTime(w.SendTime)
	res, err = stmt.ExecContext(ctx, w.Arguments, sendT, w.Touser, w.Template, w.Id)
	if err != nil {
		return tx, err
	}
	if i, _ := res.RowsAffected(); i == 0 {
		return tx, ErrNoRowsEffected
	}
	return tx, nil
}

// WeChatUpdateSendResult 修改微信发送结果
func WeChatUpdateSendResult(ctx context.Context, w *meta.DbWeChat) (*sqlx.Tx, error) {
	tx, err := storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, "UPDATE wechats SET try_num=?,status=?,result_status=?,reason=? WHERE id=?")
	if err != nil {
		return tx, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, w.TryNum, w.Status, w.ResultStatus, w.Reason, w.Id)
	return tx, err
}
