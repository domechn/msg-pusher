/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : email.go
#   Created       : 2019/1/8 16:38
#   Last Modified : 2019/1/8 16:38
#   Describe      :
#
# ====================================================*/
package model

// "platform":3,
// "platform_key":"email_test",
// "server":2,
// "title":"测试邮件",
// "content": "好嗨哟，感觉人生已经到达了高潮",
// "template": "test.blade.php",
// "arguments":"{\"name\":\"乔治君\"}",
// "destination": "923143925@qq.com",
// "send_time":"2018-07-30T07:30:00+08:00",
// "type": 2
// EmailProducer 接收email信息
type EmailProducer struct {
	PlatForm    int    `json:"platform"`
	PlatFormKey string `json:"platform_key"`
	Server      int    `json:"server"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Template    string `json:"template"`
	Arguments   string `json:"arguments"`
	Destination string `json:"destination"`
	SendTime    string `json:"send_time"`
	Type        int    `json:"type"`
}

func (e *EmailProducer) Validated() error {
	return nil
}
