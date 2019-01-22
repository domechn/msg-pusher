/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : const.go
#   Created       : 2019/1/14 11:03
#   Last Modified : 2019/1/14 11:03
#   Describe      :
#
# ====================================================*/
package cache

const (
	base     = "base-"
	lastest  = "lst-"
	template = "template_"
	lock5s   = "lock-5s-"
	lockId   = "lock-id-"
)

var (
	success = []byte{1}
)

type Marshaler interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}
