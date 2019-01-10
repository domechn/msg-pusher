/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : producer.go
#   Created       : 2019/1/9 15:49
#   Last Modified : 2019/1/9 15:49
#   Describe      :
#
# ====================================================*/
package service

type ProducerImpl producerImpl

type producerImpl struct {
}

func (producerImpl) Produce() error {
	return nil
}
