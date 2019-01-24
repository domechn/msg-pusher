/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : putter.go
#   Created       : 2019/1/11 17:03
#   Last Modified : 2019/1/11 17:03
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/storer"
)

// EmailCancelMsgByID 将email信息的发送状态设置为取消
func EmailCancelMsgByID(ctx context.Context, id string) (*sqlx.Tx, error) {
	return update(ctx,
		"EmailCancelMsgByID",
		`UPDATE emails SET status=2,result_status=2 WHERE id=?`,
		id)
}

// EmailDetailByID 按照id查询email所有字段信息，如果未找到返回error
func EmailDetailByID(ctx context.Context, id string) (*meta.DbEmail, error) {
	res := &meta.DbEmail{}
	err := query(ctx, res, "EmailDetailByID", `SELECT * FROM emails WHERE id = ? LIMIT 1`, id)

	return res, err
}

// EmailInsert 将消息插入emails表
func EmailInsert(ctx context.Context, e *meta.DbEmail) (*sqlx.Tx, error) {
	return insert(ctx,
		"EmailInsert",
		`INSERT INTO emails (id,platform,platform_key,title,content,destination,type,template,arguments,server,send_time) VALUES (?,?,?,?,?,?,?,?,?,?,?)`,
		e.Id,
		e.Platform,
		e.PlatformKey,
		e.Title,
		e.Content,
		e.Destination,
		e.Type,
		e.Template,
		e.Arguments,
		e.Server,
		changeSendTime(e.SendTime))
}

// EmailEdit 修改email发送信息的内容
func EmailEdit(ctx context.Context, e *meta.DbEmail) (*sqlx.Tx, error) {
	sendT := changeSendTime(e.SendTime)
	return update(ctx,
		"EmailEdit",
		`UPDATE emails SET destination=?,template=?,content=?,arguments=?,send_time=? WHERE id=? AND status=1`,
		e.Destination,
		e.Template,
		e.Content,
		e.Arguments,
		sendT,
		e.Id)
}

// EmailUpdateSendResult 修改短信发送结果
func EmailUpdateSendResult(ctx context.Context, e *meta.DbEmail) (*sqlx.Tx, error) {
	return update(ctx,
		"EmailUpdateSendResult",
		`UPDATE emails SET try_num=?,status=?,result_status=?,reason=? WHERE id=?`,
		e.TryNum,
		e.Status,
		e.ResultStatus,
		e.Reason,
		e.Id)
}

// EmailUpdateBatch 批量执行修改,如果不存在就插入
func EmailUpdateAndInsertBatch(ctx context.Context, es []*meta.DbEmail) error {
	var args []interface{}
	for _, e := range es {
		args = append(args,
			e.Id,
			e.Platform,
			e.PlatformKey,
			e.Title,
			e.Content,
			e.Destination,
			e.Type,
			e.Template,
			e.Arguments,
			e.Server,
			e.SendTime,
			e.TryNum,
			e.Status,
			e.ResultStatus,
			e.Reason)
	}
	err := batch(ctx,
		"emails",
		[]string{"id", "platform", "platform_key", "title", "content", "destination", "type", "template", "arguments", "server", "send_time", "try_num", "status", "result_status", "reason"},
		args...,
	)
	return err
}

// EmailUpdateBatch 批量执行修改,如果不存在就插入,性能不如EmailUpdateAndInsertBatch
// Deprecated: use EmailUpdateAndInsertBatch replace.
func EmailUpdateAndInsertExecBatch(ctx context.Context, es []*meta.DbEmail) error {
	sql := `INSERT INTO emails (id,platform,platform_key,title,content,destination,type,template,arguments,server,send_time,try_num,status,result_status,reason) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)  ON DUPLICATE KEY UPDATE platform=VALUES(platform),platform_key=VALUES(platform_key),title=VALUES(title),content=VALUES(content),destination=VALUES(destination),type=VALUES(type),template=VALUES(template),arguments=VALUES(arguments),server=VALUES(server),send_time=VALUES(send_time),try_num=VALUES(try_num),status=VALUES(status),result_status=VALUES(result_status),reason=VALUES(reason)`
	tx, _ := storer.DB.Beginx()
	stmt, _ := tx.Preparex(sql)
	defer stmt.Close()
	for _, e := range es {
		stmt.Exec(
			e.Id,
			e.Platform,
			e.PlatformKey,
			e.Title,
			e.Content,
			e.Destination,
			e.Type,
			e.Template,
			e.Arguments,
			e.Server,
			e.SendTime,
			e.TryNum,
			e.Status,
			e.ResultStatus,
			e.Reason,
		)
	}
	return tx.Commit()
}
