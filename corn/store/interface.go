/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : interface.go
#   Created       : 2019/1/23 16:54
#   Last Modified : 2019/1/23 16:54
#   Describe      :
#
# ====================================================*/
package store

type Corn interface {
	Name() string
	// 从缓存中读取数据，以切片的方式返回
	Read() ([][]byte, error)
	// 将数据写入数据库
	Write([][]byte) error
}

type Marshaler interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}
