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
func RPushWeChat(ctx context.Context, b []byte) error {
	return rPush(ctx, "RPushWeChat", weChatDB, b, storer.Cache)
}

// RPushEmail lpush到email队列
func RPushEmail(ctx context.Context, b []byte) error {
	return rPush(ctx, "RPushEmail", emailDB, b, storer.Cache)
}

// RPushSms lpush到sms队列
func RPushSms(ctx context.Context, b []byte) error {
	return rPush(ctx, "RPushSms", smsDB, b, storer.Cache)
}

// LLenWeChat 查看wechat入库队列的长度
func LLenWeChat() (int64, error) {
	return storer.Cache.LLen(context.Background(), weChatDB)
}

// LLenEmail 查看email入库队列的长度
func LLenEmail() (int64, error) {
	return storer.Cache.LLen(context.Background(), emailDB)
}

// LLenSms 查看sms入库队列的长度
func LLenSms() (int64, error) {
	return storer.Cache.LLen(context.Background(), smsDB)
}

// LPopWeChat 从wechat队列中取一条数据
func LPopWeChat() ([]byte, error) {
	return lPop(context.Background(), weChatDB, storer.Cache)
}

// LPopEmail 从email队列中取一条数据
func LPopEmail() ([]byte, error) {
	return lPop(context.Background(), emailDB, storer.Cache)
}

// LPopSms 从sms队列中取一条数据
func LPopSms() ([]byte, error) {
	return lPop(context.Background(), smsDB, storer.Cache)
}
