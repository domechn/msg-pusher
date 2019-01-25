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
func RateLimit(ctx context.Context, mobile string, sendTime time.Time) error {
	year, month, day := sendTime.Date()
	date := fmt.Sprintf("%d-%s-%d", year, month.String(), day)
	key := mobile + date
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		parentCtx := parentSpan.Context()
		span := opentracing.StartSpan("RateLimit", opentracing.ChildOf(parentCtx))
		ext.SpanKindRPCClient.Set(span)
		ext.PeerService.Set(span, "redis")
		span.SetTag("cache.key", key)
		span.SetTag("cache.type", "limit")
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	sendTime = sendTime.UTC()
	if sendTime.Before(time.Now()) {
		return nil
	}

	zeroStr := sendTime.Format("2006-01-02")
	zero, _ := time.Parse("2006-01-02", zeroStr)
	second := int(sendTime.Sub(zero).Seconds())
	ress, err := storer.Cache.ZRange(ctx, key, 0, 86400)
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
		t.C.ZAdd(context.Background(), key, second, []byte(strconv.Itoa(second)))
		if expire {
			ttl := sendTime.Sub(time.Now()).Seconds()
			fmt.Println(ttl)
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
