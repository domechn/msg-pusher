/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : email.go
#   Created       : 2019/1/11 17:03
#   Last Modified : 2019/1/11 17:03
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"uuabc.com/sendmsg/api/storer"
)

// CancelEmailMsgByID 将email信息的发送状态设置为取消
func CancelEmailMsgByID(ctx context.Context, id string) error {
	stmt, err := storer.DB.PrepareContext(ctx, "UPDATE emails SET status=2 WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, id)
	return err
}
