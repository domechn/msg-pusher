/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : metrics.go
#   Created       : 2019/1/8 17:47
#   Last Modified : 2019/1/8 17:47
#   Describe      :
#
# ====================================================*/
package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hiruok/msg-pusher/pkg/cst"
	"github.com/hiruok/msg-pusher/pkg/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	counter metrics.Counter
	summary metrics.Summary
}

func NewMetrics(name string) *Metrics {
	m := metrics.GetMetrics(name)
	return &Metrics{
		counter: m.Counter(),
		summary: m.Summary(),
	}
}

func (m *Metrics) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var begin = time.Now()

		next.ServeHTTP(w, r)

		var err error
		if sw, ok := w.(*statusCodeResponseWriter); ok {
			if sw.status >= http.StatusBadRequest {
				err = fmt.Errorf("failed request")
			}
		}

		cstSum := cst.VarRequestSummary
		cstCount := cst.VarRequestCounts

		// 统计每台服务器的每个api的响应时间(前提是请求该服务器成功)
		m.summary.Observe(prometheus.Labels{cstSum[0]: r.RequestURI}, time.Since(begin).Seconds())

		// 统计api访问的数量
		status := "ok"
		if err != nil {
			status = "fail"
		}
		m.counter.Inc(prometheus.Labels{cstCount[0]: r.RequestURI, cstCount[1]: time.Now().Format("2006-01-02"), cstCount[2]: status})

	})
}
