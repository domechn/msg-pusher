/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : transaction.go
#   Created       : 2019/1/23 14:38
#   Last Modified : 2019/1/23 14:38
#   Describe      :
#
# ====================================================*/
package cache

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/hiruok/msg-pusher/pkg/cache/redis"
	"github.com/hiruok/msg-pusher/storer"
)

type Transaction struct {
	C *redis.Client
}

func NewTransaction() *Transaction {
	c := storer.Cache.Pipeline()
	return &Transaction{
		C: c,
	}
}

// PutBaseCache 底层缓存，跟数据库数据同步，不过期
func (t *Transaction) PutBaseCache(ctx context.Context, k string, v []byte) error {
	return put(ctx, "tx-PutBaseCache", base+k, v, 0, t.C)
}

// PutBaseTemplate 将添加的模板存入缓存中
func (t *Transaction) PutBaseTemplate(ctx context.Context, k string, v []byte) error {
	return put(ctx, "tx-PutBaseTemplate", template+k, v, 0, t.C)
}

// PutSendResult 修改发送结果，一般只有发送成功的情况才需要设置
func (t *Transaction) PutSendSuccess(ctx context.Context, k string) error {
	return put(ctx, "tx-PutSendSuccess", k+"_send", success, 0, t.C)
}

// RPushMsg 将消息存入list中，用于异步存入数据库
func (t *Transaction) RPushMsg(ctx context.Context, b []byte) error {
	return rPush(ctx, "tx-RPushMsg", msgDB, b, storer.Cache)
}

// LPopMsg 从msg队列中取一条数据
func (t *Transaction) LPopMsg() ([]byte, error) {
	return lPop(context.Background(), msgDB, t.C)
}

// Commit 提交事务
func (t *Transaction) Commit(ctx context.Context) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan("tx-Commit", opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "redis")
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	_, err := t.C.Exec()
	return err
}

// CommitParam 待参数的提交
func (t *Transaction) CommitParam(ctx context.Context) ([]interface{}, error) {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan("tx-Commit", opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "redis")
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	return t.C.Exec()
}

// Rollback 回滚事务
func (t *Transaction) Rollback(ctx context.Context) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan("tx-Rollback", opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "redis")
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	return t.C.Discard()
}

func (t *Transaction) Close() error {
	return t.C.Close()
}
