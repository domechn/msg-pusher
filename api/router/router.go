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

	route = route.PathPrefix("/" + version.Info.Version).Subrouter()
	// restful
	delVRoute := route.Methods("DELETE").Subrouter()
	delVRoute.Path("/sms/{id}").HandlerFunc(handler.JsonHandler(handler.SmsCancel))
	delVRoute.Path("/sms/key/{key}").HandlerFunc(handler.JsonHandler(handler.SmsKeyCancel))
	delVRoute.Path("/wechat/{id}").HandlerFunc(handler.JsonHandler(handler.WeChatCancel))
	delVRoute.Path("/email/{id}").HandlerFunc(handler.JsonHandler(handler.EmailCancel))

	postVRoute := route.Methods("POST").Subrouter()
	postVRoute.Path("/sms").HandlerFunc(handler.JsonHandler(handler.SmsProducer))
	postVRoute.Path("/smss").HandlerFunc(handler.JsonHandler(handler.SmsProducers))
	postVRoute.Path("/wechat").HandlerFunc(handler.JsonHandler(handler.WeChatProducer))
	postVRoute.Path("/email").HandlerFunc(handler.JsonHandler(handler.EmailProducer))

	getVRoute := route.Methods("GET").Subrouter()
	getVRoute.Path("/sms/{id}").HandlerFunc(handler.URLHandler(handler.SmsIDDetail))
	getVRoute.Path("/wechat/{id}").HandlerFunc(handler.URLHandler(handler.WeChatIDDetail))
	getVRoute.Path("/email/{id}").HandlerFunc(handler.URLHandler(handler.EmailIDDetail))
	getVRoute.Path("/sms/mobile/{mobile}/page/{p}").HandlerFunc(handler.URLHandler(handler.SmsMobileDetail))

	putVRoute := route.Methods("PATCH").Subrouter()
	putVRoute.Path("/sms").HandlerFunc(handler.JsonHandler(handler.SmsEdit))
	putVRoute.Path("/wechat").HandlerFunc(handler.JsonHandler(handler.WeChatEdit))
	putVRoute.Path("/email").HandlerFunc(handler.JsonHandler(handler.EmailEdit))
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
