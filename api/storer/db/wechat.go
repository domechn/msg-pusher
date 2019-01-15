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
	"uuabc.com/sendmsg/api/model"
	"uuabc.com/sendmsg/api/storer"
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
func WeChatDetailByID(ctx context.Context, id string) (*model.DbWeChat, error) {
	res := &model.DbWeChat{}
	err := storer.DB.GetContext(ctx, res, "SELECT * FROM wechats WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// InsertWechats 将数据插入wechats表
func WeChatInsert(ctx context.Context, wechat *model.DbWeChat) (*sqlx.Tx, error) {
	tx, err := storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO wechats (id,platform,touser,type,template,url,content,send_time) VALUES (?,?,?,?,?,?,?,?)`)
	if err != nil {
		return tx, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(
		ctx,
		wechat.ID,
		wechat.Platform,
		wechat.Touser,
		wechat.Type,
		wechat.Template,
		wechat.URL,
		wechat.Content,
		wechat.SendTime,
	)
	if err != nil {
		return tx, err
	}
	return tx, nil
}

func WeChatEdit(ctx context.Context, w *model.DbWeChat) (*sqlx.Tx, error) {
	tx, err := storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	query := "UPDATE wechats SET content=?,send_time=? "
	if w.Touser != "" {
		query += ",touser=? WHERE id=?"
	} else {
		query += "WHERE id=?"
	}
	query += " AND status=1"
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return tx, err
	}
	defer stmt.Close()
	var res sql.Result
	if w.Touser != "" {
		res, err = stmt.ExecContext(ctx, w.Content, w.SendTime, w.Touser, w.ID)
	} else {
		res, err = stmt.ExecContext(ctx, w.Content, w.SendTime, w.ID)
	}
	if err != nil {
		return tx, err
	}
	if i, _ := res.RowsAffected(); i == 0 {
		return tx, ErrNoRowsEffected
	}
	return tx, nil
}
