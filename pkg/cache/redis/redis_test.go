/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : redis.go
#   Created       : 2019/1/9 18:49
#   Last Modified : 2019/1/9 18:49
#   Describe      :
#
# ====================================================*/
package redis

import (
	"context"
	"testing"

	"github.com/go-redis/redis"
)

var (
	cli, _ = NewClient([]string{"127.0.0.1:6379"}, "")
)

func TestNewClient(t *testing.T) {
	type args struct {
		addrs    []string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "new_client_case_1",
			args: args{
				addrs:    []string{"127.0.0.1:6379"},
				password: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.args.addrs, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_Put(t *testing.T) {
	type fields struct {
		c redis.Cmdable
	}
	type args struct {
		ctx context.Context
		k   string
		v   []byte
		ttl int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "put_case_1",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				k:   "hello",
				v:   []byte("test"),
				ttl: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				c: tt.fields.c,
			}
			if err := c.Put(tt.args.ctx, tt.args.k, tt.args.v, tt.args.ttl); (err != nil) != tt.wantErr {
				t.Errorf("Client.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Get(t *testing.T) {
	type fields struct {
		c redis.Cmdable
	}
	type args struct {
		ctx context.Context
		s   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "get_case_1",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				s:   "hello",
			},
		}, {
			name: "get_case_2",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				s:   "bucunzai",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				c: tt.fields.c,
			}
			_, err := c.Get(tt.args.ctx, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_Add(t *testing.T) {
	type fields struct {
		c redis.Cmdable
	}
	type args struct {
		ctx context.Context
		k   string
		v   []byte
		ttl int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "add_case_1",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				k:   "hello",
				v:   []byte("hello"),
				ttl: 10,
			},
			wantErr: true,
		}, {
			name: "add_case_2",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				k:   "hellp",
				v:   []byte("hellp"),
				ttl: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				c: tt.fields.c,
			}
			if err := c.Add(tt.args.ctx, tt.args.k, tt.args.v, tt.args.ttl); (err != nil) != tt.wantErr {
				t.Errorf("Client.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Del(t *testing.T) {
	type fields struct {
		c redis.Cmdable
	}
	type args struct {
		ctx context.Context
		k   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "del_case_1",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				k:   "hella",
			},
		}, {
			name: "del_case_2",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				k:   "hello",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				c: tt.fields.c,
			}
			if err := c.Del(tt.args.ctx, tt.args.k); (err != nil) != tt.wantErr {
				t.Errorf("Client.Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Append(t *testing.T) {
	type fields struct {
		c redis.Cmdable
	}
	type args struct {
		ctx context.Context
		k   string
		v   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "append_case_1",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				k:   "append_test",
				v:   []byte("test"),
			},
		}, {
			name: "append_case_2",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				k:   "hello",
				v:   []byte("tt"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				c: tt.fields.c,
			}
			if err := c.Append(tt.args.ctx, tt.args.k, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Client.Append() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_IsMember(t *testing.T) {
	type fields struct {
		c redis.Cmdable
	}
	type args struct {
		ctx context.Context
		k   string
		v   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ismember_case_1",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				k:   "append_test",
				v:   []byte("test"),
			},
			want: true,
		}, {
			name: "append_case_1",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				k:   "append_test1",
				v:   []byte("test"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				c: tt.fields.c,
			}
			got, err := c.IsMember(tt.args.ctx, tt.args.k, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.IsMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.IsMember() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Incr(t *testing.T) {
	type fields struct {
		c redis.Cmdable
	}
	type args struct {
		ctx context.Context
		k   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "incr_case_1",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				k:   "incr_test",
			},
			want: 1,
		}, {
			name: "incr_case_2",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				k:   "incr_test",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				c: tt.fields.c,
			}
			got, err := c.Incr(tt.args.ctx, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Incr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.Incr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Expire(t *testing.T) {
	type fields struct {
		c redis.Cmdable
	}
	type args struct {
		ctx context.Context
		k   string
		ttl int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "expire_case_1",
			fields: fields{
				c: cli.c,
			},
			args: args{
				ctx: context.Background(),
				k:   "incr_test",
				ttl: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				c: tt.fields.c,
			}
			if err := c.Expire(tt.args.ctx, tt.args.k, tt.args.ttl); (err != nil) != tt.wantErr {
				t.Errorf("Client.Expire() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Close(t *testing.T) {
	type fields struct {
		c redis.Cmdable
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
				c: tt.fields.c,
			}
			if err := c.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Client.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
