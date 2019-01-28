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

type APIName struct {
}

var (
	VarRequestCounts  = []string{"api_name", "date", "status"}
	VarRequestSummary = []string{"api_name"}
)
