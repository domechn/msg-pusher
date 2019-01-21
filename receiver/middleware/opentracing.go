/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : opentracing.go
#   Created       : 2019/1/8 16:54
#   Last Modified : 2019/1/8 16:54
#   Describe      :
#
# ====================================================*/
package middleware

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	oputil "uuabc.com/sendmsg/pkg/opentracing"
)

type statusCodeResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusCodeResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// OpenTracing 用于跟踪请求的中间件
type OpenTracing struct {
	https bool
}

func NewOpenTracing(https bool) *OpenTracing {
	return &OpenTracing{
		https: https,
	}
}

func spanStartName(r *http.Request) string {
	m := strings.ToUpper(r.Method)
	do := "Unknown"
	switch m {
	case "GET":
		do = "Detail"
	case "POST":
		do = "Produce"
	case "PATCH":
		do = "Edit"
	case "DELETE":
		do = "Cancel"
	}
	param := "Unknown"
	if strings.Contains(r.URL.Path, "email") {
		param = "Email"
	} else if strings.Contains(r.URL.Path, "sms") {
		param = "Sms"
	} else if strings.Contains(r.URL.Path, "wechat") {
		param = "WeChat"
	}
	return param + "/" + do
}

// Handler middleware接口
func (h *OpenTracing) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var span opentracing.Span
		var err error

		name := spanStartName(r)
		if parent, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header)); err != nil {
			span = opentracing.StartSpan(name)
		} else {
			span = opentracing.StartSpan(name, opentracing.ChildOf(parent))
		}
		defer span.Finish()

		host, err := os.Hostname()
		if host == "" || err != nil {
			logrus.Debugf("Failed to get host name, error: %+v", err)
			host = "unknown"
		}

		scheme := "http"
		if h.https {
			scheme = "https"
		}
		ext.HTTPMethod.Set(span, r.Method)
		ext.HTTPUrl.Set(span, fmt.Sprintf(scheme+"://"+r.Host+r.URL.Path))
		ext.Component.Set(span, "sendmsg")
		ext.SpanKind.Set(span, "server")
		if hostname, portString, err := net.SplitHostPort(r.URL.Host); err != nil {
			ext.PeerHostname.Set(span, hostname)
			if port, err := strconv.Atoi(portString); err != nil {
				ext.PeerPort.Set(span, uint16(port))
			}
		} else {
			ext.PeerHostname.Set(span, r.URL.Host)
		}

		span.SetTag("user.agent", r.UserAgent())
		span.SetTag("peer.address", r.RemoteAddr)
		span.SetTag("host.name", host)
		span.SetTag("referer", r.Referer())
		span.SetTag("request.id", RequestID(r.Context()))

		err = span.Tracer().Inject(
			span.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			logrus.Errorf("Could not inject span context into header, error: %+v", err)
		}

		w = &statusCodeResponseWriter{w, 0}

		next.ServeHTTP(w, oputil.ToContext(r, span))

		if scr, ok := w.(*statusCodeResponseWriter); ok {
			code := uint16(scr.status)
			ext.HTTPStatusCode.Set(span, code)
			if code >= http.StatusInternalServerError {
				ext.Error.Set(span, true)
			}
		}

	})
}
