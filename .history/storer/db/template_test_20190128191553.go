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

	"github.com/hiruok/msg-pusher/pkg/pb/tpl"
	"github.com/hiruok/msg-pusher/storer"
	"github.com/jmoiron/sqlx"
)

var dbtm = &tpl.DBTemplate{
	Id:       "test-template-id",
	Type:     1,
	SimpleID: "test-simple-id",
	Content:  "test-content",
}

func TestTemplateInsert(t *testing.T) {
	type args struct {
		ctx   context.Context
		templ *tpl.DBTemplate
	}
	tests := []struct {
		name    string
		args    args
		want    *sqlx.Tx
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "insert_case_1",
			args: args{
				ctx:   context.Background(),
				templ: dbtm,
			},
		}, {
			name: "inser_case_2",
			args: args{
				ctx:   context.Background(),
				templ: dbtm,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TemplateInsert(tt.args.ctx, tt.args.templ)
			if err != nil {
				RollBack(got)
			} else {
				Commit(got)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("TemplateInsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
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
		wantRes []*tpl.DBTemplate
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "list_case_1",
			args: args{
				ctx: context.Background(),
			},
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

func TestDeleteTemp(t *testing.T) {
	storer.DB.Exec("DELETE FROM template WHERE id = ?", dbtm.Id)
}
