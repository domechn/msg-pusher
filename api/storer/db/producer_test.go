/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : producer_test.go
#   Created       : 2019/1/10 13:00
#   Last Modified : 2019/1/10 13:00
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"github.com/satori/go.uuid"
	"testing"
	"uuabc.com/sendmsg/api/storer"
	"uuabc.com/sendmsg/pkg/db"
	"uuabc.com/sendmsg/pkg/pb/meta"
)

func init() {
	storer.DB, _ = db.New(db.Config{
		URL: "root:root@tcp(localhost:3306)/uuabc?charset=utf8&parseTime=True",
	})
}

func TestInsertSmss(t *testing.T) {
	err := InsertSmss(
		context.Background(),
		&meta.SmsProducer{
			Id:        uuid.NewV4().String(),
			Content:   "testcontent",
			Mobile:    "13155555555",
			Template:  "sms-13123123",
			Arguments: "testarg",
			SendTime:  "2019-01-01 01:01:11",
			Server:    2,
			Type:      2,
			Platform:  2,
		},
	)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertWechats(t *testing.T) {
	err := InsertWechats(
		context.Background(),
		&meta.WeChatProducer{
			Id:         uuid.NewV4().String(),
			Data:       "testcontent",
			Touser:     "13155555555",
			TemplateID: "sms-13123123",
			Url:        "",
			SendTime:   "2019-01-01 01:01:11",
			Type:       1,
			Platform:   1,
		},
	)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertEmails(t *testing.T) {
	err := InsertEmails(
		context.Background(),
		&meta.EmailProducer{
			Id:          uuid.NewV4().String(),
			PlatformKey: "123",
			Server:      1,
			Title:       "test",
			Content:     "test",
			Template:    "hello",
			Arguments:   "123test",
			Destination: "abc@uuabc.com",
			SendTime:    "2019-01-01 01:01:11",
			Type:        1,
			Platform:    1,
		},
	)
	if err != nil {
		t.Error(err)
	}
}
