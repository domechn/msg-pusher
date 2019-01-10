/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : sms.go
#   Created       : 2019/1/8 16:36
#   Last Modified : 2019/1/8 16:36
#   Describe      :
#
# ====================================================*/
package model

// "platform":3,
// "platform_key":"message_test",
// "server":2,
// "content": "验证码885588，您正在进行身份验证，打死不要告诉别人哦！",
// "template": "SMS_130990029",
// "arguments":"{\"code\":885588}",
// "mobile": "18516051096",
// "send_time":"2018-07-30T07:30:00+08:00",
// "type": 2
// SmsProducer 接收短信消息
type SmsProducer struct {
	PlatForm    int    `json:"platform"`
	PlatFormKey string `json:"platform_key"`
	Server      int    `json:"server"`
	Content     string `json:"content"`
	Template    string `json:"template"`
	Arguments   string `json:"arguments"`
	Mobile      string `json:"mobile"`
	SendTime    string `json:"send_time"`
	Type        int    `json:"type"`
}

func (s *SmsProducer) Validated() error {

	return nil
}
