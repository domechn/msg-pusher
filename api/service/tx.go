/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : tx.go
#   Created       : 2019/1/14 10:36
#   Last Modified : 2019/1/14 10:36
#   Describe      :
#
# ====================================================*/
package service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func rollback(tx *sqlx.Tx) error {
	if tx == nil {
		return nil
	}
	return tx.Rollback()
}

func commit(tx *sqlx.Tx) error {
	if tx == nil {
		return fmt.Errorf("commit: tx is nil")
	}
	return tx.Commit()
}
