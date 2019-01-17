/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : public.go
#   Created       : 2019/1/14 10:36
#   Last Modified : 2019/1/14 10:36
#   Describe      :
#
# ====================================================*/
package service

import (
	"context"
	"encoding/json"
	"strings"

	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/storer/cache"
)

func checkTemplateAndArguments(s string, args string) error {
	params, err := cache.LocalTemplate(s)
	if err == nil {
		return checkArguments(params, args)
	}
	params, err = cache.BaseTemplate(context.Background(), s)
	if err != nil {
		return errors.ErrTemplateTypeInvalid
	}
	cache.AddLocalTemplate(s, strings.Join(params, ","))
	return checkArguments(params, args)
}

func checkArguments(params []string, args string) error {
	var ags = make(map[string]interface{})
	if err := json.Unmarshal([]byte(args), &ags); err != nil {
		return err
	}
	var i int
	for a := range ags {
		var flag bool
		for _, v := range params {
			if v == "${"+a+"}" {
				flag = true
				i++
			}
			if !flag {
				return errors.ErrArgumentsInvalid
			}
		}
	}
	if i != len(params) {
		return errors.ErrArgumentsInvalid
	}
	return nil
}
