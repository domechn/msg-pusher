/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : metrics.go
#   Created       : 2019/1/8 17:43
#   Last Modified : 2019/1/8 17:43
#   Describe      :
#
# ====================================================*/
package metrics

// Metrics 用于统计每个请求的状况
type Metrics interface {
	// 统计请求数
	Counter() Counter
	// 统计请求返回时间
	Summary() Summary
}

type Counter interface {
	Inc(labels map[string]string)
}

type Summary interface {
	Observe(labels map[string]string, f float64)
}

// TODO 支持更多metrics
func GetMetrics(metricsName string) Metrics {
	return GetPrometheus()
}
