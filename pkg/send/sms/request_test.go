/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : request_test.go
#   Created       : 2019/1/7 16:41
#   Last Modified : 2019/1/7 16:41
#   Describe      :
#
# ====================================================*/
package sms

import (
	"fmt"
	"testing"
)

var (
	req *Request
)

func init() {
	req = NewRequest("1231231", "12421", "12421", "123123", "124w234")
	req.AccessKeyId = "213"
}

func TestRequest_Encode(t *testing.T) {
	type fields struct {
		AccessKeyId      string
		Timestamp        string
		Format           string
		SignatureMethod  string
		SignatureVersion string
		SignatureNonce   string
		Signature        string
		Action           string
		Version          string
		RegionId         string
		PhoneNumbers     string
		SignName         string
		TemplateCode     string
		TemplateParam    string
		SmsUpExtendCode  string
		OutId            string
	}
	type args struct {
		accessKeySecret string
		gatewayUrl      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			fields: fields{
				AccessKeyId:      req.AccessKeyId,
				Timestamp:        "1234",
				Format:           req.Format,
				SignatureNonce:   "nonce",
				SignatureMethod:  req.SignatureMethod,
				SignatureVersion: req.SignatureVersion,
				Signature:        req.Signature,
				Action:           req.Action,
				Version:          req.Version,
				RegionId:         req.RegionId,
				PhoneNumbers:     req.PhoneNumbers,
				SmsUpExtendCode:  req.SmsUpExtendCode,
				SignName:         req.SignName,
				TemplateCode:     req.TemplateCode,
				TemplateParam:    req.TemplateParam,
				OutId:            req.OutId,
			},
			args: args{
				accessKeySecret: "test",
				gatewayUrl:      "aabbb",
			},
			want: "aabbb?Signature=3NmFjycvE0ZBAm6qB914n7lngWs%3D&AccessKeyId=213&Action=SendSms&Format=json&OutId=124w234&PhoneNumbers=1231231&RegionId=cn-hangzhou&SignName=12421&SignatureMethod=HMAC-SHA1&SignatureNonce=nonce&SignatureVersion=1.0&SmsUpExtendCode=1234567&TemplateCode=12421&TemplateParam=123123&Timestamp=1234&Version=2017-05-25",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				AccessKeyId:      tt.fields.AccessKeyId,
				Timestamp:        tt.fields.Timestamp,
				Format:           tt.fields.Format,
				SignatureMethod:  tt.fields.SignatureMethod,
				SignatureVersion: tt.fields.SignatureVersion,
				SignatureNonce:   tt.fields.SignatureNonce,
				Signature:        tt.fields.Signature,
				Action:           tt.fields.Action,
				Version:          tt.fields.Version,
				RegionId:         tt.fields.RegionId,
				PhoneNumbers:     tt.fields.PhoneNumbers,
				SignName:         tt.fields.SignName,
				TemplateCode:     tt.fields.TemplateCode,
				TemplateParam:    tt.fields.TemplateParam,
				SmsUpExtendCode:  tt.fields.SmsUpExtendCode,
				OutId:            tt.fields.OutId,
			}
			got, err := r.Encode(tt.args.accessKeySecret, tt.args.gatewayUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
			if got != tt.want {
				t.Errorf("Request.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
