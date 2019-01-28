/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : produce.go
#   Created       : 2019/1/28 16:58
#   Last Modified : 2019/1/28 16:58
#   Describe      :
#
# ====================================================*/
package client

import (
	"github.com/domgoer/msg-pusher/pkg/pb/meta"
)

type Msg struct {
	value *meta.MsgProducer
}
