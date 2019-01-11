/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : canceler.go
#   Created       : 2019/1/11 17:45
#   Last Modified : 2019/1/11 17:45
#   Describe      :
#
# ====================================================*/
package cache

import (
	"context"
	"uuabc.com/sendmsg/api/storer"
)

const (
	cancelSet = "msg_cancel_set"
)

func CancelMsg(ctx context.Context, id string) error {
	return storer.Cache.Append(cancelSet, []byte(id))
}
