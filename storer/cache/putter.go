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
	"bytes"
	"context"
	"math/rand"

	"uuabc.com/sendmsg/storer"
)

// PutBaseCache 底层缓存，跟数据库数据同步，不过期
func PutBaseCache(ctx context.Context, k string, v []byte) error {
	return putBaseCache(ctx, k, v, false)
}

// putBaseCache 如果b为true则一个key修改之前要获取锁，获取成功再修改，
// 如果未false无需获取锁，直接修改
func putBaseCache(ctx context.Context, k string, v []byte, b bool) error {
	if b {
		if err := LockID5s(ctx, k); err != nil {
			return err
		}
		defer ReleaseLock(ctx, k)
	}
	return storer.Cache.Put(ctx, base+k, v, 0)
}

// PutLastestCache 最新缓存，保证数据时效性，默认5+n(n<5)秒缓存
func PutLastestCache(ctx context.Context, k string, v []byte) error {
	return storer.Cache.Put(ctx, lastest+k, v, int64(5+rand.Intn(5)))
}

// PutBaseTemplate 将添加的模板存入缓存中
func PutBaseTemplate(ctx context.Context, k string, v []byte) error {
	return storer.Cache.Put(ctx, template+k, v, 0)
}

// PutSendResult 在bitmap中修改发送结果，一般只有发送成功的情况才需要设置
func PutSendSuccess(ctx context.Context, k string) error {
	return storer.Cache.Put(ctx, k+"_sent", success, 0)
}

// SendResult 获取发送结果
func SendResult(ctx context.Context, k string) (bool, error) {
	res, err := storer.Cache.Get(ctx, k+"_sent")
	if err != nil {
		return false, err
	}
	if bytes.Compare(res, success) == 0 {
		return true, nil
	}
	return false, nil
}
