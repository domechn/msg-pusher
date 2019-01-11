/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : uuid.go
#   Created       : 2019/1/11 15:02
#   Last Modified : 2019/1/11 15:02
#   Describe      :
#
# ====================================================*/
package utils

import (
	"github.com/satori/go.uuid"
)

func ValidateUUIDV4(id string) error {
	_, err := uuid.FromString(id)
	return err
}
