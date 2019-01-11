/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : producer.go
#   Created       : 2019/1/10 11:52
#   Last Modified : 2019/1/10 11:52
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"uuabc.com/sendmsg/api/storer"
	"uuabc.com/sendmsg/pkg/pb/meta"
)

// InsertSmss 将数据插入smss表
func InsertSmss(ctx context.Context, sms *meta.SmsProducer) error {
	stmt, err := storer.DB.PrepareContext(ctx, `INSERT INTO smss (id,platform,content,mobile,template,arguments,send_time,server,type) VALUES (?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(
		ctx,
		sms.Id,
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
		return err
	}
	return nil
}

// InsertWechats 将数据插入wechats表
func InsertWechats(ctx context.Context, wechat *meta.WeChatProducer) error {
	stmt, err := storer.DB.PrepareContext(ctx, `INSERT INTO wechats (id,platform,touser,type,template,url,content,send_time) VALUES (?,?,?,?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(
		ctx,
		wechat.Id,
		wechat.Platform,
		wechat.Touser,
		wechat.Type,
		wechat.TemplateID,
		wechat.Url,
		wechat.Data,
		wechat.SendTime,
	)
	if err != nil {
		return err
	}
	return nil
}

// InsertEmails 将消息插入emails表
func InsertEmails(ctx context.Context, email *meta.EmailProducer) error {
	stmt, err := storer.DB.PrepareContext(ctx, `INSERT INTO emails (id,platform,platform_key,title,content,destination,type,template,arguments,server,send_time) VALUES (?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(
		ctx,
		email.Id,
		email.Platform,
		email.PlatformKey,
		email.Title,
		email.Content,
		email.Destination,
		email.Type,
		email.Template,
		email.Arguments,
		email.Server,
		email.SendTime,
	)
	if err != nil {
		return err
	}
	return nil
}
