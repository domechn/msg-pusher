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
	"sync"

	"github.com/domgoer/msg-pusher/config"
	"github.com/domgoer/msg-pusher/corn/store"
	"github.com/domgoer/msg-pusher/pkg/pb/meta"
	"github.com/domgoer/msg-pusher/pkg/utils"
	"github.com/domgoer/msg-pusher/storer/cache"
	"github.com/domgoer/msg-pusher/storer/db"
	"github.com/sirupsen/logrus"
)

type Sms struct {
	len int64
	sync.RWMutex
}

func registerSms() {
	store.MustRegisterCorn("sms-corn", NewSmsCorn(config.CornConf().MaxLen))
}

func (e *Sms) Read() ([][]byte, error) {
	return read(cache.LLenSms, cache.LPopSms, e.len)
}

func (e *Sms) Write(param [][]byte) (err error) {
	if len(param) == 0 {
		return
	}
	var li []*meta.DbSms
	for _, b := range param {
		dbSms := &meta.DbSms{}
		if err := dbSms.Unmarshal(b); err != nil {
			logrus.WithFields(logrus.Fields{
				"type":   "sms",
				"method": "read",
				"data":   string(b),
			}).Error("redis中存在错误数据")
			continue
		}
		dbSms.SetSendTime(utils.MustISO8601StrToUTCStr(dbSms.GetSendTime()))
		li = append(li, dbSms)
	}
	if err = db.SmsUpdateAndInsertBatch(context.Background(), li); err != nil {
		// 如果是数据库无法连接，就将数据回滚到redis
		if err == sql.ErrConnDone {
			logrus.Errorf("批量插入数据库失败，数据库连接已关闭，正在将数据回滚到redis")
			t := cache.NewTransaction()
			defer t.Close()
			for _, p := range param {
				t.RPushSms(context.Background(), p)
			}
			t.Commit(context.Background())
		}
	}
	return err
}

func (e *Sms) Name() string {
	return "sms"
}

// NewSmsCorn 初始化一个定时写入的任务，n为每次读取和写入的最大数据量
func NewSmsCorn(n int64) *Sms {
	return &Sms{
		len: n,
	}
}
