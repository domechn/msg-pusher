/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : email_test.go
#   Created       : 2019/1/14 10:25
#   Last Modified : 2019/1/14 10:25
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

func TestInsertEmails(t *testing.T) {
	tx, err := EmailInsert(
		context.Background(),
		&meta.DbEmail{
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

func TestEmailDetailByID(t *testing.T) {
	res, err := EmailDetailByID(context.Background(), "d1b1753f-d2d4-4c0c-b24b-bfdeeb8068bf")
	fmt.Println(res, err)
}

func TestEmailEdit(t *testing.T) {
	tx, err := EmailEdit(context.Background(), &meta.DbEmail{
		Id:       "8a9da75d-96ef-40ac-b0be-98a327d5e482",
		Content:  "hello",
		SendTime: "2018-08-08 08:08:08",
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

func TestEmailEditDestination(t *testing.T) {
	tx, err := EmailEdit(context.Background(), &meta.DbEmail{
		Id:          "8a9da75d-96ef-40ac-b0be-98a327d5e482",
		Content:     "hello",
		SendTime:    "2018-08-08 08:08:08",
		Destination: "test@email.com",
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
