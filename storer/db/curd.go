/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : curd.go
#   Created       : 2019/1/21 14:40
#   Last Modified : 2019/1/21 14:40
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"uuabc.com/sendmsg/storer"
)

// query 获取单条数据
func query(ctx context.Context, out interface{}, typeN, sql string, args ...interface{}) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan(typeN, opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "mysql")
		span.SetTag("sql.query", sql)
		span.SetTag("sql.param", args)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	return storer.DB.GetContext(ctx, out, sql, args...)
}

// list 获取集合数据
func list(ctx context.Context, out interface{}, typeN, sql string, args ...interface{}) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan(typeN, opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "mysql")
		span.SetTag("sql.query", sql)
		span.SetTag("sql.param", args)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	return storer.DB.SelectContext(ctx, out, sql, args...)
}

// update 更新数据
func update(ctx context.Context, typeN, sqlStr string, args ...interface{}) (*sqlx.Tx, error) {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan(typeN, opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "mysql")
		span.SetTag("sql.update", sqlStr)
		span.SetTag("sql.param", args)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	tx, err := storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, sqlStr)
	if err != nil {
		return tx, err
	}
	defer stmt.Close()
	var res sql.Result
	res, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return tx, err
	}
	if i, _ := res.RowsAffected(); i == 0 {
		return tx, ErrNoRowsEffected
	}
	return tx, nil
}

func insert(ctx context.Context, typeN, sqlStr string, args ...interface{}) (tx *sqlx.Tx, err error) {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan(typeN, opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "mysql")
		span.SetTag("sql.insert", sqlStr)
		span.SetTag("sql.param", args)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	tx, err = storer.DB.Beginx()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, sqlStr)
	if err != nil {
		return tx, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, args...)
	if err != nil && err.(*mysql.MySQLError).Number == 1062 {
		err = ErrUniqueKeyExsits
	}
	return tx, err
}
