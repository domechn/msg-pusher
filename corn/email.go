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
package corn

import (
	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/storer/cache"
)

type Email struct {
}

func (e *Email) Read(n int64) (inserts []Marshaler, updates []Marshaler) {
	len, err := cache.LLenEmail()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"method": "email.Read",
			"error":  err.Error(),
		}).Error("出现位置错误")
	}
	if len > n {
		len = n
	}

	for i := 0; i < int(len); i++ {
		b, err := cache.LPopEmail()
		if err != nil {
			continue
		}
		dbEmail := &meta.DbEmail{}
		if err = dbEmail.Unmarshal(b); err != nil {
			logrus.WithFields(logrus.Fields{
				"type":   "email",
				"method": "read",
				"data":   string(b),
			}).Error("redis中存在错误数据")
		}
		if dbEmail.Option == int32(meta.Insert) {
			inserts = append(inserts, dbEmail)
		} else if dbEmail.Option == int32(meta.Update) {
			updates = append(updates, dbEmail)
		}
	}

	return
}

func (e *Email) Write(inserts, updates []Marshaler) error {

	panic("implement me")
}

func (e *Email) Start() {
	panic("implement me")
}

func NewEmailCorn() *Email {
	return &Email{}
}
