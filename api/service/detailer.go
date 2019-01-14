/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : getter.go
#   Created       : 2019/1/11 14:04
#   Last Modified : 2019/1/11 14:04
#   Describe      :
#
# ====================================================*/
package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/api/model"
	"uuabc.com/sendmsg/api/storer/cache"
	"uuabc.com/sendmsg/api/storer/db"
	cache2 "uuabc.com/sendmsg/pkg/cache"
	"uuabc.com/sendmsg/pkg/errors"
)

const (
	// 默认缓存时间为1天
	defaultTTL = 60 * 60 * 24
)

var DetailerImpl detailerImpl

type detailerImpl struct {
}

func (d detailerImpl) Detail(ctx context.Context, typeN, id string) (Marshaler, error) {
	res, err := d.detail(ctx, typeN, id)
	if err == nil {
		return res, nil
	}
	// 转换err类型
	if _, ok := err.(*errors.Error); !ok {
		if err == cache2.ErrCacheMiss {
			err = errors.ErrMsgNotFound
		} else {
			logrus.Errorf("idDetail,数据操作异常，error: %v", err)
			err = errors.NewError(
				10000000,
				err.Error(),
			)
		}
	}
	return nil, err
}

// DetailByPhonePage 直接数据库中取,不走缓存
func (d detailerImpl) DetailByPhonePage(ctx context.Context, mobile string, page int) ([]byte, error) {
	return nil, nil
}

// IDDetail 根据消息id查询数据,先从缓存中查取数据，如果不存在再去数据库中查
func (d detailerImpl) detail(ctx context.Context, typeN, id string) (Marshaler, error) {
	var res Marshaler
	switch typeN {
	case sms:
		res = &model.DbSms{}
	case wechat:
		res = &model.DbWeChat{}
	case email:
		res = &model.DbEmail{}
	default:
		return nil, errors.ErrMsgTypeNotFound
	}

	var data []byte
	var err error
	data, e1 := cache.LastestDetail(ctx, id)
	if e1 == nil {
		logrus.WithFields(logrus.Fields{
			"data": string(data),
			"id":   id,
		}).Infof("在lastest缓存中找到数据，直接返回结果")
	} else {
		data, err = cache.BaseDetail(ctx, id)
	}

	// 如果最新缓存不存在，则更新最新缓存和基础缓存
	if e1 == cache2.ErrCacheMiss {
		go func() {
			// 5秒内更新一次
			gErr := cache.LockID5s(context.Background(), id)
			if gErr != nil {
				logrus.WithField("id", id).Errorf("频繁请求更新不存在的key。")
				return
			}
			// 更新缓存
			dbRes, dbErr := d.idDetail(context.Background(), typeN, id)
			if dbErr != nil {
				logrus.Errorf("后台通过数据库更新cache失败，key:%s,error: %v", id, dbErr)
				return
			}
			cache.PutLastestCache(context.Background(), id, dbRes)
			cache.PutBaseCache(context.Background(), id, dbRes)
			logrus.WithField("id", id).Errorf("后台通过数据库添加cache成功")
		}()
	}

	if err != nil {
		return nil, err
	}
	err = res.Unmarshal(data)
	return res, err
}

func (detailerImpl) idDetail(ctx context.Context, typeN, id string) ([]byte, error) {
	switch typeN {
	case wechat:
		dt, err := db.WeChatDetailByID(ctx, id)
		if err != nil {
			return nil, err
		}
		return dt.Marshal()
	case email:
		dt, err := db.EmailDetailByID(ctx, id)
		if err != nil {
			return nil, err
		}
		return dt.Marshal()
	case sms:
		dt, err := db.SmsDetailByID(ctx, id)
		if err != nil {
			return nil, err
		}
		return dt.Marshal()
	default:
		return nil, errors.ErrMsgTypeNotFound
	}
}
