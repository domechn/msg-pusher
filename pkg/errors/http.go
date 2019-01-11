/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : http.go
#   Created       : 2019/1/8 17:05
#   Last Modified : 2019/1/8 17:05
#   Describe      :
#
# ====================================================*/
package errors

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
)

// RecoveryHandler 程序奔溃时的返回
func ErrHandler(w http.ResponseWriter, _ *http.Request, err interface{}) {
	handler(w, err)
}

// handler 判断错误时panic还是分装的错误
func handler(w http.ResponseWriter, err interface{}) {
	switch internalErr := err.(type) {
	case *Error:
		logrus.WithFields(logrus.Fields{
			"code":  internalErr.ErrCode,
			"error": internalErr.Error(),
		}).Errorf("Internal error handled")
		toJson(w, 200, internalErr)
	case error:
		logrus.WithFields(logrus.Fields{
			"error": internalErr,
			"stack": string(debug.Stack()),
		}).Errorf("Internal server error handled")
		toJson(w, DoErr(internalErr), &Error{
			ErrCode: DoErr(internalErr),
			Msg:     internalErr.Error(),
			Data:    nil,
		})
	default:
		logrus.WithFields(logrus.Fields{
			"error": err,
			"stack": string(debug.Stack()),
		}).Errorf("Internal server error handled")
		toJson(w, http.StatusInternalServerError, err)
	}
}

// toJson 使用json格式返回
func toJson(w http.ResponseWriter, code int, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(buf.Bytes())
}
