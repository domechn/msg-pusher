/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : client.go
#   Created       : 2019/1/28 16:54
#   Last Modified : 2019/1/28 16:54
#   Describe      :
#
# ====================================================*/
package client

import (
	"net"
)

type Client struct {
	Addr string
}

// NewClient 返回一个客户端
func NewClient(addr string) *Client {
	return &Client{
		Addr: addr,
	}
}

// Ping 查看是否能够连接到指定的服务
func (c *Client) Ping() error {
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
