/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : error.go
#   Created       : 2019/1/8 17:07
#   Last Modified : 2019/1/8 17:07
#   Describe      :
#
# ====================================================*/
package errors

import (
	"net/http"
	"strconv"
	"strings"
)

// Error 现错误时的返回值
type Error struct {
	ErrCode int         `json:"errcode"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// New creates a new instance of Error
func NewError(code int, message string) *Error {
	return &Error{code, message, nil}
}

// Error 返回错误的信息
func (e *Error) Error() string {
	return e.Msg
}

// GetCode 返回错误码
func (e *Error) GetCode() int {
	return e.ErrCode
}

// Marshal 将错误编码
func (e *Error) Marshal() []byte {
	str := `{"code":` + strconv.FormatInt(int64(e.ErrCode), 10) + `,"msg":"` + e.Error() + `","data":{}}`
	return []byte(str)
}

// DoErr
func DoErr(err error) int {
	str := err.Error()
	if strings.Contains(str, "EOF") {
		return http.StatusInternalServerError
	} else if strings.Contains(str, "refused") {
		return http.StatusServiceUnavailable
	} else if strings.Contains(str, "canceled") {
		return http.StatusRequestTimeout
	} else if strings.Contains(str, "not found") {
		return http.StatusNotFound
	}
	return http.StatusBadRequest
}
