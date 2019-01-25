/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : sms.go
#   Created       : 2019/1/11 17:02
#   Last Modified : 2019/1/11 17:02
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"testing"

	"github.com/domgoer/msg-pusher/pkg/db"
	"github.com/domgoer/msg-pusher/pkg/pb/meta"
	"github.com/domgoer/msg-pusher/storer"
	"github.com/jmoiron/sqlx"
)

func init() {
	storer.DB, _ = db.New(db.Config{
		URL: "root:root@tcp(localhost:3306)/uuabc?charset=utf8&parseTime=True",
	})
}

var dbt = &meta.DbSms{
	Id:          "test-test-test-test-insert",
	Platform:    1,
	PlatformKey: "test-platform",
	Content:     "test-content",
	Mobile:      "14323232321",
	Type:        1,
	Template:    "test-template",
	Arguments:   "test-args",
	Server:      1,
	SendTime:    "2018-09-09 10:10:10",
}

func TestSmsInsert(t *testing.T) {

	type args struct {
		ctx context.Context
		sms *meta.DbSms
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
				sms: dbt,
			},
			wantErr: false,
		}, {
			name: "insert_case_2",
			args: args{
				ctx: context.Background(),
				sms: dbt,
			},
			wantErr: true,
		}, {
			name: "insert_case_3",
			args: args{
				ctx: context.Background(),
				sms: &meta.DbSms{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := SmsInsert(tt.args.ctx, tt.args.sms)
			if err != nil {
				RollBack(tx)
			} else {
				Commit(tx)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("SmsInsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSmsDetailByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *meta.DbSms
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "detailID_case_1",
			args: args{
				ctx: context.Background(),
				id:  "test-test-test-test-insert",
			},
		}, {
			name: "detailID_case_2",
			args: args{
				ctx: context.Background(),
				id:  "test-test-insert",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SmsDetailByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SmsDetailByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSmsDetailByPlat(t *testing.T) {
	type args struct {
		ctx         context.Context
		platform    int32
		platformKey string
	}
	tests := []struct {
		name    string
		args    args
		want    []*meta.DbSms
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "detailP_case_1",
			args: args{
				ctx:         context.Background(),
				platform:    1,
				platformKey: "test-platform",
			},
		}, {
			name: "detailP_case_2",
			args: args{
				ctx:         context.Background(),
				platform:    0,
				platformKey: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SmsDetailByPlat(tt.args.ctx, tt.args.platform, tt.args.platformKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("SmsDetailByPlat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSmsDetailByPhoneAndPage(t *testing.T) {
	type args struct {
		ctx    context.Context
		mobile string
		page   int
	}
	tests := []struct {
		name    string
		args    args
		want    []*meta.DbSms
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "detailM_case_1",
			args: args{
				ctx:    context.Background(),
				mobile: "14323232321",
				page:   1,
			},
		}, {
			name: "detailM_case-2",
			args: args{
				ctx:    context.Background(),
				mobile: "23",
				page:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SmsDetailByPhoneAndPage(tt.args.ctx, tt.args.mobile, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("SmsDetailByPhoneAndPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSmsEdit(t *testing.T) {
	dbt.Id = "test-test-test-test-insert"
	dbt.Mobile = "123"
	type args struct {
		ctx context.Context
		s   *meta.DbSms
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
				s:   dbt,
			},
		}, {
			name: "edit_case_2",
			args: args{
				ctx: context.Background(),
				s: &meta.DbSms{
					Id: "test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := SmsEdit(tt.args.ctx, tt.args.s)
			if err != nil {
				RollBack(tx)
			} else {
				Commit(tx)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("SmsEdit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSmsUpdateSendResult(t *testing.T) {
	dbt.TryNum = 3
	type args struct {
		ctx context.Context
		s   *meta.DbSms
	}
	tests := []struct {
		name    string
		args    args
		want    *sqlx.Tx
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "update_res_case_1",
			args: args{
				ctx: context.Background(),
				s:   dbt,
			},
		}, {
			name: "update_res_case_2",
			args: args{
				ctx: context.Background(),
				s: &meta.DbSms{
					Id: "test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := SmsUpdateSendResult(tt.args.ctx, tt.args.s)
			if err != nil {
				RollBack(tx)
			} else {
				Commit(tx)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("SmsUpdateSendResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSmsCancelByID(t *testing.T) {
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
				id:  "test-test-test-test-insert",
			},
			wantErr: false,
		}, {
			name: "cancel_case_2",
			args: args{
				ctx: context.Background(),
				id:  "test-test-test-test-insert",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SmsCancelByID(tt.args.ctx, tt.args.id)
			if err != nil {
				RollBack(got)
			} else {
				Commit(got)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("SmsCancelByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSmsCancelByPlat(t *testing.T) {
	type args struct {
		ctx         context.Context
		platform    int32
		platformKey string
	}
	tests := []struct {
		name    string
		args    args
		want    *sqlx.Tx
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "detailP_case_1",
			args: args{
				ctx:         context.Background(),
				platform:    1,
				platformKey: "test-platformKey",
			},
			wantErr: true,
		}, {
			name: "detailP_case_2",
			args: args{
				ctx:         context.Background(),
				platformKey: "t",
				platform:    0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SmsCancelByPlat(tt.args.ctx, tt.args.platform, tt.args.platformKey)
			if err != nil {
				RollBack(got)
			} else {
				Commit(got)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("SmsCancelByPlat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSmsUpdateAndInsertBatch(t *testing.T) {
	type args struct {
		ctx context.Context
		ds  []*meta.DbSms
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "batch_insert_case1",
			args: args{
				ctx: context.Background(),
				ds: []*meta.DbSms{
					dbt,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbt.Type = 6
			if err := SmsUpdateAndInsertBatch(tt.args.ctx, tt.args.ds); (err != nil) != tt.wantErr {
				t.Errorf("SmsUpdateAndInsertBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	storer.DB.Exec("DELETE FROM smss WHERE id = ?", dbt.Id)
}
