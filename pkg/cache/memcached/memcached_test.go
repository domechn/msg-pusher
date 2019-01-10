/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : memcached_test.go
#   Created       : 2019/1/8 11:16
#   Last Modified : 2019/1/8 11:16
#   Describe      :
#
# ====================================================*/
package memcached

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient("127.0.0.1:11211")
	if c == nil {
		t.Errorf("newclient error")
	}
}

func TestClient_Add(t *testing.T) {
	c := NewClient("127.0.0.1:11211")
	err := c.Add("test2", []byte("test"), 0)
	if err != nil {
		t.Errorf("add error:%v\n", err)
	}
}

func TestClient_Get(t *testing.T) {
	c := NewClient("127.0.0.1:11211")
	_, err := c.Get("test2")
	if err != nil {
		t.Errorf("get error:%v\n", err)
	}
}

func TestClient_Put(t *testing.T) {
	c := NewClient("127.0.0.1:11211")
	err := c.Put("test", []byte("test"), 0)
	if err != nil {
		t.Errorf("put error:%v\n", err)
	}
}

func TestClient_Del(t *testing.T) {
	c := NewClient("127.0.0.1:11211")
	if err := c.Del("test"); err != nil {
		t.Errorf("del error:%v\n", err)
	}
	if err := c.Del("test2"); err != nil {
		t.Errorf("del error:%v\n", err)
	}
}
