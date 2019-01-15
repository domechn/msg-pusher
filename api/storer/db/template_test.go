/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : template_test.go
#   Created       : 2019/1/15 15:42
#   Last Modified : 2019/1/15 15:42
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"fmt"
	"testing"
	"uuabc.com/sendmsg/pkg/pb/tpl"
)

func TestTemplateInsert(t *testing.T) {
	tx, err := TemplateInsert(context.Background(), &tpl.DBTemplate{
		Id:       "123",
		Type:     1,
		SimpleID: "123213",
		Content:  "123123",
	})
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		t.Error(err)
	}
	tx.Commit()
}

func TestTemplateList(t *testing.T) {
	res, err := TemplateList(context.Background())
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}
