/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : sms.go
#   Created       : 2019/1/11 17:02
#   Last Modified : 2019/1/11 17:02
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"uuabc.com/sendmsg/api/storer"
)

// CancelSmsMsgByID 将sms信息的发送状态设置为取消
func CancelSmsMsgByID(ctx context.Context, id string) error {
	stmt, err := storer.DB.PrepareContext(ctx, "UPDATE smss SET status=2 WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, id)
	return err
}
