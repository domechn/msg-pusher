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

// DbSms 数据库中smss表的实体类
type DbSms struct {
	ID           string `json:"id" db:"id"`
	Platform     string `json:"platform" db:"platform"`
	PlatformKey  string `json:"platform_key" db:"platform_key"`
	Content      string `json:"content" db:"content"`
	Mobile       string `json:"mobile" db:"mobile"`
	Type         int    `json:"type" db:"type"`
	Template     string `json:"template" db:"template"`
	Arguments    string `json:"arguments" db:"arguments"`
	Server       int    `json:"server" db:"server"`
	SendTime     string `json:"send_time" db:"send_time"`
	TryNum       int    `json:"try_num" db:"try_num"`
	Status       int    `json:"status" db:"status"`
	ResultStatus int    `json:"result_status" db:"result_status"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}

// DbWeChat 数据库中wechats表的实体类
type DbWeChat struct {
	ID           string `json:"id" db:"id"`
	Platform     string `json:"platform" db:"platform"`
	Touser       string `json:"touser" db:"touser"`
	Type         int    `json:"type" db:"type"`
	Template     string `json:"template" db:"template"`
	URL          string `json:"url" db:"url"`
	Content      string `json:"content" db:"content"`
	SendTime     string `json:"send_time" db:"send_time"`
	TryNum       int    `json:"try_num" db:"try_num"`
	Status       int    `json:"status" db:"status"`
	ResultStatus int    `json:"result_status" db:"result_status"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}

// DbEmail 数据库中emails表的实体类
type DbEmail struct {
	ID           string `json:"id" db:"id"`
	Platform     string `json:"platform" db:"platform"`
	PlatformKey  string `json:"platform_key" db:"platform_key"`
	Title        string `json:"title" db:"title"`
	Content      string `json:"content" db:"content"`
	Destination  string `json:"destination" db:"destination"`
	Type         string `json:"type" db:"type"`
	Template     string `json:"template" db:"template"`
	Arguments    string `json:"arguments" db:"arguments"`
	Server       int    `json:"server" db:"server"`
	SendTime     string `json:"send_time" db:"send_time"`
	TryNum       int    `json:"try_num" db:"try_num"`
	Status       int    `json:"status" db:"status"`
	ResultStatus int    `json:"result_status" db:"result_status"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}
