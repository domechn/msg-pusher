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
)

// LockID5s 独占锁
func LockID5s(ctx context.Context, k string) error {
	return add(ctx, "LockID5s", lock5s+k, []byte("lock"), 5)
}

// ReleaseLock 释放独占锁
func ReleaseLock(ctx context.Context, k string) error {
	return del(ctx, "ReleaseLock", lock5s+k)
}

// LockId 当需要操作一个公共数据时需要使用lock,锁最多持有10秒钟
func LockId(ctx context.Context, k string) error {
	return add(ctx, "LockId", lockId+k, []byte("lock"), 10)
}

// UnlockId 释放id锁
func UnlockId(ctx context.Context, k string) error {
	return del(ctx, "UnlockId", lockId+k)
}
