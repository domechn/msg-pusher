/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : editer.go
#   Created       : 2019/1/14 15:33
#   Last Modified : 2019/1/14 15:33
#   Describe      :
#
# ====================================================*/
package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/api/storer/cache"
	cache2 "uuabc.com/sendmsg/pkg/cache"
	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/pkg/pb/meta"
)

var EditerImpl editerImpl

type editerImpl struct{}

func (e editerImpl) Edit(ctx context.Context, id string, m Messager) error {
	err := e.edit(ctx, id, m)
	if err == nil {
		return nil
	}
	// 转换err类型
	if _, ok := err.(*errors.Error); !ok {
		if err == cache2.ErrCacheMiss {
			err = errors.ErrMsgNotFound
		} else {
			logrus.WithFields(logrus.Fields{
				"method": "Edit",
				"error":  err,
			}).Errorf("数据操作异常")
			err = errors.NewError(
				10000000,
				err.Error(),
			)
		}
	}
	return err
}

func (editerImpl) edit(ctx context.Context, id string, m Messager) error {
	b, err := cache.BaseDetail(ctx, id)
	if err != nil {
		return err
	}
	err = m.Unmarshal(b)
	if err != nil {
		return err
	}
	st := m.GetStatus()
	if st == meta.Status_Cancel {
		return errors.ErrMsgHasCancelled
	}
	if st == meta.Status_Final {
		return errors.ErrMsgCantEdit
	}

	return nil
}
