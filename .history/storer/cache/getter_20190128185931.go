/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : getter.go
#   Created       : 2019/1/11 17:48
#   Last Modified : 2019/1/11 17:48
#   Describe      :
#
# ====================================================*/
package cache

import (
	"bytes"
	"context"
)

func BaseDetail(ctx context.Context, k string) ([]byte, error) {
	return get(ctx, "BaseDetail", base+k)
}

func BaseTemplate(ctx context.Context, k string) (string, error) {
	b, err := get(ctx, "BaseTemplate", template+k)
	return string(b), err
}

func LastestDetail(ctx context.Context, k string) ([]byte, error) {
	return get(ctx, "LastestDetail", lastest+k)
}

func Detail(ctx context.Context, id string) ([]byte, error) {
	return get(ctx, "Detail", id)
}

// SendResult 获取发送结果
func SendResult(ctx context.Context, k string) (bool, error) {
	res, err := get(ctx, "SendResult", k+"_sent")
	if err != nil {
		return false, err
	}
	if bytes.Compare(res, success) == 0 {
		return true, nil
	}
	return false, nil
}
