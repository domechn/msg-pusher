/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : request_test.go
#   Created       : 2019/1/7 16:41
#   Last Modified : 2019/1/7 16:41
#   Describe      :
#
# ====================================================*/
package sms

import (
	"testing"
)

func TestNewRequest(t *testing.T) {
	req := NewRequest("1231231", "12421", "12421", "123123", "124w234")
	st, err := req.Encode("asddfs", "adscvcx")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf(st)
	}
}
