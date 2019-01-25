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

	"github.com/domgoer/msgpusher/config"
	"github.com/domgoer/msgpusher/corn/store"
	"github.com/domgoer/msgpusher/pkg/pb/meta"
	"github.com/domgoer/msgpusher/pkg/utils"
	"github.com/domgoer/msgpusher/storer/cache"
	"github.com/domgoer/msgpusher/storer/db"
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

func (e *WeChat) Write(param [][]byte) error {
	if len(param) == 0 {
		return nil
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
	return db.WeChatUpdateAndInsertBatch(context.Background(), li)
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
