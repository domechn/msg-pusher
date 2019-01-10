/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : send_test.go
#   Created       : 2019/1/7 15:09
#   Last Modified : 2019/1/7 15:09
#   Describe      :
#
# ====================================================*/
package mail

import (
	"testing"
)

func TestSend(t *testing.T) {
	s := NewClient(&Config{
		ServerAddr: "smtp.exmail.qq.com:587",
		Host:       "smtp.exmail.qq.com",
		Username:   "passwordreset@uuabc.com",
		Password:   "uuAB12",
	})

	if err := s.Send(NewMessage("814172254@qq.com", []byte("hello")), nil); err != nil {
		t.Error(err)
	}
}
