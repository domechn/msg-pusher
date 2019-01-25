/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : init.go
#   Created       : 2019/1/9 15:55
#   Last Modified : 2019/1/9 15:55
#   Describe      :
#
# ====================================================*/
package storer

import (
	"github.com/domgoer/msgpusher/config"
	"github.com/domgoer/msgpusher/pkg/cache/local"
	"github.com/domgoer/msgpusher/pkg/cache/redis"
	"github.com/domgoer/msgpusher/pkg/db"
	"github.com/domgoer/msgpusher/pkg/mq"
	"github.com/jmoiron/sqlx"
)

var (
	MqCli        *mq.RabbitConn
	ExChangeName string
	DB           *sqlx.DB
	Cache        *redis.Client
	LocalCache   *local.Client
)

func Init() (err error) {
	mqConf := config.MQConf()
	MqCli, err = mq.New(mqConf.URL)
	ExChangeName = mqConf.ExChangeName
	if err != nil {
		return err
	}

	mysqlConf := config.MysqlConf()
	DB, err = db.New(db.Config{
		URL:             mysqlConf.URL,
		MaxIdleConns:    mysqlConf.MaxIdleConns,
		MaxOpenConns:    mysqlConf.MaxOpenConns,
		ConnMaxLifetime: mysqlConf.ConnMaxLifetime,
	})
	if err != nil {
		return err
	}

	redisConf := config.RedisConf()
	Cache, err = redis.NewClient(redisConf.Addrs, redisConf.Password)
	LocalCache = local.NewClient()
	return err
}

func Close() error {
	MqCli.Close()
	Cache.Close()
	DB.Close()
	return nil
}
