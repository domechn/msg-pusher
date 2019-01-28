/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : log.go
#   Created       : 2019/1/8 16:46
#   Last Modified : 2019/1/8 16:46
#   Describe      :
#
# ====================================================*/
package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/hiruok/msg-pusher/pkg/ip"
)

// LoggingMiddleware 返回日志middleware
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logrus.WithFields(logrus.Fields{"method": r.Method, "path": r.URL.Path}).Info("Started request")

		realIP := ip.RealIP(r)
		fields := logrus.Fields{
			"request-id":  RequestID(r.Context()),
			"request-ip":  realIP,
			"method":      r.Method,
			"host":        r.Host,
			"request":     r.RequestURI,
			"remote-addr": r.RemoteAddr,
			"referer":     r.Referer(),
			"user-agent":  r.UserAgent(),
		}

		defer func(begin time.Time) {
			logrus.WithFields(fields).Infof("Completed handling request, It takes about: %v", time.Since(begin))
		}(time.Now())

		next.ServeHTTP(w, r)

	})
}
