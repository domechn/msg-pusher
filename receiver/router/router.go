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

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/receiver/handler"
	mid "uuabc.com/sendmsg/receiver/middleware"
	"uuabc.com/sendmsg/receiver/version"
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
			name:    "CancelSmsByID",
			method:  "DELETE",
			path:    "/sms/{id}",
			handler: handler.URLHandler(handler.SmsCancel),
		}, {
			name:    "CancelSmsByKey",
			method:  "DELETE",
			path:    "/sms/key/{key}",
			handler: handler.URLHandler(handler.SmsKeyCancel),
		}, {
			name:    "CancelWeChatByID",
			method:  "DELETE",
			path:    "/wechat/{id}",
			handler: handler.URLHandler(handler.WeChatCancel),
		}, {
			name:    "CancelEmailByID",
			method:  "DELETE",
			path:    "/email/{id}",
			handler: handler.URLHandler(handler.EmailCancel),
		},
		// post
		{
			name:    "ProduceSms",
			method:  "POST",
			path:    "/sms",
			handler: handler.JsonHandler(handler.SmsProducer),
		}, {
			name:    "ProduceSmss",
			method:  "POST",
			path:    "/smss",
			handler: handler.JsonHandler(handler.SmsProducers),
		}, {
			name:    "ProduceWeChat",
			method:  "POST",
			path:    "/wechat",
			handler: handler.JsonHandler(handler.WeChatProducer),
		}, {
			name:    "ProduceEmail",
			method:  "POST",
			path:    "/email",
			handler: handler.JsonHandler(handler.EmailProducer),
		}, {
			name:    "AddTemplate",
			method:  "POST",
			path:    "/template",
			handler: handler.JsonHandler(handler.TemplateAdd),
		},
		// get
		{
			name:    "SmsDetailByID",
			method:  "GET",
			path:    "/sms/{id}",
			handler: handler.URLHandler(handler.SmsIDDetail),
		}, {
			name:    "WeChatDetailByID",
			method:  "GET",
			path:    "/wechat/{id}",
			handler: handler.URLHandler(handler.WeChatIDDetail),
		}, {
			name:    "EmailDetailById",
			method:  "GET",
			path:    "/email/{id}",
			handler: handler.URLHandler(handler.EmailIDDetail),
		}, {
			name:    "SmsDetailByMobileAndPage",
			method:  "GET",
			path:    "/sms/mobile/{mobile}/page/{p}",
			handler: handler.URLHandler(handler.SmsMobileDetail),
		},
		// Patch
		{
			name:    "SmsEdit",
			method:  "PATCH",
			path:    "/sms",
			handler: handler.JsonHandler(handler.SmsEdit),
		}, {
			name:    "WeChatEdit",
			method:  "PATCH",
			path:    "/wechat",
			handler: handler.JsonHandler(handler.WeChatEdit),
		}, {
			name:    "EmailEdit",
			method:  "PATCH",
			path:    "/email",
			handler: handler.JsonHandler(handler.EmailEdit),
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
