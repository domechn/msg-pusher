/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : init.go
#   Created       : 2019/1/9 16:58
#   Last Modified : 2019/1/9 16:58
#   Describe      :
#
# ====================================================*/
package receiver

import (
	"github.com/hiruok/msg-pusher/receiver/monitor"
	"github.com/hiruok/msg-pusher/receiver/router"
	"github.com/hiruok/msg-pusher/storer"
	"github.com/gorilla/mux"
)

func Init(route *mux.Router, addrMonitor string) error {
	if err := storer.Init(); err != nil {
		return err
	}
	router.Init(route)
	if err := monitor.StartMetrics(addrMonitor); err != nil {
		return err
	}
	return nil
}

func Close() error {
	return storer.Close()
}
