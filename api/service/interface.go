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
	// 返回消息类型，sms还是wechat还是email
	TypeName() string
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	// 转换必要的参数,请在validated调用后再使用
	Transfer()
	// 获取延迟发送的时间,请在Transfer调用后使用
	Delay() int64
}