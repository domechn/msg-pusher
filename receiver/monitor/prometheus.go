/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : prometheus.go
#   Created       : 2019/1/8 17:58
#   Last Modified : 2019/1/8 17:58
#   Describe      :
#
# ====================================================*/
package monitor

import (
	"fmt"
	"github.com/domgoer/msgpusher/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
)

// StartMetrics prometheus
func StartMetrics(addr string) (err error) {
	if addr == "" {
		return
	}

	pro := metrics.GetPrometheus()

	count := pro.Counter().(*metrics.CountVec).Count
	sum := pro.Summary().(*metrics.SummaryVec).Summary

	prometheus.Register(count)
	prometheus.Register(sum)

	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())

	listen, err := net.Listen("tcp", addr)

	if err != nil {
		return fmt.Errorf("failed metrics init: %v", err)
	}

	go http.Serve(listen, http.DefaultServeMux)
	return
}
