/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : validate.go
#   Created       : 2019/1/15 15:22
#   Last Modified : 2019/1/15 15:22
#   Describe      :
#
# ====================================================*/
package tpl

import (
	"uuabc.com/sendmsg/pkg/errors"
)

func (m *TemplateAdder) Validate() error {
	if m.Content == "" {
		return errors.ErrTemplateContentIsNil
	}
	_, ok := TemplateType_name[m.Type]
	if m.Type == 0 || !ok {
		return errors.ErrTemplateTypeInvalid
	}
	if m.SimpleID == "" {
		return errors.ErrTemplateSimpleInvalid
	}
	return nil
}
