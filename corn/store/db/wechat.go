/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : wechat.go
#   Created       : 2019/1/24 11:27
#   Last Modified : 2019/1/24 11:27
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"database/sql"

	"github.com/domgoer/msg-pusher/config"
	"github.com/domgoer/msg-pusher/corn/store"
	"github.com/domgoer/msg-pusher/pkg/pb/meta"
	"github.com/domgoer/msg-pusher/pkg/utils"
	"github.com/domgoer/msg-pusher/storer/cache"
	"github.com/domgoer/msg-pusher/storer/db"
	"github.com/sirupsen/logrus"
)

type WeChat struct {
	len int64
}

func registerWeChat() {
	store.MustRegisterCorn("wechat-corn", NewWeChatCorn(config.CornConf().MaxLen))
}

func (e *WeChat) Read() ([][]byte, error) {
	return read(cache.LLenWeChat, cache.LPopWeChat, e.len)
}

func (e *WeChat) Write(param [][]byte) (err error) {
	if len(param) == 0 {
		return
	}
	var li []*meta.DbWeChat
	for _, b := range param {
		dbWeChat := &meta.DbWeChat{}
		if err := dbWeChat.Unmarshal(b); err != nil {
			logrus.WithFields(logrus.Fields{
				"type":   "wechat",
				"method": "read",
				"data":   string(b),
			}).Error("redis中存在错误数据")
			continue
		}
		dbWeChat.SetSendTime(utils.MustISO8601StrToUTCStr(dbWeChat.GetSendTime()))
		li = append(li, dbWeChat)
	}
	if err = db.WeChatUpdateAndInsertBatch(context.Background(), li); err != nil {
		// 如果是数据库无法连接，就将数据回滚到redis
		if err == sql.ErrConnDone {
			logrus.Errorf("批量插入数据库失败，数据库连接已关闭，正在将数据回滚到redis")
			t := cache.NewTransaction()
			defer t.Close()
			for _, p := range param {
				t.RPushWeChat(context.Background(), p)
			}
			t.Commit(context.Background())
		}
	}
	return
}

func (e *WeChat) Name() string {
	return "weChat"
}

// NewWeChatCorn 初始化一个定时写入的任务，n为每次读取和写入的最大数据量
func NewWeChatCorn(n int64) *WeChat {
	return &WeChat{
		len: n,
	}
}
