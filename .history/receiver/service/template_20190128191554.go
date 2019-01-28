/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : template.go
#   Created       : 2019/1/15 15:21
#   Last Modified : 2019/1/15 15:21
#   Describe      :
#
# ====================================================*/
package service

import (
	"context"

	"github.com/hiruok/msg-pusher/pkg/errors"
	"github.com/hiruok/msg-pusher/pkg/pb/tpl"
	"github.com/hiruok/msg-pusher/storer/cache"
	"github.com/hiruok/msg-pusher/storer/db"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

var TemplateImpl templateImpl

type templateImpl struct {
}

// AddTemplate 添加模板
func (t templateImpl) AddTemplate(ctx context.Context, a *tpl.TemplateAdder) (string, error) {
	uid := uuid.NewV4().String()

	tx, err := t.add(ctx, uid, a)
	if err != nil {
		db.RollBack(tx)
		return "", err
	}
	err = db.Commit(tx)
	if err == nil {
		// 更新本地缓存
		cache.AddLocalTemplate(ctx, a.SimpleID, a.Content)
	}
	return uid, nil
}

func (templateImpl) add(ctx context.Context, id string, a *tpl.TemplateAdder) (tx *sqlx.Tx, err error) {
	tx, err = db.TemplateInsert(ctx, &tpl.DBTemplate{
		Id:       id,
		SimpleID: a.SimpleID,
		Type:     a.Type,
		Content:  a.Content,
	})
	if err != nil {
		if err == db.ErrUniqueKeyExsits {
			err = errors.ErrTemplateIsExisted
		}
		return
	}
	err = cache.PutBaseTemplate(ctx, a.SimpleID, []byte(a.Content))
	return
}
