/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : producer.go
#   Created       : 2019/1/8 16:32
#   Last Modified : 2019/1/8 16:32
#   Describe      :
#
# ====================================================*/
package handler

import (
	"context"
	"uuabc.com/sendmsg/api/service"
)

func processData(ctx context.Context, p service.Meta) (err error) {
	if err = p.Validated(); err != nil {
		return
	}
	p.Transfer()
	err = service.ProducerImpl.Produce(ctx, p)
	return
}
