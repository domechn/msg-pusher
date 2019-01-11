/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : canceler.go
#   Created       : 2019/1/10 19:58
#   Last Modified : 2019/1/10 19:58
#   Describe      :
#
# ====================================================*/
package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/api/storer/cache"
	"uuabc.com/sendmsg/api/storer/db"
	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/pkg/retry/backoff"
)

var Canceler cancelerImpl

type getFunc func(context.Context, string) (Messager, error)
type updateFunc func(i context.Context, s string) error

type cancelerImpl struct {
}

// Cancel 根据id取消发送消息
func (c cancelerImpl) Cancel(ctx context.Context, typeN, id string) error {
	err := c.cancel(ctx, typeN, id)
	if err == nil {
		return nil
	}
	if _, ok := err.(*errors.Error); !ok {
		err = errors.NewError(
			10000000,
			err.Error(),
		)
	}
	return err
}

func (cancelerImpl) cancel(ctx context.Context, typeN, id string) (err error) {
	var g getFunc
	var u updateFunc
	switch typeN {
	case wechat:
		g = func(i context.Context, s string) (Messager, error) {
			return db.WeChatDetailByID(i, s)
		}
		u = func(i context.Context, s string) error {
			return db.CancelWeChatMsgByID(i, s)
		}
	case sms:
		g = func(i context.Context, s string) (Messager, error) {
			return db.SmsDetailByID(i, s)
		}
		u = func(i context.Context, s string) error {
			return db.CancelSmsMsgByID(i, s)
		}
	case email:
		g = func(i context.Context, s string) (Messager, error) {
			return db.EmailDetailByID(i, s)
		}
		u = func(i context.Context, s string) error {
			return db.CancelEmailMsgByID(i, s)
		}
	default:
		return errors.ErrMsgTypeNotFound
	}
	err = cancel(ctx, id, g, u)

	return
}

func cancel(ctx context.Context, id string, g getFunc, u updateFunc) error {
	// TODO 检查msg是否存在
	data, err := g(ctx, id)
	if err != nil {
		return err
	}
	if data.GetStatus() == meta.Status_Cancel {
		return errors.ErrMsgHasCancelled
	}
	if data.GetStatus() == meta.Status_Final {
		return errors.ErrMsgCantEdit
	}

	if err := cache.CancelMsg(ctx, id); err != nil {
		return err
	}

	go func() {
		// 异步将数据插入数据库
		go func() {
			var err error
			var count int
			// 重试机制，最多试两次
			retryFunc := func() error {
				if count > 2 {
					logrus.WithFields(logrus.Fields{
						"id": id,
					}).Errorf("cancel:取消消息请求执行成功，但是更新数据库字段时失败，请手动更新,error: %v", err)
					return nil
				}
				err = u(context.Background(), id)
				return err
			}
			back := backoff.NewServiceBackOff()
			err = backoff.Retry(retryFunc, back)
			if err != nil {
				logrus.Errorf("unexpected error: %s", err.Error())
			}
		}()
	}()

	return nil
}
