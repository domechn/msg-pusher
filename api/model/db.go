/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : db.go
#   Created       : 2019/1/11 10:09
#   Last Modified : 2019/1/11 10:09
#   Describe      :
#
# ====================================================*/
package model

import (
	"github.com/json-iterator/go"
	"uuabc.com/sendmsg/pkg/pb/meta"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// DbSms 数据库中smss表的实体类
type DbSms struct {
	ID           string `json:"id" db:"id"`
	Platform     int32  `json:"platform" db:"platform"`
	PlatformKey  string `json:"platform_key" db:"platform_key"`
	Content      string `json:"content" db:"content"`
	Mobile       string `json:"mobile" db:"mobile"`
	Type         int32  `json:"type" db:"type"`
	Template     string `json:"template" db:"template"`
	Arguments    string `json:"arguments" db:"arguments"`
	Server       int32  `json:"server" db:"server"`
	SendTime     string `json:"send_time" db:"send_time"`
	TryNum       int32  `json:"try_num" db:"try_num"`
	Status       int32  `json:"status" db:"status"`
	ResultStatus int32  `json:"result_status" db:"result_status"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}

func (d *DbSms) SetStatus(s int32) {
	d.Status = s
}

func (d *DbSms) SetResult(s int32) {
	d.ResultStatus = s
}

func (d *DbSms) Marshal() ([]byte, error) {
	return json.Marshal(d)
}

func (d *DbSms) Unmarshal(b []byte) error {
	return json.Unmarshal(b, d)
}

func (d *DbSms) GetStatus() meta.Status {
	return meta.Status(d.Status)
}

// DbWeChat 数据库中wechats表的实体类
type DbWeChat struct {
	ID           string `json:"id" db:"id"`
	Platform     int32  `json:"platform" db:"platform"`
	Touser       string `json:"touser" db:"touser"`
	Type         int32  `json:"type" db:"type"`
	Template     string `json:"template" db:"template"`
	URL          string `json:"url" db:"url"`
	Content      string `json:"content" db:"content"`
	SendTime     string `json:"send_time" db:"send_time"`
	TryNum       int32  `json:"try_num" db:"try_num"`
	Status       int32  `json:"status" db:"status"`
	ResultStatus int32  `json:"result_status" db:"result_status"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}

func (d *DbWeChat) SetStatus(s int32) {
	d.Status = s
}

func (d *DbWeChat) SetResult(s int32) {
	d.ResultStatus = s
}

func (d *DbWeChat) Marshal() ([]byte, error) {
	return json.Marshal(d)
}

func (d *DbWeChat) Unmarshal(b []byte) error {
	return json.Unmarshal(b, d)
}

func (d *DbWeChat) GetStatus() meta.Status {
	return meta.Status(d.Status)
}

// DbEmail 数据库中emails表的实体类
type DbEmail struct {
	ID           string `json:"id" db:"id"`
	Platform     int32  `json:"platform" db:"platform"`
	PlatformKey  string `json:"platform_key" db:"platform_key"`
	Title        string `json:"title" db:"title"`
	Content      string `json:"content" db:"content"`
	Destination  string `json:"destination" db:"destination"`
	Type         int32  `json:"type" db:"type"`
	Template     string `json:"template" db:"template"`
	Arguments    string `json:"arguments" db:"arguments"`
	Server       int32  `json:"server" db:"server"`
	SendTime     string `json:"send_time" db:"send_time"`
	TryNum       int32  `json:"try_num" db:"try_num"`
	Status       int32  `json:"status" db:"status"`
	ResultStatus int32  `json:"result_status" db:"result_status"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}

func (d *DbEmail) Marshal() ([]byte, error) {
	return json.Marshal(d)
}

func (d *DbEmail) Unmarshal(b []byte) error {
	return json.Unmarshal(b, d)
}

func (d *DbEmail) GetStatus() meta.Status {
	return meta.Status(d.Status)
}

func (d *DbEmail) SetStatus(s int32) {
	d.Status = s
}

func (d *DbEmail) SetResult(s int32) {
	d.ResultStatus = s
}
