/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : wechat_test.go
#   Created       : 2019/1/14 10:24
#   Last Modified : 2019/1/14 10:24
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
	"uuabc.com/sendmsg/api/model"
	"uuabc.com/sendmsg/api/storer"
	"uuabc.com/sendmsg/pkg/db"
)

func init() {
	storer.DB, _ = db.New(db.Config{
		URL: "root:root@tcp(localhost:3306)/uuabc?charset=utf8&parseTime=True",
	})
}

func TestInsertWechats(t *testing.T) {
	tx, err := WeChatInsert(
		context.Background(),
		&model.DbWeChat{
			ID:       uuid.NewV4().String(),
			Content:  "testcontent",
			Touser:   "13155555555",
			Template: "sms-13123123",
			URL:      "",
			SendTime: "2019-01-01 01:01:11",
			Type:     1,
			Platform: 1,
		},
	)
	if err != nil {
		if tx != nil {
			if err := tx.Rollback(); err != nil {
				t.Error(err)
			}
		}
		t.Error(err)
	} else {
		if err := tx.Commit(); err != nil {
			t.Error(err)
		}
	}
}

func TestWeChatDetailByID(t *testing.T) {
	res, err := WeChatDetailByID(context.Background(), "00a87366-c607-43d8-9673-6f2cd143273c")
	fmt.Println(res, err)
}
