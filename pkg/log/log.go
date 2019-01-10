/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : log.go
#   Created       : 2019/1/8 15:12
#   Last Modified : 2019/1/8 15:12
#   Describe      :
#
# ====================================================*/
package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func Init(path, level string) {
	l := logrus.InfoLevel
	switch strings.ToUpper(level) {
	case "DEBUG":
		l = logrus.DebugLevel
	case "WARN":
		l = logrus.WarnLevel
	case "ERROR":
		l = logrus.ErrorLevel
	}
	logrus.SetLevel(l)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(f)
	} else {
		logrus.Warn("logrus init:cannot open log file,log will only be printed on std")
	}
}
