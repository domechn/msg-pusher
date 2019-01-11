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
	"fmt"
	"io"
)

var (
	ErrCacheMiss = fmt.Errorf("cache miss")
	ErrKeyExsit  = fmt.Errorf("key exsits")
)

// Cache 缓存接口，实现了增删改查功能
type Cache interface {
	Get(s string) ([]byte, error)
	// 如果ttl等于0 则kv永久有效,ttl单位 秒
	Put(k string, v []byte, ttl int32) error
	Del(k string) error
	// 如果key存在就报错，只有不存在时才能设置成功,ttl单位 秒
	Add(k string, v []byte, ttl int32) error

	io.Closer
}
