/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : getter.go
#   Created       : 2019/1/11 17:48
#   Last Modified : 2019/1/11 17:48
#   Describe      :
#
# ====================================================*/
package cache

import (
	"context"
	"strings"
	"uuabc.com/sendmsg/api/storer"
)

func BaseDetail(ctx context.Context, k string) ([]byte, error) {
	return storer.Cache.Get(base + k)
}

func BaseTemplate(ctx context.Context, k string) ([]string, error) {
	b, err := storer.Cache.Get(template + k)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(b), ","), nil
}

func LastestDetail(ctx context.Context, k string) ([]byte, error) {
	return storer.Cache.Get(lastest + k)
}

func Detail(ctx context.Context, id string) ([]byte, error) {
	return storer.Cache.Get(id)
}

func StoreDetail(ctx context.Context, id string, value []byte, ttl int64) error {
	return storer.Cache.Put(id, value, ttl)
}
