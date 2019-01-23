/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : putter.go
#   Created       : 2019/1/11 17:03
#   Last Modified : 2019/1/11 17:03
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/storer"
)

var dbe = &meta.DbEmail{
	Id:          "test-test-test-email",
	Platform:    1,
	PlatformKey: "test-platform",
	Title:       "hello",
	Content:     "test-content",
	Destination: "heelo@abc.com",
	Type:        1,
	Template:    "1-1",
	Arguments:   "test",
	SendTime:    "2018-09-09 09:09:09",
	Status:      1,
}

func TestEmailInsert(t *testing.T) {
	type args struct {
		ctx context.Context
		e   *meta.DbEmail
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
				ctx: context.Background(),
				e:   dbe,
			},
		}, {
			name: "insert_case_2",
			args: args{
				ctx: context.Background(),
				e:   dbe,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EmailInsert(tt.args.ctx, tt.args.e)
			if err != nil {
				RollBack(got)
			} else {
				Commit(got)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("EmailInsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEmailDetailByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *meta.DbEmail
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "detailD_case_1",
			args: args{
				ctx: context.Background(),
				id:  dbe.Id,
			},
		}, {
			name: "detailD_case_2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := EmailDetailByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("EmailDetailByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEmailEdit(t *testing.T) {
	dbe.Destination = "hello@aaa.com"
	type args struct {
		ctx context.Context
		e   *meta.DbEmail
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
				e:   dbe,
			},
		}, {
			name: "edit_case_2",
			args: args{
				ctx: context.Background(),
				e: &meta.DbEmail{
					Id: "test11",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EmailEdit(tt.args.ctx, tt.args.e)
			if err != nil {
				RollBack(got)
			} else {
				Commit(got)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("EmailEdit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEmailUpdateSendResult(t *testing.T) {
	dbe.ResultStatus = 3
	type args struct {
		ctx context.Context
		e   *meta.DbEmail
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
				e:   dbe,
			},
		}, {
			name: "update_case_2",
			args: args{
				ctx: context.Background(),
				e: &meta.DbEmail{
					Id: "ttt",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EmailUpdateSendResult(tt.args.ctx, tt.args.e)
			if err != nil {
				RollBack(got)
			} else {
				Commit(got)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("EmailUpdateSendResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEmailCancelMsgByID(t *testing.T) {
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
				id:  dbe.Id,
			},
		}, {
			name: "caccel_case_2",
			args: args{
				ctx: context.Background(),
				id:  "wu",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EmailCancelMsgByID(tt.args.ctx, tt.args.id)
			if err != nil {
				RollBack(got)
			} else {
				Commit(got)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("EmailCancelMsgByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEmailUpdateBatch(t *testing.T) {
	s := time.Now().Format("2006-01-02 13:04:05")
	dbe.CreatedAt = s
	dbe.UpdatedAt = s
	dbe.Type = 7

	type args struct {
		ctx context.Context
		es  []*meta.DbEmail
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "batch_update_case_1",
			args: args{
				ctx: context.Background(),
				es: []*meta.DbEmail{
					dbe,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EmailUpdateAndInsertBatch(tt.args.ctx, tt.args.es); (err != nil) != tt.wantErr {
				t.Errorf("EmailUpdateBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEDelete(t *testing.T) {
	storer.DB.Exec("DELETE FROM emails WHERE id = ? ", dbe.Id)
}
