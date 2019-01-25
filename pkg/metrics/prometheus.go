/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : prometheus.go
#   Created       : 2019/1/8 17:44
#   Last Modified : 2019/1/8 17:44
#   Describe      :
#
# ====================================================*/
package metrics

import (
	"github.com/domgoer/msg-pusher/pkg/cst"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type Prometheus struct {
	counter Counter
	summary Summary
}

var (
	once sync.Once
	pro  Metrics
)

func GetPrometheus() Metrics {
	once.Do(func() {
		pro = &Prometheus{
			counter: newCountVec(),
			summary: newSummaryVec(),
		}
	})
	return pro
}

// Counter 返回counter
func (p *Prometheus) Counter() Counter {
	return p.counter
}

// Summary 返回summary
func (p *Prometheus) Summary() Summary {
	return p.summary
}

type CountVec struct {
	Count *prometheus.CounterVec
}

func newCountVec() *CountVec {
	c := new(CountVec)
	c.Count = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "sendmsg",
		Subsystem: "receiver",
		Name:      "http_request_counts",
		Help:      "Http Request Counts",
	}, cst.VarRequestCounts)
	return c
}

// Inc 给label +1
func (c *CountVec) Inc(labels map[string]string) {
	c.Count.With(labels).Inc()
}

type SummaryVec struct {
	Summary *prometheus.SummaryVec
}

func newSummaryVec() *SummaryVec {
	s := new(SummaryVec)
	s.Summary = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: "sendmsg",
		Subsystem: "receiver",
		Name:      "http_request_summary",
		Help:      "Http Request Summary",
	}, cst.VarRequestSummary)
	return s
}

// Observe 记录label的运行时间
func (s *SummaryVec) Observe(labels map[string]string, f float64) {
	s.Summary.With(labels).Observe(f)
}
