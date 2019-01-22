/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : interface.go
#   Created       : 2019/1/10 11:33
#   Last Modified : 2019/1/10 11:33
#   Describe      :
#
# ====================================================*/
package service

type Meta interface {
	GetId() string
	// 验证参数
	Validated() error
	Marshaler
	// 转换必要的参数,请在validated调用后再使用
	Transfer(bool)
	// 获取延迟发送的时间,请在Transfer调用后使用
	Delay() int64
	// 返回参数
	GetArguments() string
	// 返回模板
	GetTemplate() string

	GetSendTo() string
	SetSendTo(string)
	GetSendTime() string
	ValidateEdit() error
}

type Marshaler interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

type Messager interface {
	Marshaler
	GetId() string
	GetContent() string
	SetContent(string)
	GetStatus() int32
	SetStatus(int32)
	SetResult(int32)
	SetCreatedAt(string)
	SetUpdatedAt(string)
	GetSendTo() string
	SetSendTo(string)
	SetArguments(string)
	GetArguments() string
	SetTemplate(string)
	GetTemplate() string
	SetSendTime(string)
	GetSendTime() string
	SetOption(int32)
	GetOption() int32
}
