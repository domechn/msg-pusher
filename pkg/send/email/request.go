/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : request.go
#   Created       : 2019/1/8 12:03
#   Last Modified : 2019/1/8 12:03
#   Describe      :
#
# ====================================================*/
package email

import (
	"strings"
)

// Message implements send.Message
type Request struct {
	to      string
	data    string
	subject string
}

func NewMessage(to, subject, data string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		data:    data,
	}
}

func (m *Request) To() string {
	return m.to
}

func (m *Request) Content() []byte {
	return []byte(m.data)
}

func (m *Request) textType(data string) string {
	if strings.Contains(data, "<html>") {
		return "html"
	}
	return "plain"
}
