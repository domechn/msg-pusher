/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : public.go
#   Created       : 2019/1/24 14:37
#   Last Modified : 2019/1/24 14:37
#   Describe      :
#
# ====================================================*/
package db

func Register() {
	registerMsg()
}

func read(lenF func() (int64, error), popF func() ([]byte, error), n int64) ([][]byte, error) {
	var res [][]byte
	// 查看list的长度
	len, err := lenF()
	if err != nil {
		return nil, err
	}
	// 如果超过单次最大量
	if len > n {
		len = n
	}

	for i := 0; i < int(len); i++ {
		b, err := popF()
		if err != nil {
			continue
		}
		// 说明已经读完
		if b == nil {
			return res, nil
		}
		res = append(res, b)
	}
	return res, nil
}
