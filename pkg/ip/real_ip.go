// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package ip

import (
	"net"
	"net/http"
	"strings"
)

// RealIP 获取客户端的真实ip
func RealIP(r *http.Request) string {

	if contextIp := r.Context().Value("remote_addr"); contextIp != nil {
		return contextIp.(string)
	}

	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}

	if fw := r.Header.Get("X-Forwarded-For"); fw != "" {
		// X-Forwarded-For 不包含端口
		if i := strings.IndexByte(fw, ','); i >= 0 {

			return fw[:i]
		}

		return fw
	}

	// 如果上述都未找到则从remoteAddr中找IP地址，remoteAddr格式：[host:port]
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}
