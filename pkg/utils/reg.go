/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : reg.go
#   Created       : 2019/1/10 16:37
#   Last Modified : 2019/1/10 16:37
#   Describe      :
#
# ====================================================*/
package utils

import (
	"regexp"
)

func ValidatePhone(s string) bool {
	b, _ := regexp.Match(`^1\d{10}$`, []byte(s))
	return b
}

func ValidateTemplate(s string) bool {
	b, _ := regexp.Match(`^SMS_[0-9]{9}`, []byte(s))
	return b
}

func ValidateEmailAddr(s string) bool {
	b, _ := regexp.Match(`^[A-Za-z\d]+([-_.][A-Za-z\d]+)*@([A-Za-z\d]+[-.])+[A-Za-z\d]{2,4}$`, []byte(s))
	return b
}
