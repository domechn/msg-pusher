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
)

var (
	ErrNoRowsEffected = errors.New("db: no rows affected")
)
