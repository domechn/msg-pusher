/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : response.go
#   Created       : 2019/1/8 10:32
#   Last Modified : 2019/1/8 10:32
#   Describe      :
#
# ====================================================*/
package wechat

// Response 微信接口返回值
type Response struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
