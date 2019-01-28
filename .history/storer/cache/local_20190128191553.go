/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : local.go
#   Created       : 2019/1/15 19:27
#   Last Modified : 2019/1/15 19:27
#   Describe      :
#
# ====================================================*/
package cache

import (
	"context"

	"github.com/hiruok/msg-pusher/storer"
	"github.com/opentracing/opentracing-go"
)

// AddLocalTemplate 将模板添加到本地缓存,
// 因为当前模板是不可变的所以，默认缓存1小时，如果以后变动频繁，可以适当降低缓存刷新时间。
func AddLocalTemplate(ctx context.Context, s string, v string) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan("AddLocalTemplate", opentracing.ChildOf(parentCtx))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	return storer.LocalCache.Put(ctx, s, []byte(v), 360)
}

// LocalTemplate 从本地缓存中取出模板
func LocalTemplate(ctx context.Context, s string) (string, error) {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan("LocalTemplate", opentracing.ChildOf(parentCtx))
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	b, err := storer.LocalCache.Get(ctx, s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
