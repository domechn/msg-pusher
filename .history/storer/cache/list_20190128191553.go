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

	"github.com/hiruok/msg-pusher/storer"
)

func RPushMsg(ctx context.Context, b []byte) error {
	return rPush(ctx, "RPushMsg", msgDB, b, storer.Cache)
}

// LLenMsg 查看入库队列的长度
func LLenMsg() (int64, error) {
	return storer.Cache.LLen(context.Background(), msgDB)
}

// LPopMsg 从入库队列中取一条数据
func LPopMsg() ([]byte, error) {
	return lPop(context.Background(), msgDB, storer.Cache)
}
