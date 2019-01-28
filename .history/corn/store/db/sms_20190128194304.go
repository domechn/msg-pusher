/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : sms.go
#   Created       : 2019/1/24 11:27
#   Last Modified : 2019/1/24 11:27
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"database/sql"

	"github.com/hiruok/msg-pusher/config"
	"github.com/hiruok/msg-pusher/corn/store"
	"github.com/hiruok/msg-pusher/pkg/pb/meta"
	"github.com/hiruok/msg-pusher/pkg/utils"
	"github.com/hiruok/msg-pusher/storer/cache"
	"github.com/hiruok/msg-pusher/storer/db"
	"github.com/sirupsen/logrus"
)

// Msg 定时读取缓存中的数据，持久化到数据库
type Msg struct {
	len int64
}

func registerMsg() {
	store.MustRegisterCorn("msg-corn", NewMsgCorn(config.CornConf().MaxLen))
}

// Read 读取缓存中的信息
func (e *Msg) Read() ([][]byte, error) {
	return read(cache.LLenMsg, cache.LPopMsg, e.len)
}

// Write 将缓存中的信息批量存入数据库
func (e *Msg) Write(param [][]byte) (err error) {
	if len(param) == 0 {
		return
	}
	var li []*meta.DbMsg
	for _, b := range param {
		dbSms := &meta.DbMsg{}
		if err := dbSms.Unmarshal(b); err != nil {
			logrus.WithFields(logrus.Fields{
				"type":   "msg",
				"method": "read",
				"data":   string(b),
			}).Error("redis中存在错误数据")
			continue
		}
		dbSms.SetSendTime(utils.MustISO8601StrToUTCStr(dbSms.GetSendTime()))
		li = append(li, dbSms)
	}
	if err = db.UpdateAndInsertMsgBatch(context.Background(), li); err != nil {
		// 如果是数据库无法连接，就将数据回滚到redis
		if err == sql.ErrConnDone {
			logrus.Errorf("批量插入数据库失败，数据库连接已关闭，正在将数据回滚到redis")
			t := cache.NewTransaction()
			defer t.Close()
			for _, p := range param {
				t.RPushMsg(context.Background(), p)
			}
			t.Commit(context.Background())
		}
	}
	return err
}

func (e *Msg) Name() string {
	return "msg"
}

// NewMsgCorn 初始化一个定时写入的任务，n为每次读取和写入的最大数据量
func NewMsgCorn(n int64) *Msg {
	return &Msg{
		len: n,
	}
}
