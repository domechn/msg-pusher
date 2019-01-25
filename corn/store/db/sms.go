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
	"sync"

	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/config"
	"uuabc.com/sendmsg/corn/store"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/pkg/utils"
	"uuabc.com/sendmsg/storer/cache"
	"uuabc.com/sendmsg/storer/db"
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

func (e *Sms) Write(param [][]byte) error {
	if len(param) == 0 {
		return nil
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
	return db.SmsUpdateAndInsertBatch(context.Background(), li)
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
