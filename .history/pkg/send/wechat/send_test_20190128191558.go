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
	"testing"

	"github.com/hiruok/msg-pusher/config"
	"github.com/hiruok/msg-pusher/pkg/cache"
	"github.com/hiruok/msg-pusher/pkg/cache/redis"
	"github.com/hiruok/msg-pusher/pkg/send"
)

var (
	wc  *config.WeChat
	rc  *config.Redis
	cli *Client
)

func init() {
	err := config.Init("../../../conf.yaml")
	if err != nil {
		panic(err)
	}
	wc = config.WeChatConf()
	rc = config.RedisConf()
	rcli, _ := redis.NewClient([]string{"127.0.0.1:6379"}, "")
	cli = NewClient(Config{
		APPId:     wc.AppId,
		APPSecret: wc.AppSecret,
	}, rcli)
}

func TestClient_Send(t *testing.T) {
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
			name: "send_case_1",
			fields: fields{
				httpCli: cli.httpCli,
				cached:  cli.cached,
				cfg:     cli.cfg,
			},
			args: args{
				msg: NewRequest("abc", "abc", "abc", []byte("abc")),
			},
			wantErr: true,
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
