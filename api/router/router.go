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
	"github.com/gorilla/mux"
	"uuabc.com/sendmsg/api/handler"
	mid "uuabc.com/sendmsg/api/middleware"
	"uuabc.com/sendmsg/api/version"
	"uuabc.com/sendmsg/pkg/errors"
)

const prometheus = "prometheus"

func Init(route *mux.Router) {
	route.Use(
		mid.RequestIDMiddleware,
		mid.LoggingMiddleware,
		mid.RecoveryMiddleware(errors.ErrHandler),
		mid.NewOpenTracing(false).Handler,
		mid.NewMetrics(prometheus).Handler,
	)

	postRoute := route.Methods("POST")

	versionRoute := postRoute.PathPrefix("/" + version.Info.Version).Subrouter()
	msgRoute := versionRoute.PathPrefix("/messages").Subrouter()
	msgRoute.Path("/producer").HandlerFunc(handler.JsonHandler(handler.SmsProducer))
	msgRoute.Path("/producers").HandlerFunc(handler.JsonHandler(handler.SmsProducers))
	msgRoute.Path("/detail/id").HandlerFunc(handler.JsonHandler(handler.IDDetail))
	msgRoute.Path("/detail/mobile").HandlerFunc(handler.JsonHandler(handler.MobileDetail))
	msgRoute.Path("/edit").HandlerFunc(handler.JsonHandler(handler.Edit))
	msgRoute.Path("/cancel").HandlerFunc(handler.JsonHandler(handler.Cancel))
	msgRoute.Path("/detail/key").HandlerFunc(handler.JsonHandler(handler.KeyDetail))
	msgRoute.Path("/cancel_key").HandlerFunc(handler.JsonHandler(handler.KeyCancel))

	versionRoute.Path("/wechats/producer").HandlerFunc(handler.JsonHandler(handler.WeChatProducer))
	versionRoute.Path("/emails/producer").HandlerFunc(handler.JsonHandler(handler.EmailProducer))
}
