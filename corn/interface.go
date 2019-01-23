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
package corn

type Corn interface {
	// 从缓存中读取数据，以切片的方式返回
	Read(int64) (inserts []Marshaler, updates []Marshaler)
	// 将数据写入数据库
	Write(inserts, updates []Marshaler) error
	// 开始执行定时任务，将redis中的数据定时入库
	Start()
}

type Marshaler interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}
