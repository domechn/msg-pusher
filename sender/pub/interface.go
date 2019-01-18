/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : interface.go
#   Created       : 2019/1/16 16:54
#   Last Modified : 2019/1/16 16:54
#   Describe      :
#
# ====================================================*/
package pub

type Messager interface {
	GetId() string
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	GetStatus() int32
	SetStatus(int32)
	SetTryNum(int32)
	SetResult(int32)
	SetReason(string)
	GetSendTime() string
}
