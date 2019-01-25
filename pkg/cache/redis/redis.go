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

	"github.com/domgoer/msg-pusher/pkg/cache"
	"github.com/go-redis/redis"
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

func (c *Client) ZAdd(ctx context.Context, k string, score int, v []byte) error {
	return c.c.ZAdd(k, redis.Z{
		Score:  float64(score),
		Member: v,
	}).Err()
}

func (c *Client) ZRange(ctx context.Context, k string, start, end int64) ([]string, error) {
	return c.c.ZRange(k, start, end).Result()
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
