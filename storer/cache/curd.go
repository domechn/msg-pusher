/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : curd.go
#   Created       : 2019/1/21 15:13
#   Last Modified : 2019/1/21 15:13
#   Describe      :
#
# ====================================================*/
package cache

import (
	"context"

	"github.com/domgoer/msgpusher/pkg/cache/redis"
	"github.com/domgoer/msgpusher/storer"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func get(ctx context.Context, typeN, k string) ([]byte, error) {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan(typeN, opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "redis")
		span.SetTag("cache.type", "get")
		span.SetTag("cache.key", k)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	return storer.Cache.Get(ctx, k)
}

func put(ctx context.Context, typeN, k string, b []byte, ttl int64, c *redis.Client) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan(typeN, opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "redis")
		span.SetTag("cache.type", "put")
		span.SetTag("cache.key", k)
		span.SetTag("cache.value", string(b))
		span.SetTag("cache.ttl", ttl)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	return c.Put(ctx, k, b, ttl)
}

func limit(ctx context.Context, typeN, k string, ttl int64) (int64, error) {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan(typeN, opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "redis")
		span.SetTag("cache.key", k)
		span.SetTag("cache.type", "limit")
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	res, err := storer.Cache.Incr(ctx, k)
	if err != nil {
		return res, err
	}
	if res == 1 {
		return res, storer.Cache.Expire(ctx, k, ttl)
	}
	return res, nil
}

func add(ctx context.Context, typeN, k string, v []byte, ttl int64) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan(typeN, opentracing.ChildOf(parentCtx))
		ext.PeerService.Set(span, "redis")
		ext.SpanKindRPCClient.Set(span)
		span.SetTag("cache.type", "add")
		span.SetTag("cache.key", k)
		span.SetTag("cache.value", string(v))
		span.SetTag("cache.ttl", ttl)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	return storer.Cache.Add(ctx, k, v, ttl)
}

func del(ctx context.Context, typeN, k string) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan(typeN, opentracing.ChildOf(parentCtx))
		ext.PeerService.Set(span, "redis")
		ext.SpanKindRPCClient.Set(span)
		span.SetTag("cache.type", "del")
		span.SetTag("cache.id", k)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	return storer.Cache.Del(ctx, k)
}

func rPush(ctx context.Context, typeN, k string, v []byte, c *redis.Client) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan(typeN, opentracing.ChildOf(parentCtx))
		ext.PeerService.Set(span, "redis")
		ext.SpanKindRPCClient.Set(span)
		span.SetTag("cache.type", "rPush")
		span.SetTag("cache.id", k)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	return c.RPush(ctx, k, v)
}

func lPop(ctx context.Context, k string, c *redis.Client) ([]byte, error) {
	return c.LPop(ctx, k)
}
