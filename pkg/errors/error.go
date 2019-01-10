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

// Error gateway代理出现错误时的返回值
type Error struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// New creates a new instance of Error
func NewError(code int, message string) *Error {
	return &Error{code, message, nil}
}

func (e *Error) Error() string {
	return e.Msg
}

func (e *Error) GetCode() int {
	return e.Code
}

func (e *Error) Marshal() []byte {
	str := `{"code":` + strconv.FormatInt(int64(e.Code), 10) + `,"msg":"` + e.Error() + `","data":{}}`
	return []byte(str)
}

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
