/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : send.go
#   Created       : 2019/1/7 14:34
#   Last Modified : 2019/1/7 14:34
#   Describe      :
#
# ====================================================*/
package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"uuabc.com/sendmsg/pkg/cache/redis"

	"uuabc.com/sendmsg/pkg/cache"
	"uuabc.com/sendmsg/pkg/cst"
	"uuabc.com/sendmsg/send"
)

const (
	defaultTimeout = time.Second * 10
	tokenURL       = "https://api.weixin.qq.com/cgi-bin/token"
	sendURL        = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="
)

var (
	ErrTokenOverdue = errors.New("weixin_token is overdued")
)

type Client struct {
	httpCli *http.Client
	cached  cache.Cache
	cfg     *Config
}

func NewClient(cfg *Config) *Client {
	return &Client{
		cfg:    cfg,
		cached: redis.NewClient(cfg.CacheAddrs, cfg.CachePwd),
		httpCli: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

// accessTokenData 先从缓存中获取token，如果不存在再http请求获取
func (c *Client) accessTokenData() (string, error) {
	token, err := c.token()
	if err == nil {
		return token, nil
	}
	result, err := c.requestAccessToken()
	if err != nil {
		return "", err
	}
	go func() {
		c.storeToken(result.AccessToken, result.ExpiresIn)
	}()
	return result.AccessToken, nil
}

// requestAccessToken 通过http请求获取weixin_token
func (c *Client) requestAccessToken() (*Response, error) {
	url := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", tokenURL, c.cfg.APPId, c.cfg.APPSecret)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := &Response{}
	err = json.Unmarshal(b, result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, fmt.Errorf("get accessTokenData failed,code: %d error: %v\n", result.ErrCode, result.ErrMsg)
	}
	return result, nil
}

// token 从缓存中获取token
func (c *Client) token() (string, error) {
	b, err := c.cached.Get(cst.WeiXinAccessToken)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// storeToken 在缓存中存储token
func (c *Client) storeToken(v string, e int) error {
	return c.cached.Put(cst.WeiXinAccessToken, []byte(v), int32(e/60))
}

func (c *Client) Send(msg send.Message, do send.DoRes) error {
	token, err := c.accessTokenData()
	if err != nil {
		return err
	}
	url := sendURL + token
	// 微信token失效返回json {"errcode":40001,"errmsg":"invalid credential, access_token is invalid or not latest hint: [CVJXrA0584vr61!]"}
	r := bytes.NewReader(msg.Content())
	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		return err
	}
	resp, err := c.httpCli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	result := &Response{}
	err = json.Unmarshal(b, result)
	if err != nil {
		return err
	}
	if result.ErrCode == 40001 {
		return ErrTokenOverdue
	}
	if result.ErrCode != 0 {
		return fmt.Errorf("send msg failed,code: %d error: %v\n", result.ErrCode, result.ErrMsg)
	}
	if do != nil {
		do(result)
	}
	return nil
}
