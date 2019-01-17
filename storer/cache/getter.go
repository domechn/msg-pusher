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

	"uuabc.com/sendmsg/storer"
)

func BaseDetail(ctx context.Context, k string) ([]byte, error) {
	return storer.Cache.Get(ctx, base+k)
}

func BaseTemplate(ctx context.Context, k string) (string, error) {
	b, err := storer.Cache.Get(ctx, template+k)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func LastestDetail(ctx context.Context, k string) ([]byte, error) {
	return storer.Cache.Get(ctx, lastest+k)
}

func Detail(ctx context.Context, id string) ([]byte, error) {
	return storer.Cache.Get(ctx, id)
}

func StoreDetail(ctx context.Context, id string, value []byte, ttl int64) error {
	return storer.Cache.Put(ctx, id, value, ttl)
}
