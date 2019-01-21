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

	"github.com/jmoiron/sqlx"
	"uuabc.com/sendmsg/pkg/pb/meta"
)

// SmsCancelMsgByID 将sms信息的发送状态设置为取消
func SmsCancelMsgByID(ctx context.Context, id string) (*sqlx.Tx, error) {
	return update(ctx,
		"SmsCancelMsgByID",
		`UPDATE smss SET status=2,result_status=2 WHERE id = ?`,
		id)
}

// SmsDetailByID 按照id查询sms所有字段信息，如果未找到返回error
func SmsDetailByID(ctx context.Context, id string) (*meta.DbSms, error) {
	res := &meta.DbSms{}
	err := query(ctx, res, "SmsDetailByID", `SELECT * FROM smss WHERE id = ? LIMIT 1`, id)
	return res, err
}

func SmsDetailByPhoneAndPage(ctx context.Context, mobile string, page int) ([]*meta.DbSms, error) {
	size := (page - 1) * 10

	var res []*meta.DbSms
	err := list(ctx, &res, "SmsDetailByPhoneAndPage", `SELECT * FROM smss WHERE mobile=? LIMIT ?,10`, mobile, size)
	return res, err
}

// SmsInsert 将数据插入smss表
func SmsInsert(ctx context.Context, sms *meta.DbSms) (*sqlx.Tx, error) {
	return insert(ctx,
		"SmsInsert",
		`INSERT INTO smss (id,platform,platform_key,content,mobile,template,arguments,send_time,server,type) VALUES (?,?,?,?,?,?,?,?,?,?)`,
		sms.Id,
		sms.Platform,
		sms.PlatformKey,
		sms.Content,
		sms.Mobile,
		sms.Template,
		sms.Arguments,
		changeSendTime(sms.SendTime),
		sms.Server,
		sms.Type)
}

func SmsEdit(ctx context.Context, s *meta.DbSms) (*sqlx.Tx, error) {
	sendT := changeSendTime(s.SendTime)

	return update(ctx,
		"SmsEdit",
		`UPDATE smss SET arguments=?,send_time=?,template=?,mobile=? WHERE id=? AND status=1`,
		s.Arguments,
		sendT,
		s.Template,
		s.Mobile,
		s.Id)
}

// SmsUpdateSendResult 更新短信发送结果
func SmsUpdateSendResult(ctx context.Context, s *meta.DbSms) (*sqlx.Tx, error) {
	return update(ctx,
		"SmsUpdateSendResult",
		`UPDATE smss SET try_num=?,status=?,result_status=?,reason=? WHERE id=?`,
		s.TryNum,
		s.Status,
		s.ResultStatus,
		s.Reason,
		s.Id)
}
