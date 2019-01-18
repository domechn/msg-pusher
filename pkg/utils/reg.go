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
	valid "github.com/asaskevich/govalidator"
	"regexp"
)

// ValidatePhone 验证手机号，符合就返回true
func ValidatePhone(s string) bool {
	b, _ := regexp.Match(`^1\d{10}$`, []byte(s))
	return b
}

// ValidateEmailAddr 验证邮箱格式，符合返回true
func ValidateEmailAddr(s string) bool {
	return valid.IsEmail(s)
}

// StrFromCurlyBraces 从文本中获取${...}的值
func StrFromCurlyBraces(s string) []string {
	rg, _ := regexp.Compile(`\${(.*?)}`)
	return rg.FindAllString(s, -1)
}
