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
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// Init 初始化日志的输出等级，输出路径
func Init(typeN, path, level string) {
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
	if path != "" {
		// 日志分割 "/path/typeN.log.20060102"
		rl, err := rotatelogs.New(transferPath(path)+typeN+".log.%Y%m%d", rotatelogs.WithClock(rotatelogs.UTC))
		rotate(rl)
		if err == nil {
			logrus.SetOutput(rl)
		} else {
			logrus.Warn("logrus init:cannot open log file,log will only be printed on std")
		}
	}
}

func transferPath(s string) string {
	if s == "" {
		return "/"
	}
	if s[len(s)-1] != '/' {
		return s + "/"
	}
	return s
}

// 当进程启动时如果遇到log文件名冲突，则为表单的数字后缀“。1”、“。2“,”。3”等等被附加到日志文件的末尾
func rotate(rl *rotatelogs.RotateLogs) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP)
	go func(ch chan os.Signal) {
		<-ch
		rl.Rotate()
	}(ch)
}
