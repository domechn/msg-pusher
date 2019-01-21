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
	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/receiver/handler"
	mid "uuabc.com/sendmsg/receiver/middleware"
	"uuabc.com/sendmsg/receiver/version"
)

const prometheus = "prometheus"

func Init(route *mux.Router) {
	route = route.PathPrefix("/" + version.Info.Version).Subrouter()

	// restful
	delVRoute := route.Methods("DELETE").Subrouter()
	delVRoute.Path("/sms/{id}").HandlerFunc(handler.URLHandler(handler.SmsCancel))
	delVRoute.Path("/sms/key/{key}").HandlerFunc(handler.URLHandler(handler.SmsKeyCancel))
	delVRoute.Path("/wechat/{id}").HandlerFunc(handler.URLHandler(handler.WeChatCancel))
	delVRoute.Path("/email/{id}").HandlerFunc(handler.URLHandler(handler.EmailCancel))
	handlerMiddleware(delVRoute)

	postVRoute := route.Methods("POST").Subrouter()
	postVRoute.Path("/sms").HandlerFunc(handler.JsonHandler(handler.SmsProducer))
	postVRoute.Path("/smss").HandlerFunc(handler.JsonHandler(handler.SmsProducers))
	postVRoute.Path("/wechat").HandlerFunc(handler.JsonHandler(handler.WeChatProducer))
	postVRoute.Path("/email").HandlerFunc(handler.JsonHandler(handler.EmailProducer))
	postVRoute.Path("/template").HandlerFunc(handler.JsonHandler(handler.TemplateAdd))
	handlerMiddleware(postVRoute)

	getVRoute := route.Methods("GET").Subrouter()
	getVRoute.Path("/sms/{id}").HandlerFunc(handler.URLHandler(handler.SmsIDDetail))
	getVRoute.Path("/wechat/{id}").HandlerFunc(handler.URLHandler(handler.WeChatIDDetail))
	getVRoute.Path("/email/{id}").HandlerFunc(handler.URLHandler(handler.EmailIDDetail))
	getVRoute.Path("/sms/mobile/{mobile}/page/{p}").HandlerFunc(handler.URLHandler(handler.SmsMobileDetail))
	handlerMiddleware(getVRoute)

	putVRoute := route.Methods("PATCH").Subrouter()
	putVRoute.Path("/sms").HandlerFunc(handler.JsonHandler(handler.SmsEdit))
	putVRoute.Path("/wechat").HandlerFunc(handler.JsonHandler(handler.WeChatEdit))
	putVRoute.Path("/email").HandlerFunc(handler.JsonHandler(handler.EmailEdit))
	handlerMiddleware(putVRoute)

	// postRoute := route.Methods("POST")
	//
	// versionRoute := postRoute.PathPrefix("/" + version.Info.Version).Subrouter()
	// smsRoute := versionRoute.PathPrefix("/smss").Subrouter()
	// smsRoute.Path("/producer").HandlerFunc(handler.JsonHandler(handler.SmsProducer))
	// smsRoute.Path("/producers").HandlerFunc(handler.JsonHandler(handler.SmsProducers))
	// smsRoute.Path("/detail/id").HandlerFunc(handler.JsonHandler(handler.IDDetail))
	// smsRoute.Path("/detail/mobile").HandlerFunc(handler.JsonHandler(handler.MobileDetail))
	// smsRoute.Path("/edit").HandlerFunc(handler.JsonHandler(handler.Edit))
	// smsRoute.Path("/cancel").HandlerFunc(handler.JsonHandler(handler.SmsCancel))
	// smsRoute.Path("/detail/key").HandlerFunc(handler.JsonHandler(handler.KeyDetail))
	// smsRoute.Path("/cancel_key").HandlerFunc(handler.JsonHandler(handler.SmsKeyCancel))
	//
	//
	// wechatRoute := versionRoute.PathPrefix("/wechats").Subrouter()
	// wechatRoute.Path("/producer").HandlerFunc(handler.JsonHandler(handler.WeChatProducer))
	//
	// emailRoute := versionRoute.PathPrefix("/emails").Subrouter()
	// emailRoute.Path("/producer").HandlerFunc(handler.JsonHandler(handler.EmailProducer))
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
