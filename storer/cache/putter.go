/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : putter.go
#   Created       : 2019/1/14 10:58
#   Last Modified : 2019/1/14 10:58
#   Describe      :
#
# ====================================================*/
package cache

import (
	"context"
	"math/rand"
	"uuabc.com/sendmsg/storer"
)

// PutBaseCache 底层缓存，跟数据库数据同步，不过期,并发安全
func PutBaseCache(ctx context.Context, k string, v []byte) error {
	if err := storer.Cache.Add(ctx, "lock-key-"+k, []byte("lock"), 3); err != nil {
		return err
	}
	defer storer.Cache.Del(ctx, "lock-key-"+k)
	return storer.Cache.Put(ctx, base+k, v, 0)
}

// PutLastestCache 最新缓存，保证数据时效性，默认5+n(n<5)秒缓存
func PutLastestCache(ctx context.Context, k string, v []byte) error {
	return storer.Cache.Put(ctx, lastest+k, v, int64(5+rand.Intn(5)))
}

// LockID5s 独占锁
func LockID5s(ctx context.Context, k string) error {
	return storer.Cache.Add(ctx, "lock-5s-"+k, []byte("lock"), 5)
}

// ReleaseLock 释放独占锁
func ReleaseLock(ctx context.Context, k string) error {
	return storer.Cache.Del(ctx, "lock-5s-"+k)
}

func PutBaseTemplate(ctx context.Context, k string, v []byte) error {
	return storer.Cache.Put(ctx, template+k, v, 0)
}
