/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : lock.go
#   Created       : 2019/1/18 17:35
#   Last Modified : 2019/1/18 17:35
#   Describe      :
#
# ====================================================*/
package cache

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"uuabc.com/sendmsg/storer"
)

// LockID5s 独占锁
func LockID5s(ctx context.Context, k string) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan("LockID5s", opentracing.ChildOf(parentCtx))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	return storer.Cache.Add(ctx, lock5s+k, []byte("lock"), 5)
}

// ReleaseLock 释放独占锁
func ReleaseLock(ctx context.Context, k string) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan("ReleaseLock", opentracing.ChildOf(parentCtx))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	return storer.Cache.Del(ctx, lock5s+k)
}

// LockId 当需要操作一个公共数据时需要使用lock,锁最多持有10秒钟
func LockId(k string) error {
	return storer.Cache.Add(context.Background(), "lock-id"+k, []byte("lock"), 10)
}

// UnlockId 释放id锁
func UnlockId(k string) error {
	return storer.Cache.Del(context.Background(), "lock-id"+k)
}
