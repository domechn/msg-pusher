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
	"testing"

	"github.com/satori/go.uuid"
	"uuabc.com/sendmsg/pkg/db"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/storer"
)

func init() {
	storer.DB, _ = db.New(db.Config{
		URL: "root:root@tcp(localhost:3306)/uuabc?charset=utf8&parseTime=True",
	})
}

func TestInsertWechats(t *testing.T) {
	tx, err := WeChatInsert(
		context.Background(),
		&meta.DbWeChat{
			Id:       uuid.NewV4().String(),
			Content:  "testcontent",
			Touser:   "13155555555",
			Template: "sms-13123123",
			Url:      "",
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

func TestWeChatEdit(t *testing.T) {
	tx, err := WeChatEdit(context.Background(), &meta.DbWeChat{
		Id:        "db84e690-fbf6-4e4e-b113-44275282c6fd",
		Arguments: "{\"code\":800}",
		SendTime:  "2018-08-08 08:08:08",
	})
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		t.Error(err)
	} else {
		tx.Commit()
	}
}

func TestWeChatEditToUser(t *testing.T) {
	tx, err := WeChatEdit(context.Background(), &meta.DbWeChat{
		Id:        "db84e690-fbf6-4e4e-b113-44275282c6fd",
		Arguments: "{\"code\":800}",
		SendTime:  "2018-08-08 08:08:08",
		Touser:    "me",
	})
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		t.Error(err)
	} else {
		tx.Commit()
	}
}

func TestWeChatUpdateSendResult(t *testing.T) {
	tx, err := WeChatUpdateSendResult(context.Background(), &meta.DbWeChat{
		Id:           "8ea57cba-64dd-4477-9389-3a85d7269d38",
		Status:       3,
		ResultStatus: 1,
		TryNum:       2,
	})
	if err != nil {
		if tx != nil {
			if err := tx.Rollback(); err != nil {
				t.Error(err)
			}
		}
		t.Error(err)
	}
	if err := tx.Commit(); err != nil {
		t.Error(err)
	}
}
