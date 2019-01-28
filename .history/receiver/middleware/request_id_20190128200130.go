/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : request_id.go
#   Created       : 2019/1/8 16:49
#   Last Modified : 2019/1/8 16:49
#   Describe      :
#
# ====================================================*/
package middleware

import (
	"context"
	"net/http"

	"github.com/satori/go.uuid"
)

const (
	reqIDKey        = 0
	requestIDHeader = "X-Request-ID"
)

// RequestIDMiddleware 从req中获取request-id ， 如果request不存在就生成一个，并将request写入ResponseWriter
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		idHeader := r.Header.Get(requestIDHeader)

		if idHeader == "" {
			idHeader = uuid.NewV4().String()
			r.Header.Set(requestIDHeader, idHeader)
		}

		w.Header().Set(requestIDHeader, idHeader)

		ctx := r.Context()
		ctx = context.WithValue(ctx, reqIDKey, idHeader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequestID 从context中获取request-id
func RequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if requestID, ok := ctx.Value(reqIDKey).(string); ok {
		return requestID
	}

	return ""
}
