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

type Sms struct {
	RateLimit  RateLimit    `yaml:"rateLimit"`
	ServerList []*SmsServer `yaml:"serverList"`
}

type SmsServer struct {
	Server       string `yaml:"server"`
	AccessKeyId  string `yaml:"accessKeyId"`
	AccessSecret string `yaml:"accessSecret"`
	GatewayURL   string `yaml:"gatewayURL"`
}

type WeChat struct {
	RateLimit RateLimit `yaml:"rateLimit"`
	AppId     string    `yaml:"appId"`
	AppSecret string    `yaml:"appSecret"`
}

type Email struct {
	RateLimit  RateLimit      `yaml:"rateLimit"`
	ServerList []*EmailServer `yaml:"serverList"`
}

type EmailServer struct {
	Server   string `yaml:"server"`
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	TLS      bool   `yaml:"tls"`
}

type Corn struct {
	Interval int64 `yaml:"interval"`
	MaxLen   int64 `yaml:"maxLen"`
}

type Config struct {
	Mysql  *Mysql    `yaml:"mysql"`
	MQ     *RabbitMQ `yaml:"mq"`
	Redis  *Redis    `yaml:"redis"`
	Sms    *Sms      `yaml:"sms"`
	WeChat *WeChat   `yaml:"wechat"`
	Email  *Email    `yaml:"email"`
	Corn   *Corn     `yaml:"corn"`
}

type RateLimit struct {
	Every1Min  int `yaml:"every1Min"`
	Every1Hour int `yaml:"every1Hour"`
	Every1Day  int `yaml:"every1Day"`
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

func MQConf() *RabbitMQ {
	return conf.MQ
}

func RedisConf() *Redis {
	return conf.Redis
}

func SmsConf() map[string]*SmsServer {
	var res = make(map[string]*SmsServer)
	for _, v := range conf.Sms.ServerList {
		res[v.Server] = v
	}
	return res
}

func WeChatConf() *WeChat {
	return conf.WeChat
}

func EmailConf() map[string]*EmailServer {
	var res = make(map[string]*EmailServer)
	for _, v := range conf.Email.ServerList {
		res[v.Server] = v
	}
	return res
}

func CornConf() *Corn {
	return conf.Corn
}
