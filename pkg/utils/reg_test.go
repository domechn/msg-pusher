/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : reg.go
#   Created       : 2019/1/10 16:37
#   Last Modified : 2019/1/10 16:37
#   Describe      :
#
# ====================================================*/
package utils

import (
	"reflect"
	"testing"
)

func TestValidatePhone(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				s: "13143234543",
			},
			want: true,
		}, {
			name: "case2",
			args: args{
				s: "123",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePhone(tt.args.s); got != tt.want {
				t.Errorf("ValidatePhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateEmailAddr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				s: "safd",
			},
		}, {
			name: "case2",
			args: args{
				s: "abc@a.b",
			},
			want: true,
		}, {
			name: "case3",
			args: args{
				s: "abc.cmb_@ac.com",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateEmailAddr(tt.args.s); got != tt.want {
				t.Errorf("ValidateEmailAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrFromCurlyBraces(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				s: "hello,${abc}",
			},
			want: []string{"${abc}"},
		}, {
			name: "case2",
			args: args{
				s: "hello,${abc}${bcd}",
			},
			want: []string{"${abc}", "${bcd}"},
		}, {
			name: "case3",
			args: args{
				s: "asd",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrFromCurlyBraces(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrFromCurlyBraces() = %v, want %v", got, tt.want)
			}
		})
	}
}
