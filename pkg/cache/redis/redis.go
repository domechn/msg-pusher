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
	"context"
	"github.com/go-redis/redis"
	"io"
	"time"
	"uuabc.com/sendmsg/pkg/cache"
)

type Client struct {
	c redis.Cmdable
}

func NewClient(addrs []string, password string) (*Client, error) {
	if len(addrs) == 1 {
		return newClient(addrs[0], password)
	}
	return newClusterClient(addrs, password)
}

func newClient(addr string, password string) (*Client, error) {
	c := new(Client)
	c.c = redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    password,
		DialTimeout: time.Second * 5,
	})
	err := c.c.Ping().Err()
	return c, err
}

func newClusterClient(addrs []string, password string) (*Client, error) {
	c := new(Client)
	c.c = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       addrs,
		Password:    password,
		DialTimeout: time.Second * 10,
	})
	err := c.c.Ping().Err()
	return c, err
}

func (c *Client) Get(ctx context.Context, s string) ([]byte, error) {
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

func (c *Client) Put(ctx context.Context, k string, v []byte, ttl int64) error {
	res := c.c.Set(k, v, time.Second*time.Duration(ttl))
	return res.Err()
}

func (c *Client) Del(ctx context.Context, k string) error {
	return c.c.Del(k).Err()
}

func (c *Client) Add(ctx context.Context, k string, v []byte, ttl int64) error {
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

func (c *Client) Append(ctx context.Context, k string, v []byte) error {
	res := c.c.SAdd(k, v)
	return res.Err()
}

func (c *Client) IsMember(ctx context.Context, k string, v []byte) (bool, error) {
	res := c.c.SIsMember(k, v)
	return res.Result()
}

// PutBit 修改k中offset的值，val in (0,1)
func (c *Client) PutBit(ctx context.Context, k string, offset int64, val int) error {
	return c.c.SetBit(k, offset, val).Err()
}

// GetBit 从k中获取offset的值（0或1）
func (c *Client) GetBit(ctx context.Context, k string, offset int64) (int, error) {
	res := c.c.GetBit(k, offset)
	r, err := res.Result()
	return int(r), err
}

func (c *Client) Close() error {
	if c.c != nil {
		if v, ok := c.c.(io.Closer); ok {
			return v.Close()
		}
	}
	return nil
}
