/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : limit.go
#   Created       : 2019/1/25 16:09
#   Last Modified : 2019/1/25 16:09
#   Describe      : 限制短信发送频次
#
# ====================================================*/
package cache

import (
	"context"
	"fmt"
	"github.com/domgoer/msg-pusher/pkg/errors"
	"github.com/domgoer/msg-pusher/storer"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"math"
	"strconv"
	"time"
)

// RateLimit 限速器，如果超过流量返回失败，sendDate格式2006-01-02,second为当天0点到该时的秒数
func RateLimit(ctx context.Context, toUser string, sendTime time.Time) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan("RateLimit", opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "redis")
		span.SetTag("cache.key", toUser)
		span.SetTag("cache.type", "limit")
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	if sendTime.Add(time.Minute).Before(time.Now().UTC()) {
		return nil
	}
	zeroStr := sendTime.Format("2006-01-02")
	key := toUser + zeroStr
	zero, _ := time.Parse("2006-01-02", zeroStr)
	second := int(sendTime.Sub(zero).Seconds())
	ress, err := storer.Cache.SMembers(ctx, key)
	if err != nil {
		return err
	}
	var sends []int
	for _, v := range ress {
		if vv, err := strconv.Atoi(v); err == nil {
			sends = append(sends, vv)
		}
	}
	if len(sends) >= 10 {
		return errors.ErrMsg1DayLimit
	}

	var expire bool

	if len(sends) == 0 {
		expire = true
	}

	var hourCount int

	for _, s := range sends {
		if math.Abs(float64(second-s)) < 60 {
			return errors.ErrMsg1MinuteLimit
		}
		if math.Abs(float64(second-s)) < 3600 {
			hourCount++
		}
	}
	if hourCount > 5 {
		return errors.ErrMsg1HourLimit
	}

	go func() {
		t := NewTransaction()
		defer t.Close()
		fmt.Println(t.C.Append(context.Background(), key, []byte(strconv.Itoa(second))))
		if expire {
			ttl := zero.Add(time.Hour * 24).Sub(time.Now()).Seconds()
			if ttl <= 0 {
				t.Rollback(context.Background())
				return
			}
			t.C.Expire(context.Background(), key, int64(ttl))
		}
		t.Commit(context.Background())
	}()
	return nil
}

// RemoveLimit 移除特定时间的限制
func RemoveLimit(ctx context.Context, toUser string, sendTime time.Time) error {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan("RemoveLimit", opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "redis")
		span.SetTag("cache.key", toUser)
		span.SetTag("cache.type", "remove-limit")
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	sendTime = sendTime.UTC()
	if sendTime.Before(time.Now()) {
		return nil
	}
	zeroStr := sendTime.Format("2006-01-02")
	key := toUser + zeroStr
	zero, _ := time.Parse("2006-01-02", zeroStr)
	second := int(sendTime.Sub(zero).Seconds())

	return storer.Cache.SRem(ctx, key, []byte(strconv.Itoa(second)))
}
