/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : var.go
#   Created       : 2019/1/14 14:26
#   Last Modified : 2019/1/14 14:26
#   Describe      :
#
# ====================================================*/
package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

var (
	ErrNoRowsEffected  = errors.New("db: no rows affected")
	ErrUniqueKeyExsits = errors.New("db: Duplicate entry for key")
)

// changeSendTime 修改发送时间格式"2006-01-02T15:04:05Z"->"2006-01-02 15:05:05"
func changeSendTime(s string) string {
	s = strings.Replace(s, "T", " ", -1)
	return strings.Replace(s, "Z", "", -1)
}

// 带空判断的回滚
func RollBack(tx *sqlx.Tx) error {
	if tx != nil {
		return tx.Rollback()
	}
	return nil
}

// Commit 带空判断的提交
func Commit(tx *sqlx.Tx) error {
	if tx == nil {
		return fmt.Errorf("tx is nil")
	}
	return tx.Commit()
}
