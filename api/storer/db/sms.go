/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : sms.go
#   Created       : 2019/1/11 17:02
#   Last Modified : 2019/1/11 17:02
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

// SmsCancelMsgByID 将sms信息的发送状态设置为取消
func SmsCancelMsgByID(ctx context.Context, id string) (*sqlx.Tx, error) {
	tx, err := storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, "UPDATE smss SET status=2,result_status=2 WHERE id = ?")
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

// SmsDetailByID 按照id查询sms所有字段信息，如果未找到返回error
func SmsDetailByID(ctx context.Context, id string) (*model.DbSms, error) {
	res := &model.DbSms{}
	err := storer.DB.GetContext(ctx, res, `SELECT * FROM smss WHERE id = ? LIMIT 1`, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SmsDetailByPhoneAndPage(ctx context.Context, mobile string, page int) ([]*model.DbSms, error) {
	var res []*model.DbSms
	size := (page - 1) * 10
	err := storer.DB.SelectContext(ctx, &res, `SELECT * FROM smss WHERE mobile=? LIMIT ?,10`, mobile, size)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SmsInsert 将数据插入smss表
func SmsInsert(ctx context.Context, sms *model.DbSms) (*sqlx.Tx, error) {
	tx, err := storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO smss (id,platform,content,mobile,template,arguments,send_time,server,type) VALUES (?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return tx, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(
		ctx,
		sms.ID,
		sms.Platform,
		sms.Content,
		sms.Mobile,
		sms.Template,
		sms.Arguments,
		sms.SendTime,
		sms.Server,
		sms.Type,
	)
	if err != nil {
		return tx, err
	}
	return tx, nil
}

func SmsEdit(ctx context.Context, s *model.DbSms) (*sqlx.Tx, error) {
	tx, err := storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	query := "UPDATE smss SET content=?,send_time=? "
	if s.Mobile != "" {
		query += ",mobile=? WHERE id=?"
	} else {
		query += "WHERE id=?"
	}
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return tx, err
	}
	defer stmt.Close()
	var res sql.Result
	if s.Mobile != "" {
		res, err = stmt.ExecContext(ctx, s.Content, s.SendTime, s.Mobile, s.ID)
	} else {
		res, err = stmt.ExecContext(ctx, s.Content, s.SendTime, s.ID)
	}
	if err != nil {
		return tx, err
	}
	if i, _ := res.RowsAffected(); i == 0 {
		return tx, ErrNoRowsEffected
	}
	return tx, nil
}
