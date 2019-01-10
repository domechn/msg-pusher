/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : mysql_test.go
#   Created       : 2019/1/9 16:22
#   Last Modified : 2019/1/9 16:22
#   Describe      :
#
# ====================================================*/
package db

import (
	"testing"
)

func TestNew(t *testing.T) {
	cases := []struct {
		name string
		cfg  Config
		want error
	}{
		{
			name: "case1",
			cfg: Config{
				URL: "root:root@tcp(localhost:3306)/uuabc?charset=utf8&parseTime=True",
			},
			want: nil,
		},
	}

	for _, v := range cases {
		if m, err := New(v.cfg); err != v.want {
			t.Errorf("func: New() failed,caseName: %s,want: %v,actual: %v", v.name, v.want, err)
		} else {
			m.Close()
		}
	}
}
