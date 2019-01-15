/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : local.go
#   Created       : 2019/1/15 19:27
#   Last Modified : 2019/1/15 19:27
#   Describe      :
#
# ====================================================*/
package cache

import (
	"strings"
	"uuabc.com/sendmsg/api/storer"
)

func AddLocalTemplate(s string, v string) error {
	return storer.LocalCache.Put(s, []byte(v), 60)
}

func LocalTemplate(s string) ([]string, error) {
	b, err := storer.LocalCache.Get(s)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(b), ","), nil
}
