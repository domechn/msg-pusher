/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : send_test.go
#   Created       : 2019/1/7 20:09
#   Last Modified : 2019/1/7 20:09
#   Describe      :
#
# ====================================================*/
package wechat

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_AccessTokenData(t *testing.T) {
	c := NewClient(&Config{
		APPId:      "wxb08951fc140c4b9d",
		APPSecret:  "46f69c013b01ad309405bb9bf1ec6bf1",
		CacheAddrs: []string{"127.0.0.1:6379"},
	})
	if err := c.Send(&Request{}, nil); err != nil {
		t.Error(err)
	}
}

func TestClient_Send(t *testing.T) {
	c := NewClient(&Config{
		APPId:     "wxb08951fc140c4b9d",
		APPSecret: "46f69c013b01ad309405bb9bf1ec6bf1",
	})
	res := &Response{}
	msg := map[string]interface{}{
		"first": map[string]string{
			"hello": "hi",
		},
	}
	b, _ := json.Marshal(msg)
	err := c.Send(NewRequest("hello", "1", "", b), func(a interface{}) {
		if v, ok := a.(*Response); ok {
			res = v
		}
	})
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}
