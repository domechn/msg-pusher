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
	"uuabc.com/sendmsg/api/model"
	"uuabc.com/sendmsg/api/storer"
)

// SmsDetailByID 按照id查询sms所有字段信息，如果未找到返回error
func SmsDetailByID(id string) (*model.DbSms, error) {
	res := &model.DbSms{}
	err := storer.DB.Get(res, "SELECT * FROM smss WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// WeChatDetailByID 按照id查询wechat所有字段信息，如果未找到返回error
func WeChatDetailByID(id string) (*model.DbWeChat, error) {
	res := &model.DbWeChat{}
	err := storer.DB.Get(res, "SELECT * FROM wechats WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// EmailDetailByID 按照id查询email所有字段信息，如果未找到返回error
func EmailDetailByID(id string) (*model.DbEmail, error) {
	res := &model.DbEmail{}
	err := storer.DB.Get(res, "SELECT * FROM emails WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
