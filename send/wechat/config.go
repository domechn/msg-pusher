/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : config.go
#   Created       : 2019/1/7 19:56
#   Last Modified : 2019/1/7 19:56
#   Describe      :
#
# ====================================================*/
package wechat

type Config struct {
	CacheAddrs []string
	CachePwd   string

	APPId     string
	APPSecret string
}
