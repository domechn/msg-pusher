/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : wechat.go
#   Created       : 2019/1/11 16:58
#   Last Modified : 2019/1/11 16:58
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"testing"

	"github.com/domgoer/msgpusher/pkg/pb/meta"
	"github.com/domgoer/msgpusher/storer"
	"github.com/jmoiron/sqlx"
)

var dbw = &meta.DbWeChat{
	Id:          "test-test-test-wechat",
	Platform:    1,
	PlatformKey: "test-platform",
	Touser:      "to-user-test",
	Status:      1,
	Type:        1,
	Template:    "test-templ",
	Content:     "test",
	Arguments:   "test",
	SendTime:    "2019-09-09 09:09:09",
	Url:         "test-url",
}

func TestWeChatInsert(t *testing.T) {
	type args struct {
		ctx    context.Context
		wechat *meta.DbWeChat
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
				ctx:    context.Background(),
				wechat: dbw,
			},
		}, {
			name: "insert_case_2",
			args: args{
				ctx:    context.Background(),
				wechat: dbw,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := WeChatInsert(tt.args.ctx, tt.args.wechat)
			if err != nil {
				RollBack(tx)
			} else {
				Commit(tx)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("WeChatInsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestWeChatEdit(t *testing.T) {
	dbw.Content = "test-c"
	type args struct {
		ctx context.Context
		w   *meta.DbWeChat
	}
	tests := []struct {
		name    string
		args    args
		want    *sqlx.Tx
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "edit_case_1",
			args: args{
				ctx: context.Background(),
				w:   dbw,
			},
		}, {
			name: "edit_case_2",
			args: args{
				ctx: context.Background(),
				w:   dbw,
			},
			wantErr: true,
		}, {
			name: "edit_case_3",
			args: args{
				ctx: context.Background(),
				w: &meta.DbWeChat{
					Id: "wu",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WeChatEdit(tt.args.ctx, tt.args.w)
			if err != nil {
				RollBack(got)
			} else {
				Commit(got)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("WeChatEdit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestWeChatDetailByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *meta.DbWeChat
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "detailD_case_1",
			args: args{
				ctx: context.Background(),
				id:  dbw.Id,
			},
		}, {
			name: "detailD_case_2",
			args: args{
				ctx: context.Background(),
				id:  "wu",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := WeChatDetailByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("WeChatDetailByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestWeChatUpdateSendResult(t *testing.T) {
	dbw.Status = 2
	type args struct {
		ctx context.Context
		w   *meta.DbWeChat
	}
	tests := []struct {
		name    string
		args    args
		want    *sqlx.Tx
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "update_case_1",
			args: args{
				ctx: context.Background(),
				w:   dbw,
			},
		}, {
			name: "update_case_2",
			args: args{
				ctx: context.Background(),
				w:   dbw,
			},
			wantErr: true,
		}, {
			name: "update_case_3",
			args: args{
				ctx: context.Background(),
				w: &meta.DbWeChat{
					Id: "wu",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WeChatUpdateSendResult(tt.args.ctx, tt.args.w)
			if err != nil {
				RollBack(got)
			} else {
				Commit(got)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("WeChatUpdateSendResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestWeChatCancelMsgByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *sqlx.Tx
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "cancel_case_1",
			args: args{
				ctx: context.Background(),
				id:  dbw.Id,
			},
		}, {
			name: "cancel_case_2",
			args: args{
				ctx: context.Background(),
				id:  "wu",
			},
			wantErr: true,
		}, {
			name: "cancel_casse_3",
			args: args{
				ctx: context.Background(),
				id:  dbw.Id,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WeChatCancelMsgByID(tt.args.ctx, tt.args.id)
			if err != nil {
				RollBack(got)
			} else {
				Commit(got)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("WeChatCancelMsgByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDeleteW(t *testing.T) {
	storer.DB.Exec("DELETE FROM wechats WHERE id= ?", dbw.Id)
}

func TestWeChatUpdateAndInsertBatch(t *testing.T) {
	type args struct {
		ctx context.Context
		dw  []*meta.DbWeChat
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_batch_case1",
			args: args{
				ctx: context.Background(),
				dw: []*meta.DbWeChat{
					dbw,
				},
			},
		},
	}
	for _, tt := range tests {
		dbw.Type = 1
		dbw.Version = 1
		t.Run(tt.name, func(t *testing.T) {
			if err := WeChatUpdateAndInsertBatch(tt.args.ctx, tt.args.dw); (err != nil) != tt.wantErr {
				t.Errorf("WeChatUpdateAndInsertBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
