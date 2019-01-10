/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : request.go
#   Created       : 2019/1/7 20:14
#   Last Modified : 2019/1/7 20:14
#   Describe      :
#
# ====================================================*/
package wechat

import (
	"encoding/json"
)

type Request struct {
	ToUser     string `json:"touser"`
	TemplateId string `json:"template_id"`
	URL        string `json:"url"`
	Data       string `json:"data"`
}

func NewRequest(to, templateId, url string, data []byte) *Request {
	return &Request{
		ToUser:     to,
		TemplateId: templateId,
		URL:        url,
		Data:       string(data),
	}
}

func (r *Request) Content() []byte {
	b, _ := json.Marshal(r)
	return b
}

func (r *Request) To() string {
	return r.ToUser
}
