/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : interface.go
#   Created       : 2019/1/23 15:06
#   Last Modified : 2019/1/23 15:06
#   Describe      :
#
# ====================================================*/
package sender

import (
	"context"
)

type Cache interface {
	RPushEmail(context.Context, []byte) error
	RPushWeChat(context.Context, []byte) error
	RPushSms(context.Context, []byte) error
}
