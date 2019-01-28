/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : mysql.go
#   Created       : 2019/1/9 15:56
#   Last Modified : 2019/1/9 15:56
#   Describe      :
#
# ====================================================*/
package db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	defaultMaxIdleConns    = 10
	defaultMaxOpenConns    = 30
	defaultConnMaxLifetime = 1800
)

type Config struct {
	URL             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

func New(cfg Config) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("mysql", cfg.URL)
	if err != nil {
		return nil, err
	}
	mic := defaultMaxIdleConns
	moc := defaultMaxOpenConns
	cml := defaultConnMaxLifetime
	if cfg.MaxIdleConns > 0 {
		mic = cfg.MaxIdleConns
	}
	if cfg.MaxOpenConns > 0 {
		moc = cfg.MaxOpenConns
	}
	if cfg.ConnMaxLifetime > 0 {
		cml = cfg.ConnMaxLifetime
	}

	db.SetMaxIdleConns(mic)
	db.SetMaxOpenConns(moc)
	db.SetConnMaxLifetime(time.Second * time.Duration(cml))

	return db, db.Ping()
}
