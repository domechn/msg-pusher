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
	"time"

	"github.com/go-redis/redis"
	"github.com/hiruok/msg-pusher/pkg/cache"
)

// Client 返回redis的客户端
type Client struct {
	c redis.Cmdable
}

// NewClient 新建一个redis的客户端，addr为集群地址，如果只有一个默认单机
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

// Get 按key获取value，如果不存在返回ErrCacheMiss
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

// Put 将数据放入缓存，ttl为过期时间 单位（s）
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

func (c *Client) SMembers(ctx context.Context, k string) ([]string, error) {
	return c.c.SMembers(k).Result()
}

func (c *Client) SRem(ctx context.Context, k string, v []byte) error {
	return c.c.SRem(k, v).Err()
}

func (c *Client) Incr(ctx context.Context, k string) (int64, error) {
	res := c.c.Incr(k)
	return res.Result()
}

func (c *Client) Expire(ctx context.Context, k string, ttl int64) error {
	res := c.c.Expire(k, time.Second*time.Duration(ttl))
	return res.Err()
}

func (c *Client) RPush(ctx context.Context, k string, v []byte) error {
	return c.c.RPush(k, v).Err()
}

func (c *Client) LLen(ctx context.Context, k string) (int64, error) {
	return c.c.LLen(k).Result()
}

func (c *Client) LPop(ctx context.Context, k string) ([]byte, error) {
	return c.c.LPop(k).Bytes()
}

func (c *Client) Pipeline() *Client {
	return &Client{
		c: c.c.TxPipeline(),
	}
}

func (c *Client) Discard() error {
	return c.c.(redis.Pipeliner).Discard()
}

func (c *Client) Exec() ([]interface{}, error) {
	ss, err := c.c.(redis.Pipeliner).Exec()
	if err != nil {
		return nil, err
	}
	var res []interface{}
	for _, v := range ss {
		res = append(res, v.Args()...)
	}
	return res, nil
}

func (c *Client) Close() error {
	if c.c != nil {
		if v, ok := c.c.(redis.Pipeliner); ok {
			return v.Close()
		}
	}
	return nil
}
