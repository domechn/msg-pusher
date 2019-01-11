/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : service.go
#   Created       : 2019/1/11 17:07
#   Last Modified : 2019/1/11 17:07
#   Describe      :
#
# ====================================================*/
package backoff

import (
	"time"
)

// NewServiceBackOff 返回一个贴合项目需要的backoff
func NewServiceBackOff() *ExponentialBackOff {
	back := NewExponentialBackOff()
	back.InitialInterval = time.Millisecond * 100
	back.Multiplier = 1.2
	back.MaxInterval = time.Millisecond * 300
	back.MaxElapsedTime = time.Second * 5
	return back
}
