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

import (
	"context"
	"github.com/hiruok/msg-pusher/pkg/pb/meta"
)

type Messager interface {
	GetId() string
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	GetStatus() meta.Status
	SetStatus(meta.Status)
	SetTryNum(int32)
	SetResult(meta.Result)
	SetReason(string)
	GetSendTime() string
	SetUpdatedAt(string)
	GetVersion() int32
	SetVersion(int32)
}

type Cache interface {
	RPush(context.Context, []byte) error
}
