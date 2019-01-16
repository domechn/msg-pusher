/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : handler.go
#   Created       : 2019/1/11 14:53
#   Last Modified : 2019/1/11 14:53
#   Describe      : 通用的组件，用来解析提交的数据，传给各个不同的handler
#
# ====================================================*/
package handler

import (
	"context"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"uuabc.com/sendmsg/api/storer/db"
	"uuabc.com/sendmsg/pkg/cache"
	"uuabc.com/sendmsg/pkg/errors"
)

var (
	// {
	//    "errcode": 0,
	//    "msg": "success",
	//    "data": null
	// }
	successResp = []byte{123, 34, 101, 114, 114, 99, 111, 100, 101, 34, 58, 48, 44, 34, 109, 115, 103, 34, 58, 34, 115, 117, 99, 99, 101, 115, 115, 34, 44, 34, 100, 97, 116, 97, 34, 58, 110, 117, 108, 108, 125}
)

// BodyHandler 需要从request.Body中取数据的func
type BodyHandler func(ctx context.Context, data []byte) ([]byte, error)

// PathHandler 从url中数据的func
type PathHandler func(ctx context.Context, data map[string]string) ([]byte, error)

// JsonHandler 针对需要从request.Body中取json的handler
func JsonHandler(sh BodyHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errors.ErrHandler(w, r, err)
			return
		}
		res, err := sh(r.Context(), b)
		if err != nil {
			err = changeErr(err)
			errors.ErrHandler(w, r, err)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(res)
	}
}

// URLHandler handler所需要的数据在url中
func URLHandler(sh PathHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := mux.Vars(r)
		res, err := sh(r.Context(), data)
		if err != nil {
			err = changeErr(err)
			errors.ErrHandler(w, r, err)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(res)
	}

}

// changeErr 转换error
func changeErr(err error) error {
	// 转换err类型
	if _, ok := err.(*errors.Error); !ok {
		if err == cache.ErrCacheMiss {
			err = errors.ErrMsgNotFound
		} else if err == db.ErrNoRowsEffected {
			err = errors.ErrMsgCantEdit
		} else {
			err = errors.NewError(
				10000000,
				err.Error(),
			)
		}
	}
	return err
}
