/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : response.go
#   Created       : 2019/1/7 15:50
#   Last Modified : 2019/1/7 15:50
#   Describe      :
#
# ====================================================*/
package sms

// The response code which stands for a sms is sent successfully.
const ResponseCodeOk = "OK"

// @see https://help.aliyun.com/document_detail/55284.html#出参列表
// The Response of sending sms API.
type Response struct {
	// The raw response from server.
	RawResponse []byte `json:"-"`
	/* Response body */
	RequestId string `json:"RequestId"`
	Code      string `json:"Code"`
	Message   string `json:"Message"`
	BizId     string `json:"BizId"`
}

func (m *Response) IsSuccessful() bool {
	return m.Code == ResponseCodeOk
}
