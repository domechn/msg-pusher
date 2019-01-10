/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : wechat.go
#   Created       : 2019/1/8 16:39
#   Last Modified : 2019/1/8 16:39
#   Describe      :
#
# ====================================================*/
package model

// "platform":3,
// "touser": "oAdVsv5nQ-qtBF0F5WGU-xPcrpGY",
// "template_id": "RlJfVX1SCBW2ncbIblbOE_8PaOUyoxBmr2MKnjzcY80",
// "url": "",
// "data":"{\"first\": {\"value\":\"亲爱的George家长，UU哥提示您：本次课程已经结束，本节课配有课后作业哦，作业信息如下：\",\"color\":\"#173177\"},\"keyword1\":{\"value\":\"Hello\",\"color\":\"#173177\"},\"keyword2\": {\"value\":\"World\",\"color\":\"#173177\"},\"remark\":{\"value\":\"请提醒小朋友前往学习记录完成课后作业，完成作业还会有额外奖励哦。。如有任何问题可微信对话框留言或拨打服务热线：4001636161\",\"color\":\"#173177\"}}",
// "send_time":"0",
// "type": 2
// WeChatProducer 接收微信消息
type WeChatProducer struct {
	PlatForm   int    `json:"platform"`
	ToUser     string `json:"touser"`
	TemplateID string `json:"template_id"`
	URL        string `json:"url"`
	Data       string `json:"data"`
	SendTime   string `json:"send_time"`
	Type       int    `json:"type"`
}

func (w *WeChatProducer) Validated() error {
	return nil
}
