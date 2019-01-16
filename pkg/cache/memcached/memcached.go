/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : mem.go
#   Created       : 2019/1/8 10:36
#   Last Modified : 2019/1/8 10:36
#   Describe      : 弃用，无法满足要求
#
# ====================================================*/
package memcached

import (
	"context"

	"github.com/bradfitz/gomemcache/memcache"
	"uuabc.com/sendmsg/pkg/cache"
)

type Client struct {
	c *memcache.Client
}

// NewClient 返回一个新的memcachedclient，client implements Cached
func NewClient(server ...string) *Client {
	cli := memcache.New(server...)
	client := &Client{
		c: cli,
	}
	return client
}

func (c *Client) Get(ctx context.Context, s string) ([]byte, error) {
	item, err := c.c.Get(s)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return nil, cache.ErrCacheMiss
		}
		return nil, err
	}
	return item.Value, nil
}

func (c *Client) Put(ctx context.Context, k string, v []byte, ttl int32) error {
	item := &memcache.Item{
		Key:   k,
		Value: v,
	}
	if ttl > 0 {
		item.Expiration = ttl
	}
	return c.c.Set(item)
}

func (c *Client) Del(ctx context.Context, k string) error {
	return c.c.Delete(k)
}

func (c *Client) Add(ctx context.Context, k string, v []byte, ttl int32) error {
	item := &memcache.Item{
		Key:   k,
		Value: v,
	}
	if ttl > 0 {
		item.Expiration = ttl
	}
	err := c.c.Add(item)
	if err != nil {
		if err == memcache.ErrNotStored {
			return cache.ErrKeyExsit
		}
	}
	return err
}

func (c *Client) Append(ctx context.Context, k string, v []byte) error {
	return nil
}

func (c *Client) IsMember(ctx context.Context, k string, v []byte) (bool, error) {
	return false, nil
}

func (c *Client) Close() error {
	return nil
}
