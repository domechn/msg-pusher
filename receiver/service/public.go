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
	"strconv"
	"strings"

	"github.com/domgoer/msgpusher/pkg/errors"
	"github.com/domgoer/msgpusher/pkg/utils"
	"github.com/domgoer/msgpusher/storer/cache"
)

type ArgParams map[string]string

// checkTemplateAndArguments 验证参数,并返回具体模板内容和处理好的参数
func checkTemplateAndArguments(ctx context.Context, templateID string, args string) (string, ArgParams, error) {
	template, err := cache.LocalTemplate(ctx, templateID)
	params := utils.StrFromCurlyBraces(template)
	if err == nil {
		ag, er := checkArguments(params, args)
		return template, ag, er
	}
	template, err = cache.BaseTemplate(context.Background(), templateID)
	params = utils.StrFromCurlyBraces(template)
	if err != nil {
		return template, nil, errors.ErrTemplateTypeInvalid
	}
	cache.AddLocalTemplate(ctx, templateID, template)
	ag, er := checkArguments(params, args)
	return template, ag, er
}

// checkArguments 查看参数是否合法，是否和模板匹配
func checkArguments(params []string, args string) (ArgParams, error) {
	var ags = make(map[string]interface{})
	if err := json.Unmarshal([]byte(args), &ags); err != nil {
		return nil, err
	}
	var res = make(ArgParams, len(ags))
	var i int
	for a, g := range ags {
		switch g.(type) {
		case string:
			res[a] = g.(string)
		case float64:
			res[a] = strconv.Itoa(int(g.(float64)))
		default:
			return nil, errors.ErrArgumentsInvalid
		}
		var flag bool
		for _, v := range params {
			if v == "${"+a+"}" {
				flag = true
				i++
			}
		}
		if !flag {
			return nil, errors.ErrArgumentsInvalid
		}
	}
	if i != len(params) {
		return nil, errors.ErrArgumentsInvalid
	}
	return res, nil
}

// getContent 将模板内的参数替换成用户传入的值
func getContent(params map[string]string, template string) string {
	for k, v := range params {
		trans := "${" + k + "}"
		template = strings.Replace(template, trans, v, -1)
	}

	return template
}
