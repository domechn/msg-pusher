/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : send_test.go
#   Created       : 2019/1/7 17:11
#   Last Modified : 2019/1/7 17:11
#   Describe      :
#
# ====================================================*/
package sms

import (
	"fmt"
	"testing"
)

func TestClient_Send(t *testing.T) {
	client, err := NewClient(&Config{
		AccessKeyId:  "LTAIz878ukp5olep",
		AccessSecret: "01T1gdW8wZXhFzrR1B142EHdLW3AjF",
		GatewayURL:   "http://dysmsapi.aliyuncs.com/",
	})
	if err != nil {
		t.Error(err)
	}
	res := &Response{}
	do := func(a interface{}) {
		if v, ok := a.(*Response); ok {
			res = v
		}
	}
	err = client.Send(NewRequest("13151576692", "阿里云短信测试专用", "SMS_130990029", `{"code":"4401"}`, "12345"),
		do)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(res.RawResponse))
	}
}
