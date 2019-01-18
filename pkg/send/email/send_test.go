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

	"uuabc.com/sendmsg/config"
	"uuabc.com/sendmsg/pkg/send"
)

var (
	ec *config.Email
)

func init() {
	err := config.Init("../../../conf.yaml")
	if err != nil {
		panic(err)
	}
	ec = config.EmailConf()
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
			name: "case1",
			fields: fields{
				cfg: Config{
					ServerAddr: "smtp.qq.com:25",
					Username:   ec.Username,
					Password:   ec.Password,
					Host:       "smtp.qq.com",
				},
				auth: smtp.PlainAuth("", ec.Username, ec.Password, ec.Host),
			},
			args: args{
				msg: NewMessage("814172254@qq.com;mengcheng.dou@uuabc.com",
					"test",
					"Hello World",
					false),
			},
		}, {
			name: "caseTls",
			fields: fields{
				cfg: Config{
					ServerAddr: "smtp.qq.com:465",
					Username:   ec.Username,
					Password:   ec.Password,
					Host:       "smtp.qq.com",
				},
				auth: smtp.PlainAuth("", ec.Username, ec.Password, ec.Host),
			},
			args: args{
				msg: NewMessage("814172254@qq.com;mengcheng.dou@uuabc.com",
					"test",
					"Hello World",
					true),
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
