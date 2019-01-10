/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : redis.go
#   Created       : 2019/1/9 18:49
#   Last Modified : 2019/1/9 18:49
#   Describe      :
#
# ====================================================*/
package redis

import (
	"github.com/go-redis/redis"
	"io"
	"time"
	"uuabc.com/sendmsg/pkg/cache"
)

type Client struct {
	c redis.Cmdable
}

func NewClient(addrs []string, password string) *Client {
	if len(addrs) == 1 {
		return newClient(addrs[0], password)
	}
	return newClusterClient(addrs, password)
}

func newClient(addr string, password string) *Client {
	c := new(Client)
	c.c = redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    password,
		DialTimeout: time.Second * 5,
	})
	return c
}

func newClusterClient(addrs []string, password string) *Client {
	c := new(Client)
	c.c = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       addrs,
		Password:    password,
		DialTimeout: time.Second * 10,
	})
	return c
}

func (c *Client) Get(s string) ([]byte, error) {
	res := c.c.Get(s)
	byt, err := res.Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, cache.ErrCacheMiss
		}
		return nil, err
	}
	return byt, nil
}

func (c *Client) Put(k string, v []byte, ttl int32) error {
	res := c.c.Set(k, v, time.Second*time.Duration(ttl))
	return res.Err()
}

func (c *Client) Del(k string) error {
	return c.c.Del(k).Err()
}

func (c *Client) Add(k string, v []byte, ttl int32) error {
	b := c.c.SetNX(k, v, time.Second*time.Duration(ttl))
	if err := b.Err(); err != nil {
		return err
	}
	res, err := b.Result()
	if err != nil {
		return err
	}
	if !res {
		return cache.ErrKeyExsit
	}
	return nil
}

func (c *Client) Close() error {
	if c.c != nil {
		if v, ok := c.c.(io.Closer); ok {
			return v.Close()
		}
	}
	return nil
}
