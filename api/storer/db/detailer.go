/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : detailer.go
#   Created       : 2019/1/10 20:04
#   Last Modified : 2019/1/10 20:04
#   Describe      : 按照id在数据库中查找到对应数据的所有字段的信息，如果没有找到对应数据则返回error
#
# ====================================================*/
package db

import (
	"context"
	"uuabc.com/sendmsg/api/model"
	"uuabc.com/sendmsg/api/storer"
)

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

// WeChatDetailByID 按照id查询wechat所有字段信息，如果未找到返回error
func WeChatDetailByID(ctx context.Context, id string) (*model.DbWeChat, error) {
	res := &model.DbWeChat{}
	err := storer.DB.GetContext(ctx, res, "SELECT * FROM wechats WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// EmailDetailByID 按照id查询email所有字段信息，如果未找到返回error
func EmailDetailByID(ctx context.Context, id string) (*model.DbEmail, error) {
	res := &model.DbEmail{}
	err := storer.DB.GetContext(ctx, res, "SELECT * FROM emails WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
