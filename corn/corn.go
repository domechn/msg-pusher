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
	"uuabc.com/sendmsg/corn/store"
	"uuabc.com/sendmsg/corn/store/db"
)

func Start() {
	db.Register()

	store.Start()
}
