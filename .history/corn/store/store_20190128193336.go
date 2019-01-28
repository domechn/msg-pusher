/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : init.go
#   Created       : 2019/1/24 11:09
#   Last Modified : 2019/1/24 11:09
#   Describe      :
#
# ====================================================*/
package store

import (
	"time"

	"github.com/domgoer/gotask"
	"github.com/hiruok/msg-pusher/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	cornMap = make(map[string]Corn, 5)
	// ErrCornExist 已存在定时任务
	ErrCornExist = errors.New("corn: corn is existed")
)

// MustRegisterCorn 将定时任务注册到任务列表，如果重复注入返回错误
func RegisterCorn(n string, c Corn) error {
	if _, ok := cornMap[n]; ok {
		return ErrCornExist
	}
	cornMap[n] = c
	return nil
}

// MustRegisterCorn 将定时任务注册到任务列表，如果重复注入直接panic
func MustRegisterCorn(n string, c Corn) {
	if _, ok := cornMap[n]; ok {
		panic(ErrCornExist)
	}
	cornMap[n] = c
}

// Start 启动所有定时任务，此方法不会阻塞
func Start() {
	var tasks []gotask.Tasker
	for k, v := range cornMap {
		t, _ := gotask.NewTask(time.Millisecond*time.Duration(config.CornConf().Interval), StartCorn(v))
		logrus.WithFields(logrus.Fields{
			"type": "corn",
			"name": k,
		}).Info("start corn task")
		tasks = append(tasks, t)
	}
	gotask.AddToTaskList(tasks...)
}

func StartCorn(c Corn) func() {
	return func() {
		b, err := c.Read()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"type":   c.Name(),
				"method": "corn.Read",
				"error":  err.Error(),
			}).Error("读取失败")
		}
		if err = c.Write(b); err != nil {
			logrus.WithFields(logrus.Fields{
				"type":   c.Name(),
				"method": "corn.Write",
				"error":  err.Error(),
			}).Error("写入数据库失败")
		}
	}
}
