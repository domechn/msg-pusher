/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : wechat.go
#   Created       : 2019/1/11 16:58
#   Last Modified : 2019/1/11 16:58
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"uuabc.com/sendmsg/api/storer"
)

// CancelWeChatMsgByID 将wechat信息的发送状态设置为取消
func CancelWeChatMsgByID(ctx context.Context, id string) error {
	stmt, err := storer.DB.PrepareContext(ctx, "UPDATE wechats SET status=2 WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, id)
	return err
}
