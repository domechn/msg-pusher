/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : send.go
#   Created       : 2019-01-07 14:20:59
#   Last Modified : 2019-01-07 14:20:59
#   Describe      :
#
# ====================================================*/
package email

import (
	"net/smtp"
	"testing"

	"github.com/domgoer/msgpusher/config"
	"github.com/domgoer/msgpusher/pkg/send"
)

var (
	ec  *config.Email
	cli *Client
)

func init() {
	err := config.Init("../../../conf.yaml")
	if err != nil {
		panic(err)
	}
	ec = config.EmailConf()
	cli = NewClient(Config{
		Username:   ec.Username,
		Password:   ec.Password,
		Host:       ec.Host,
		ServerAddr: ec.ServerAddr,
		TLS:        ec.TLS,
	})
}

func TestClient_Send(t *testing.T) {
	type fields struct {
		cfg  Config
		auth smtp.Auth
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
				cfg:  cli.cfg,
				auth: cli.auth,
			},
			args: args{
				msg: NewMessage("test@abc.com", "hello", "hello"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				cfg:  tt.fields.cfg,
				auth: tt.fields.auth,
			}
			if err := c.Send(tt.args.msg, tt.args.do); (err != nil) != tt.wantErr {
				t.Errorf("Client.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
