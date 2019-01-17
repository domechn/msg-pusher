/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : template.go
#   Created       : 2019/1/15 15:39
#   Last Modified : 2019/1/15 15:39
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"testing"

	"uuabc.com/sendmsg/pkg/pb/tpl"
)

func TestTemplateInsert(t *testing.T) {
	type args struct {
		ctx   context.Context
		templ *tpl.DBTemplate
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				ctx: context.Background(),
				templ: &tpl.DBTemplate{
					Type:     3,
					SimpleID: "test-id",
					Content: "<html><head><title>忘记密码验证码</title><style>p " +
						"{margin:0px;margin-bottom:5px;}</style></head><body>" +
						"<div style=\"margin-bottom:25px\"><p>亲爱的学生,</p></div>" +
						"<div style=\"margin-bottom:25px;\"><p>你正在使用找回密码功能.</p>" +
						"</div><div style=\"margin-bottom:25px;\"><p>验证码是 ${code}</p>" +
						"</div><div><p>Sincerely,</p><p>UUabc</p></div></body></html>",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TemplateInsert(tt.args.ctx, tt.args.templ)
			if (err != nil) != tt.wantErr {
				RollBack(got)
				t.Errorf("TemplateInsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			Commit(got)
		})
	}
}

func TestTemplateList(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "case1",
			args:    args{context.Background()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := TemplateList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TemplateList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
