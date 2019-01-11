/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : json.go
#   Created       : 2019/1/10 11:15
#   Last Modified : 2019/1/10 11:15
#   Describe      :
#
# ====================================================*/
package handler

import (
	"context"
	"io/ioutil"
	"net/http"
	"uuabc.com/sendmsg/pkg/errors"
)

type ServiceHandler func(ctx context.Context, data []byte) ([]byte, error)

func JsonHandler(sh ServiceHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errors.ErrHandler(w, r, err)
			return
		}
		data, err := sh(r.Context(), b)
		if err != nil {
			errors.ErrHandler(w, r, err)
			return
		}
		w.Write(data)
	}
}
