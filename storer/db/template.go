/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : template.go
#   Created       : 2019/1/15 15:39
#   Last Modified : 2019/1/15 15:39
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	"uuabc.com/sendmsg/pkg/pb/tpl"
)

// TemplateInsert 插入消息模板到数据库，如果唯一键重复返回键已存在错误
func TemplateInsert(ctx context.Context, templ *tpl.DBTemplate) (*sqlx.Tx, error) {
	return insert(ctx,
		"TemplateInsert",
		`INSERT INTO template (id,type,simple_id,content) VALUES (?,?,?,?)`,
		templ.Id,
		templ.Type,
		templ.SimpleID,
		templ.Content)
}

// TemplateList 获取所有模板
func TemplateList(ctx context.Context) (res []*tpl.DBTemplate, err error) {
	err = list(ctx, &res, "TemplateList", `SELECT * FROM template`)
	return
}
