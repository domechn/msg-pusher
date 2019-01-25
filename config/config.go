/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : config.go
#   Created       : 2019/1/8 19:37
#   Last Modified : 2019/1/8 19:37
#   Describe      :
#
# ====================================================*/
package config

import (
	"github.com/hiruok/gconf"
)

type Mysql struct {
	URL             string `yaml:"url"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
}

type Redis struct {
	Addrs    []string `yaml:"addrs"`
	Password string   `yaml:"password"`
}

type Memcached struct {
	Addr string `yaml:"addr"`
}

type RabbitMQ struct {
	URL          string `yaml:"url"`
	ExChangeName string `yaml:"exChangeName"`
}

type Aliyun struct {
	AccessKeyId  string `yaml:"accessKeyId"`
	AccessSecret string `yaml:"accessSecret"`
	GatewayURL   string `yaml:"gatewayURL"`
}

type WeChat struct {
	AppId     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

type Email struct {
	ServerAddr string `yaml:"serverAddr"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	TLS        bool   `yaml:"tls"`
}

type Corn struct {
	Interval int64 `yaml:"interval"`
	MaxLen   int64 `yaml:"maxLen"`
}

type Config struct {
	Mysql     *Mysql     `yaml:"mysql"`
	Memcached *Memcached `yaml:"memcached"`
	MQ        *RabbitMQ  `yaml:"mq"`
	Redis     *Redis     `yaml:"redis"`
	Aliyun    *Aliyun    `yaml:"aliyun"`
	WeChat    *WeChat    `yaml:"wechat"`
	Email     *Email     `yaml:"email"`
	Corn      *Corn      `yaml:"corn"`
}

var (
	conf = &Config{}
)

func Init(path string) error {
	return gconf.Read2Struct(path, conf)
}

func MysqlConf() *Mysql {
	return conf.Mysql
}

func MemCachedConf() *Memcached {
	return conf.Memcached
}

func MQConf() *RabbitMQ {
	return conf.MQ
}

func RedisConf() *Redis {
	return conf.Redis
}

func AliyunConf() *Aliyun {
	return conf.Aliyun
}

func WeChatConf() *WeChat {
	return conf.WeChat
}

func EmailConf() *Email {
	return conf.Email
}

func CornConf() *Corn {
	return conf.Corn
}
