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
	"math/rand"
	"uuabc.com/sendmsg/api/storer"
)

// PutBaseCache 底层缓存，跟数据库数据同步，不过期,并发安全
func PutBaseCache(k string, v []byte) error {
	if err := storer.Cache.Add("lock-key-"+k, []byte("lock"), 3); err != nil {
		return err
	}
	defer storer.Cache.Del("lock-key-" + k)
	return storer.Cache.Put(base+k, v, 0)
}

// PutLastestCache 最新缓存，保证数据时效性，默认5+n(n<5)秒缓存
func PutLastestCache(k string, v []byte) error {
	return storer.Cache.Put(lastest+k, v, int64(5+rand.Intn(5)))
}

func LockID5s(k string) error {
	return storer.Cache.Add("lock-5s-"+k, []byte("lock"), 5)
}
