/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : router.go
#   Created       : 2019/1/8 15:19
#   Last Modified : 2019/1/8 15:19
#   Describe      :
#
# ====================================================*/
package router

import (
	"net/http"

	"github.com/hiruok/msg-pusher/pkg/errors"
	"github.com/hiruok/msg-pusher/receiver/handler"
	mid "github.com/hiruok/msg-pusher/receiver/middleware"
	"github.com/hiruok/msg-pusher/receiver/version"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const prometheus = "prometheus"

var (
	routes = []struct {
		name    string
		method  string
		path    string
		handler http.Handler
	}{
		// delete
		{
			name:    "CancelByID",
			method:  "DELETE",
			path:    "/msg/{id}",
			handler: handler.URLHandler(handler.MsgCancel),
		}, {
			name:    "CancelByKey",
			method:  "DELETE",
			path:    "/msg/plat/{plat}/key/{key}",
			handler: handler.URLHandler(handler.MsgCancelByKey),
		},
		// post
		{
			name:    "Produce",
			method:  "POST",
			path:    "/msg",
			handler: handler.JsonHandler(handler.MsgProducer),
		}, /* {
			name:    "ProduceBatch",
			method:  "POST",
			path:    "/msgs",
			handler: handler.JsonHandler(handler.MsgProducer),
		},*/{
			name:    "AddTemplate",
			method:  "POST",
			path:    "/template",
			handler: handler.JsonHandler(handler.TemplateAdd),
		},
		// get
		{
			name:    "DetailByID",
			method:  "GET",
			path:    "/msg/{id}",
			handler: handler.URLHandler(handler.MsgIDDetail),
		}, {
			name:    "DetailByPlat",
			method:  "GET",
			path:    "/msg/key/{key}/p/{page}",
			handler: handler.URLHandler(handler.MsgDetailByKey),
		}, {
			name:    "DetailByToAndPage",
			method:  "GET",
			path:    "/msg/to/{to}/page/{p}",
			handler: handler.URLHandler(handler.MsgDetailByTo),
		},
		// Patch
		{
			name:    "Edit",
			method:  "PATCH",
			path:    "/msg",
			handler: handler.JsonHandler(handler.MsgEdit),
		},
	}
)

func Init(route *mux.Router) {

	handlerMiddleware(route)

	for _, v := range routes {
		logrus.WithFields(logrus.Fields{
			"type":   "route",
			"method": v.method,
			"name":   v.name,
			"path":   "/" + version.Info.Version + v.path,
		}).Info()
		route.StrictSlash(true).
			Methods(v.method).
			Path("/" + version.Info.Version + v.path).
			Name(v.name).
			Handler(v.handler)
	}
}

func handlerMiddleware(r *mux.Router) {
	r.Use(
		mid.RequestIDMiddleware,
		mid.LoggingMiddleware,
		mid.RecoveryMiddleware(errors.ErrHandler),
		mid.NewOpenTracing(false).Handler,
		mid.NewMetrics(prometheus).Handler,
	)
}
