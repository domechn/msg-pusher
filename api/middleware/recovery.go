/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : recovery.go
#   Created       : 2019/1/8 16:48
#   Last Modified : 2019/1/8 16:48
#   Describe      :
#
# ====================================================*/
package middleware

import (
	"net/http"
)

// RecoveryMiddleware 防止程序因为不可预测的原因退出
func RecoveryMiddleware(returnFunc func(w http.ResponseWriter, r *http.Request, err interface{})) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					returnFunc(w, r, err)
				}
			}()

			handler.ServeHTTP(w, r)
		})
	}
}
