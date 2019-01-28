/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : msg.go
#   Created       : 2019/1/28 11:37
#   Last Modified : 2019/1/28 11:37
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"github.com/domgoer/msg-pusher/pkg/db"
	"github.com/domgoer/msg-pusher/storer"
	"reflect"
	"testing"

	"github.com/domgoer/msg-pusher/pkg/pb/meta"
)

func init() {
	storer.DB, _ = db.New(db.Config{
		URL:             "root:root@tcp(localhost:3306)/uuabc?charset=utf8&parseTime=True",
		MaxIdleConns:    10,
		MaxOpenConns:    20,
		ConnMaxLifetime: 3600,
	})
}

var dbMsg = &meta.DbMsg{
	Id:           "test",
	SubId:        "hello",
	SendTo:       "test-to",
	Content:      "test-content",
	Arguments:    "test-args",
	Template:     "test-template",
	Reserved:     "test-reserved",
	SendTime:     "2006-09-09 09:09:09",
	Type:         meta.Email,
	Server:       meta.AliYun,
	ResultStatus: meta.Success,
	Status:       meta.Final,
}

func TestUpdateAndInsertMsgBatch(t *testing.T) {
	type args struct {
		ctx context.Context
		ds  []*meta.DbMsg
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "batch_case_1",
			args: args{
				ctx: context.Background(),
				ds: []*meta.DbMsg{
					dbMsg,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateAndInsertMsgBatch(tt.args.ctx, tt.args.ds); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAndInsertMsgBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDetailByToAndPage(t *testing.T) {
	type args struct {
		ctx  context.Context
		to   string
		page int
	}
	tests := []struct {
		name    string
		args    args
		want    []*meta.DbMsg
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetailByToAndPage(tt.args.ctx, tt.args.to, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetailByToAndPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetailByToAndPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetailByKey(t *testing.T) {
	type args struct {
		ctx  context.Context
		key  string
		page int
	}
	tests := []struct {
		name    string
		args    args
		want    []*meta.DbMsg
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetailByKey(tt.args.ctx, tt.args.key, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetailByKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetailByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaitingMsgByKey(t *testing.T) {
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WaitingMsgByKey(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("WaitingMsgByKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WaitingMsgByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
