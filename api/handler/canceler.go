/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : canceler.go
#   Created       : 2019/1/8 16:32
#   Last Modified : 2019/1/8 16:32
#   Describe      :
#
# ====================================================*/
package handler

import (
	"context"
	"encoding/json"
	"uuabc.com/sendmsg/api/model"
	"uuabc.com/sendmsg/api/service"
)

func Cancel(ctx context.Context, body []byte) (res []byte, err error) {
	q := &model.CancelReq{}
	if err = json.Unmarshal(body, q); err != nil {
		return
	}
	if err = service.Canceler.Cancel(q.ID); err != nil {
		return
	}
	res = []byte{}
	return
}

func KeyCancel(ctx context.Context, body []byte) (res []byte, err error) {
	return
}
