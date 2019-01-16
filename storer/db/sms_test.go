/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : sms_test.go
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

func TestInsertSmss(t *testing.T) {
	tx, err := SmsInsert(
		context.Background(),
		&meta.DbSms{
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

func TestSmsDetailByID(t *testing.T) {
	res, err := SmsDetailByID(context.Background(), "42b86258-e242-4b67-b172-40dafa539972")
	fmt.Println(res, err)
}

func TestSmsDetailByPhoneAndPage(t *testing.T) {
	res, err := SmsDetailByPhoneAndPage(context.Background(), "18516051096", 1)
	if err != nil {
		t.Error(err)
	} else {
		for _, v := range res {
			fmt.Println(v)
		}
	}
}

func TestSmsCancelMsgByID(t *testing.T) {
	tx, err := SmsCancelMsgByID(context.Background(), "835dc583-466f-4b3b-94fc-2feef9ec8098")
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		t.Error(err)
	} else {
		tx.Commit()
	}
}

func TestSmsEdit(t *testing.T) {
	tx, err := SmsEdit(context.Background(), &meta.DbSms{
		Id:       "835dc583-466f-4b3b-94fc-2feef9ec8098",
		Content:  "hello",
		SendTime: "2018-08-08 08:08:08",
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

func TestSmsEditMobile(t *testing.T) {
	tx, err := SmsEdit(context.Background(), &meta.DbSms{
		Id:       "835dc583-466f-4b3b-94fc-2feef9ec8098",
		Content:  "hello2",
		SendTime: "2018-08-18 08:08:08",
		Mobile:   "13151545542",
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
