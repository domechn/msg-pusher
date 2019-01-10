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
package mail

// Message implements send.Message
type Request struct {
	to   string
	data []byte
}

func NewMessage(to string, data []byte) *Request {
	return &Request{
		to:   to,
		data: data,
	}
}

func (m *Request) To() string {
	return m.to
}

func (m *Request) Content() []byte {
	return m.data
}
