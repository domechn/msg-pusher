/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : reg_test.go
#   Created       : 2019/1/15 16:12
#   Last Modified : 2019/1/15 16:12
#   Describe      :
#
# ====================================================*/
package utils

import (
	"testing"
)

func TestValidateUUIDV4(t *testing.T) {

}

func TestValidateEmailAddr(t *testing.T) {

}

func TestValidatePhone(t *testing.T) {

}

func TestValidateTemplate(t *testing.T) {

}

func TestStrFromCurlyBraces(t *testing.T) {
	var cases = []struct {
		name string
		v    string
		want []string
	}{
		{
			name: "case1",
			v:    "${aa}asddsfg${cc}asdafd${abc}",
			want: []string{"${aa}", "${cc}", "${abc}"},
		}, {
			name: "case2",
			v:    "abc",
			want: []string{},
		}, {
			name: "case3",
			v:    "${aa}asddsfg{cc}asdafd${abc}",
			want: []string{"${aa}", "${abc}"},
		},
	}
	for _, v := range cases {
		if res := StrFromCurlyBraces(v.v); !equal(res, v.want) {
			t.Errorf("case:%s,do StrFromCurlyBraces() error,want: %v,actual: %v", v.name, v.want, res)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		var flag bool
		for _, v2 := range b {
			if v2 == v {
				flag = true
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
