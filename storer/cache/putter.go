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

	"github.com/domgoer/msg-pusher/storer"
)

// PutBaseCache 底层缓存，跟数据库数据同步，不过期
func PutBaseCache(ctx context.Context, k string, v []byte) error {
	return put(ctx, "PutBaseCache", base+k, v, 0, storer.Cache)
}

// PutBaseTemplate 将添加的模板存入缓存中
func PutBaseTemplate(ctx context.Context, k string, v []byte) error {
	return put(ctx, "PutBaseTemplate", template+k, v, 0, storer.Cache)
}

// PutSendResult 在bitmap中修改发送结果，一般只有发送成功的情况才需要设置
func PutSendSuccess(ctx context.Context, k string) error {
	return put(ctx, "PutSendSuccess", k+"_send", success, 0, storer.Cache)
}
