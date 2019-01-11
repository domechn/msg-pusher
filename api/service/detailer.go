/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : detailer.go
#   Created       : 2019/1/11 14:04
#   Last Modified : 2019/1/11 14:04
#   Describe      :
#
# ====================================================*/
package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"time"
	"uuabc.com/sendmsg/api/model"
	"uuabc.com/sendmsg/api/storer/cache"
	"uuabc.com/sendmsg/api/storer/db"
	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/pkg/pb/meta"
)

const (
	// 默认缓存时间为1天
	defaultTTL = 60 * 60 * 24
)

var DetailerImpl detailerImpl

type detailerImpl struct {
}

func (d detailerImpl) Detail(ctx context.Context, typeN, id string) (*model.Response, error) {
	res, err := d.detail(ctx, typeN, id)
	if err == nil {
		return res, nil
	}
	// 转换err类型
	if _, ok := err.(*errors.Error); !ok {
		if err == sql.ErrNoRows {
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
func (d detailerImpl) DetailByPhonePage(ctx context.Context, mobile string, page int) (*model.Response, error) {
	return nil, nil
}

// IDDetail 根据消息id查询数据,先从缓存中查取数据，如果不存在再去数据库中查
func (d detailerImpl) detail(ctx context.Context, typeN, id string) (*model.Response, error) {
	// TODO 使用bitmap先查看消息是否存在，防止数据库被刷爆,id hash+bitmap
	data, err := cache.Detail(ctx, id)
	if err == nil {
		logrus.WithFields(logrus.Fields{
			"data": string(data),
			"id":   id,
		}).Infof("在缓存中找到数据，直接返回结果")
		ca := make(map[string]interface{})
		if err := json.Unmarshal(data, &ca); err != nil {
			logrus.Errorf("IDDetail: unexpected error :%v", err)
			goto GetDataInDB
		}
		return model.NewResponseDataKey("detail", ca), nil
	}
GetDataInDB:
	res, ttl, err := d.idDetail(ctx, typeN, id)
	if err != nil {
		return nil, err
	}
	logrus.WithFields(logrus.Fields{
		"id": id,
	}).Info("查询数据库获取msg具体信息成功")
	// 异步添加缓存
	go func() {
		if ttl <= 0 {
			return
		}
		// TODO 修改编码格式
		b, err := json.Marshal(res)
		if err != nil {
			return
		}
		if err = cache.StoreDetail(ctx, id, b, ttl); err == nil {
			logrus.WithFields(logrus.Fields{
				"data": string(data),
				"id":   id,
			}).Infof("数据异步添加缓存成功")
		}
	}()
	return model.NewResponseDataKey("detail", res), nil
}

func (detailerImpl) idDetail(ctx context.Context, typeN, id string) (res interface{}, ttl int32, err error) {
	switch typeN {
	case wechat:
		dt, er := db.WeChatDetailByID(ctx, id)
		ttl = calculateTTL(dt.Status, dt.SendTime)
		res = dt
		err = er
	case email:
		dt, er := db.EmailDetailByID(ctx, id)
		ttl = calculateTTL(dt.Status, dt.SendTime)
		res = dt
		err = er
	case sms:
		dt, er := db.SmsDetailByID(ctx, id)
		ttl = calculateTTL(dt.Status, dt.SendTime)
		res = dt
		err = er
	default:
		res, err = nil, errors.ErrMsgTypeNotFound
	}
	if err != nil {
		return nil, 0, err
	}

	return
}

// calculateTTL 根据消息的状态和发送时间计算存入缓存中的时间
func calculateTTL(status int, sendTime string) (ttl int32) {
	ttl = defaultTTL
	// 只要状态不是待发送，就将缓存设为1天
	if meta.Status(status) != meta.Status_Wait {
		return
	}
	// 如果消息是待发送状态 就将缓存有效时间设置为发送时间减去当前时间再减去10秒
	t, err := time.Parse("2006-01-02T15:04:05Z", sendTime)
	if err != nil {
		logrus.Errorf("calculateTTL:send_time转换异常，%s", sendTime)
		return -1
	}
	ttl = int32((t.Unix()-time.Now().Unix())/1000 - 10)
	if ttl <= 0 {
		ttl = defaultTTL
	}
	return
}
