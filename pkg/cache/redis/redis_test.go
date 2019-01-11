/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : redis_test.go
#   Created       : 2019/1/9 19:12
#   Last Modified : 2019/1/9 19:12
#   Describe      :
#
# ====================================================*/
package redis

import (
	"bytes"
	"fmt"
	"testing"
)

var cli *Client
var err error

func init() {
	cli, err = NewClient([]string{"127.0.0.1:6379"}, "")
}

var cases = []struct {
	name  string
	key   string
	value []byte
	res   []byte
	want  error
}{
	{
		name:  "case1",
		key:   "test-abc",
		value: []byte("hello-test"),
		res:   []byte("hello-test"),
		want:  nil,
	},
}

func TestPut(t *testing.T) {
	fmt.Printf("%v", err)
	for _, v := range cases {
		if err := cli.Put(v.key, v.value, 0); err != v.want {
			t.Errorf("%s test func Put() failed,want: %v actual: %v", v.name, v.want, err)
		}
	}
}

func TestAdd(t *testing.T) {
	for _, v := range cases {
		if err := cli.Add(v.key, v.value, 0); err == v.want {
			t.Errorf("%s test func Add() failed,want: %v actual: %v", v.name, v.want, err)
		}
	}
}

func TestGet(t *testing.T) {
	for _, v := range cases {
		if res, err := cli.Get(v.key); err != v.want || bytes.Compare(res, v.res) != 0 {
			t.Errorf("%s test func Get() failed,want: %v res:%s, actual: %v res:%s", v.name, v.want, err, string(res), string(v.res))
		}
	}
}

func TestDel(t *testing.T) {
	for _, v := range cases {
		if err := cli.Del(v.key); err != v.want {
			t.Errorf("%s test func Del() failed,want: %v actual: %v", v.name, v.want, err)
		}
		if r, er := cli.Get(v.key); er == nil {
			t.Errorf("%s test func Del() failed,want: %v actual: %v", v.name, string(v.res), string(r))
		}
	}
}

func TestNewClient(t *testing.T) {
	cli, err := NewClient([]string{"127.0.0.1:6379"}, "")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cli.Get("hello"))
	cli.Close()
}
