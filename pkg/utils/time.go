/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : time.go
#   Created       : 2019/1/24 16:19
#   Last Modified : 2019/1/24 16:19
#   Describe      :
#
# ====================================================*/
package utils

import (
	"strconv"
	"time"
)

const (
	layout = "2006-01-02 15:04:05"
)

// NowTimeStampStr 获取当前时间的时间戳
func NowTimeStampStr() string {
	now := time.Now().Unix()
	timeStamp := strconv.Itoa(int(now))
	return timeStamp
}

// MustISO8601StrToUTCStr 将iso8601格式的字符串换成utc时间的标准格式,
// 如果转换失败就返回layout
func MustISO8601StrToUTCStr(s string) string {
	t, err := time.Parse(layout, s)
	if err != nil {
		return layout
	}
	return t.UTC().String()
}
