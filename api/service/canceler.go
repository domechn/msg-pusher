/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : canceler.go
#   Created       : 2019/1/10 19:58
#   Last Modified : 2019/1/10 19:58
#   Describe      :
#
# ====================================================*/
package service

var Canceler canceler

type canceler struct {
}

func (canceler) Cancel(id string) error {

	return nil
}
