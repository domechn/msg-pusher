/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : corn.go
#   Created       : 2019/1/24 16:49
#   Last Modified : 2019/1/24 16:49
#   Describe      :
#
# ====================================================*/
package corn

import (
	"github.com/hiruok/msg-pusher/corn/store"
	"github.com/hiruok/msg-pusher/corn/store/db"
)

// Start 启动定时任务
func Start() {
	db.Register()

	store.Start()
}
