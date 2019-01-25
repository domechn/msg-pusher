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

	"uuabc.com/sendmsg/storer"
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

// MobileCache1Min 一分钟限流器+1，并返回+1后的结果，限制每个号码每分钟发送的频率
func MobileCache1Min(ctx context.Context, mobile string) (int64, error) {
	return limit(ctx, "MobileCache1Hour", mobile+"_1_min", 60)
}

// MobileCache1Hour 一小时限流器+1，并返回+1后的结果，限制一个号码的发送频率
func MobileCache1Hour(ctx context.Context, mobile string) (int64, error) {
	return limit(ctx, "MobileCache1Hour", mobile+"_1_hour", 60*60)
}

// MobileCache1Day 一天限流器+1，并返回+1后的结果，限制一个号码每天的发送频率
func MobileCache1Day(ctx context.Context, mobile string) (int64, error) {
	return limit(ctx, "MobileCache1Day", mobile+"_1_day", 60*60*24)
}
