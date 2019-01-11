/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : detailer_test.go
#   Created       : 2019/1/10 20:12
#   Last Modified : 2019/1/10 20:12
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"fmt"
	"testing"
	"uuabc.com/sendmsg/api/storer"
	"uuabc.com/sendmsg/pkg/db"
)

func init() {
	storer.DB, _ = db.New(db.Config{
		URL: "root:root@tcp(localhost:3306)/uuabc?charset=utf8&parseTime=True",
	})
}

func TestSmsDetailByID(t *testing.T) {
	res, err := SmsDetailByID(context.Background(), "42b86258-e242-4b67-b172-40dafa539972")
	fmt.Println(res, err)
}

func TestWeChatDetailByID(t *testing.T) {
	res, err := WeChatDetailByID(context.Background(), "00a87366-c607-43d8-9673-6f2cd143273c")
	fmt.Println(res, err)
}

func TestEmailDetailByID(t *testing.T) {
	res, err := EmailDetailByID(context.Background(), "d1b1753f-d2d4-4c0c-b24b-bfdeeb8068bf")
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
