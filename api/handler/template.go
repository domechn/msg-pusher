/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : template.go
#   Created       : 2019/1/15 15:01
#   Last Modified : 2019/1/15 15:01
#   Describe      :
#
# ====================================================*/
package handler

import (
	"context"
	"uuabc.com/sendmsg/api/model"
	"uuabc.com/sendmsg/api/service"
	"uuabc.com/sendmsg/pkg/pb/tpl"
)

// @router("POST","/version/tpl")
// TemplateAdd 添加模板
func TemplateAdd(ctx context.Context, data []byte) (res []byte, err error) {
	p := &tpl.TemplateAdder{}
	if err = json.Unmarshal(data, p); err != nil {
		return
	}

	if err = p.Validate(); err != nil {
		return
	}

	var id string
	if id, err = service.TemplateImpl.AddTemplate(ctx, p); err != nil {
		return
	}

	res = model.NewResponseDataKey("id", id).MustMarshal()
	return
}
