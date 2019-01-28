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

// Mysql 数据库的配置信息
type Mysql struct {
	URL             string `yaml:"url"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
}

// Redis 缓存的配置信息
type Redis struct {
	Addrs    []string `yaml:"addrs"`
	Password string   `yaml:"password"`
}

// RabbitMQ mq的配置信息
type RabbitMQ struct {
	URL          string `yaml:"url"`
	ExChangeName string `yaml:"exChangeName"`
}

// Sms 短信发送的配置信息
type Sms struct {
	RateLimit  RateLimit    `yaml:"rateLimit"`
	ServerList []*SmsServer `yaml:"serverList"`
}

// SmsServer 短信服务商的具体配置
type SmsServer struct {
	Server       string `yaml:"server"`
	AccessKeyId  string `yaml:"accessKeyId"`
	AccessSecret string `yaml:"accessSecret"`
	GatewayURL   string `yaml:"gatewayURL"`
}

// WeChat 微信公众号发送信息的配置
type WeChat struct {
	RateLimit RateLimit `yaml:"rateLimit"`
	AppId     string    `yaml:"appId"`
	AppSecret string    `yaml:"appSecret"`
}

// Email 邮件发送的配置
type Email struct {
	RateLimit  RateLimit      `yaml:"rateLimit"`
	ServerList []*EmailServer `yaml:"serverList"`
}

// EmailServer 邮件发送服务商的配置
type EmailServer struct {
	Server   string `yaml:"server"`
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	TLS      bool   `yaml:"tls"`
}

// Corn 定时任务的配置
type Corn struct {
	Interval int64 `yaml:"interval"`
	MaxLen   int64 `yaml:"maxLen"`
}

// Config 用于读取具体的yaml
type Config struct {
	Mysql  *Mysql    `yaml:"mysql"`
	MQ     *RabbitMQ `yaml:"mq"`
	Redis  *Redis    `yaml:"redis"`
	Sms    *Sms      `yaml:"sms"`
	WeChat *WeChat   `yaml:"wechat"`
	Email  *Email    `yaml:"email"`
	Corn   *Corn     `yaml:"corn"`
}

// RateLimit 发送频次限制
type RateLimit struct {
	Every1Min  int `yaml:"every1Min"`
	Every1Hour int `yaml:"every1Hour"`
	Every1Day  int `yaml:"every1Day"`
}

var (
	conf = &Config{}
)

// Init 根据文件的路径初始化配置信息
func Init(path string) error {
	return gconf.Read2Struct(path, conf)
}

// MysqlConf 获取数据库的配置信息
func MysqlConf() *Mysql {
	return conf.Mysql
}

// MQConf 获取mq的配置信息
func MQConf() *RabbitMQ {
	return conf.MQ
}

// RedisConf 获取缓存的配置信息
func RedisConf() *Redis {
	return conf.Redis
}

// SmsConf 获取短信服务商的配置
func SmsConf() map[string]*SmsServer {
	var res = make(map[string]*SmsServer)
	for _, v := range conf.Sms.ServerList {
		res[v.Server] = v
	}
	return res
}

// WeChatConf 获取微信的配置
func WeChatConf() *WeChat {
	return conf.WeChat
}

// EmailConf 获取邮件的配置信息
func EmailConf() map[string]*EmailServer {
	var res = make(map[string]*EmailServer)
	for _, v := range conf.Email.ServerList {
		res[v.Server] = v
	}
	return res
}

// CornConf 获取定时任务的配置
func CornConf() *Corn {
	return conf.Corn
}
