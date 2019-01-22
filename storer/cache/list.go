/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : list.go
#   Created       : 2019/1/22 19:34
#   Last Modified : 2019/1/22 19:34
#   Describe      :
#
# ====================================================*/
package cache

import (
	"context"

	"uuabc.com/sendmsg/storer"
)

// RPushWeChat lpush到redis，用来代替存入数据库，提高并发能力
func RPushWeChat(b []byte) error {
	return storer.Cache.RPush(context.Background(), "", b)
}

func RPushEmail(b []byte) error {
	return storer.Cache.RPush(context.Background(), "", b)
}

func RPushSms(b []byte) error {
	return storer.Cache.RPush(context.Background(), "", b)
}

// LLenFromList 查看list中的数据量
func LLenFromList() (int64, error) {
	return storer.Cache.LLen(context.Background(), "")
}

// LPopFromList 从list中取出数据
func LPopFromList() ([]byte, error) {
	return storer.Cache.LPop(context.Background(), "")
}
