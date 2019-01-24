/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : email.go
#   Created       : 2019/1/23 16:58
#   Last Modified : 2019/1/23 16:58
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/config"
	"uuabc.com/sendmsg/corn/store"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/pkg/utils"
	"uuabc.com/sendmsg/storer/cache"
	"uuabc.com/sendmsg/storer/db"
)

type Email struct {
	len int64
}

func registerEmail() {
	store.MustRegisterCorn("email-corn", NewEmailCorn(config.CornConf().MaxLen))
}

func (e *Email) Read() ([][]byte, error) {
	return read(cache.LLenEmail, cache.LPopEmail, e.len)
}

func (e *Email) Write(param [][]byte) error {
	if len(param) == 0 {
		return nil
	}
	var li []*meta.DbEmail
	for _, b := range param {
		dbEmail := &meta.DbEmail{}
		if err := dbEmail.Unmarshal(b); err != nil {
			logrus.WithFields(logrus.Fields{
				"type":   "email",
				"method": "read",
				"data":   string(b),
			}).Error("redis中存在错误数据")
			continue
		}
		dbEmail.SetSendTime(utils.MustISO8601StrToUTCStr(dbEmail.GetSendTime()))
		li = append(li, dbEmail)
	}
	var err error
	if err = db.EmailUpdateAndInsertBatch(context.Background(), li); err != nil {
		// 如果是数据库无法连接，就将数据回滚到redis
		if err == sql.ErrConnDone {
			logrus.Errorf("批量插入数据库失败，数据库连接已关闭，正在将数据回滚到redis")
			t := cache.NewTransaction()
			defer t.Close()
			for _, p := range param {
				t.RPushEmail(context.Background(), p)
			}
			t.Commit()
		}
	}
	return err
}

func (e *Email) Name() string {
	return "email"
}

// NewEmailCorn 初始化一个定时写入的任务，n为每次读取和写入的最大数据量
func NewEmailCorn(n int64) *Email {
	return &Email{
		len: n,
	}
}
