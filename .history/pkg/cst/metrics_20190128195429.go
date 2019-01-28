/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : metrics.go
#   Created       : 2019/1/8 17:44
#   Last Modified : 2019/1/8 17:44
#   Describe      :
#
# ====================================================*/
package cst

var (
	// VarRequestCounts 记录每个api的访问时间和返回状态
	VarRequestCounts = []string{"api_name", "date", "status"}
	// VarRequestSummary 记录每个api的访问次数
	VarRequestSummary = []string{"api_name"}
)
