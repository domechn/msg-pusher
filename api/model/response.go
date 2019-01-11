/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : response.go
#   Created       : 2019/1/11 14:07
#   Last Modified : 2019/1/11 14:07
#   Describe      :
#
# ====================================================*/
package model

type Response struct {
	Errcode int         `json:"errcode"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// NewResponseData return {"errcode":0,"msg":"success","data"=data}
func NewResponseData(data interface{}) *Response {
	return &Response{
		Errcode: 0,
		Msg:     "success",
		Data:    data,
	}
}

// NewResponseDataKey return {"errcode":0,"msg":"success","data":{key:data}}
func NewResponseDataKey(key string, data interface{}) *Response {
	return &Response{
		Errcode: 0,
		Msg:     "success",
		Data: map[string]interface{}{
			key: data,
		},
	}
}
