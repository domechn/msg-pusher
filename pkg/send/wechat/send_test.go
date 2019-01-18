/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : send.go
#   Created       : 2019/1/7 14:34
#   Last Modified : 2019/1/7 14:34
#   Describe      :
#
# ====================================================*/
package wechat

import (
	"net/http"
	"reflect"
	"testing"

	"uuabc.com/sendmsg/config"
	"uuabc.com/sendmsg/pkg/cache"
	"uuabc.com/sendmsg/pkg/cache/redis"
	"uuabc.com/sendmsg/pkg/send"
)

var (
	wc *config.WeChat
	rc *config.Redis
)

func init() {
	err := config.Init("../../../conf.yaml")
	if err != nil {
		panic(err)
	}
	wc = config.WeChatConf()
	rc = config.RedisConf()
}

func TestClient_accessTokenData(t *testing.T) {
	caCli, err := redis.NewClient(rc.Addrs, rc.Password)
	if err != nil {
		t.Error(err)
		return
	}
	type fields struct {
		httpCli *http.Client
		cached  cache.Cache
		cfg     Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "case1",
			fields: fields{
				httpCli: &http.Client{},
				cached:  caCli,
				cfg: Config{
					APPId:     wc.AppId,
					APPSecret: wc.AppSecret,
				},
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpCli: tt.fields.httpCli,
				cached:  tt.fields.cached,
				cfg:     tt.fields.cfg,
			}
			_, err := c.accessTokenData()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.accessTokenData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_requestAccessToken(t *testing.T) {
	type fields struct {
		httpCli *http.Client
		cached  cache.Cache
		cfg     Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpCli: tt.fields.httpCli,
				cached:  tt.fields.cached,
				cfg:     tt.fields.cfg,
			}
			got, err := c.requestAccessToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.requestAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.requestAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_token(t *testing.T) {
	type fields struct {
		httpCli *http.Client
		cached  cache.Cache
		cfg     Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpCli: tt.fields.httpCli,
				cached:  tt.fields.cached,
				cfg:     tt.fields.cfg,
			}
			got, err := c.token()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.token() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.token() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_storeToken(t *testing.T) {
	type fields struct {
		httpCli *http.Client
		cached  cache.Cache
		cfg     Config
	}
	type args struct {
		v string
		e int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpCli: tt.fields.httpCli,
				cached:  tt.fields.cached,
				cfg:     tt.fields.cfg,
			}
			if err := c.storeToken(tt.args.v, tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("Client.storeToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_lockTokenGet(t *testing.T) {
	type fields struct {
		httpCli *http.Client
		cached  cache.Cache
		cfg     Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpCli: tt.fields.httpCli,
				cached:  tt.fields.cached,
				cfg:     tt.fields.cfg,
			}
			if err := c.lockTokenGet(); (err != nil) != tt.wantErr {
				t.Errorf("Client.lockTokenGet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_unLockTokenGet(t *testing.T) {
	type fields struct {
		httpCli *http.Client
		cached  cache.Cache
		cfg     Config
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpCli: tt.fields.httpCli,
				cached:  tt.fields.cached,
				cfg:     tt.fields.cfg,
			}
			c.unLockTokenGet()
		})
	}
}

func TestClient_Send(t *testing.T) {
	caCli, err := redis.NewClient(rc.Addrs, rc.Password)
	if err != nil {
		t.Error(err)
		return
	}
	type fields struct {
		httpCli *http.Client
		cached  cache.Cache
		cfg     Config
	}
	type args struct {
		msg send.Message
		do  send.DoRes
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			fields: fields{
				httpCli: &http.Client{},
				cached:  caCli,
				cfg: Config{
					APPSecret: wc.AppSecret,
					APPId:     wc.AppId,
				},
			},
			args: args{
				msg: NewRequest(
					"oAdVsv5nQ-qtBF0F5WGU-xPcrpGY",
					"RlJfVX1SCBW2ncbIblbOE_8PaOUyoxBmr2MKnjzcY80",
					"",
					[]byte("{\"first\": {\"value\":\"亲爱的George家长，UU哥提示您：本次课程已经结束，本节课配有课后作业哦，"+
						"作业信息如下：\",\"color\":\"#173177\"},\"keyword1\":"+
						"{\"value\":\"Hello\",\"color\":\"#173177\"},\"keyword2\": "+
						"{\"value\":\"World\",\"color\":\"#173177\"},\"remark\":"+
						"{\"value\":\"请提醒小朋友前往学习记录完成课后作业，完成作业还会有额外奖励哦。"+
						"如有任何问题可微信对话框留言或拨打服务热线：4001636161\",\"color\":\"#173177\"}}"),
				),
				do: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				httpCli: tt.fields.httpCli,
				cached:  tt.fields.cached,
				cfg:     tt.fields.cfg,
			}
			if err := c.Send(tt.args.msg, tt.args.do); (err != nil) != tt.wantErr {
				t.Errorf("Client.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
