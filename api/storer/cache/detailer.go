/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : detailer.go
#   Created       : 2019/1/11 17:48
#   Last Modified : 2019/1/11 17:48
#   Describe      :
#
# ====================================================*/
package cache

import (
	"context"
	"uuabc.com/sendmsg/api/storer"
)

func Detail(ctx context.Context, id string) ([]byte, error) {
	return storer.Cache.Get(id)
}

func StoreDetail(ctx context.Context, id string, value []byte, ttl int32) error {
	return storer.Cache.Put(id, value, ttl)
}
