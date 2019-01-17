/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : send.go
#   Created       : 2019/1/7 14:33
#   Last Modified : 2019/1/7 14:33
#   Describe      :
#
# ====================================================*/
package sms

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/json-iterator/go"
	"uuabc.com/sendmsg/pkg/send"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const defaultTimeout = time.Second * 10

// Client 发送短信客户端，implements send.Sender
type Client struct {
	cfg    Config
	client *http.Client
}

func NewClient(cfg Config) *Client {
	return &Client{
		cfg: cfg,
		client: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

// Send 发送短信，msg需要使用sms.NewRequest(...)生成
// 可以使用do将请求结果回传,也可以传nil忽略结果
func (c *Client) Send(msg send.Message, do send.DoRes) error {
	if c.client == nil {
		return fmt.Errorf("sender initializes the exception, use the NewSender() method to initialize it")
	}
	var req *Request
	var ok bool
	if req, ok = msg.(*Request); !ok {
		return fmt.Errorf("this type is not supported, use sms.NewRequest()")
	}
	req.AccessKeyId = c.cfg.AccessKeyId
	sign, err := req.Encode(c.cfg.AccessSecret, c.cfg.GatewayURL)
	if err != nil {
		return nil
	}
	request, _ := http.NewRequest("GET", sign, nil)
	resp, err := c.client.Do(request)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	result := &Response{}
	if err := json.Unmarshal(data, result); err != nil {
		return err
	}
	result.RawResponse = data
	if do != nil {
		do(result)
	}
	if !result.IsSuccessful() {
		return fmt.Errorf("send msg failed,\ncode: %s\nmsg: %s", result.Code, result.Message)
	}
	return nil
}
