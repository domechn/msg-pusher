/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : cache.go
#   Created       : 2019/1/8 10:43
#   Last Modified : 2019/1/8 10:43
#   Describe      :
#
# ====================================================*/
package cache

import (
	"context"
	"fmt"
	"io"
)

var (
	// ErrCacheMiss 未找到对应的缓存
	ErrCacheMiss = fmt.Errorf("cache miss")
	// ErrKeyExsit 需要缓存的key已存在
	ErrKeyExsit = fmt.Errorf("key exsits")
)

// Cache 缓存接口，实现了增删改查功能
type Cache interface {
	Get(ctx context.Context, s string) ([]byte, error)
	// 如果ttl等于0 则kv永久有效,ttl单位 秒
	Put(ctx context.Context, k string, v []byte, ttl int64) error
	Del(ctx context.Context, k string) error
	// 如果key存在就报错，只有不存在时才能设置成功,ttl单位 秒
	Add(ctx context.Context, k string, v []byte, ttl int64) error

	// 将数据添加到set，k为列表名称，v为值
	Append(ctx context.Context, k string, v []byte) error
	// v是否在k列表中
	IsMember(ctx context.Context, k string, v []byte) (bool, error)

	io.Closer
}
