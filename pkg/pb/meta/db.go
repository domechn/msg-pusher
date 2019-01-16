/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : db.go
#   Created       : 2019/1/16 13:47
#   Last Modified : 2019/1/16 13:47
#   Describe      :
#
# ====================================================*/
package meta

func (d *DbSms) SetStatus(s int32) {
	d.Status = s
}

func (d *DbSms) SetResult(s int32) {
	d.ResultStatus = s
}

func (d *DbSms) SetCreatedAt(s string) {
	d.CreatedAt = s
}

func (d *DbSms) SetUpdatedAt(s string) {
	d.UpdatedAt = s
}

func (d *DbSms) SetTryNum(s int32) {
	d.TryNum = s
}

func (d *DbEmail) SetStatus(s int32) {
	d.Status = s
}

func (d *DbEmail) SetResult(s int32) {
	d.ResultStatus = s
}

func (d *DbEmail) SetCreatedAt(s string) {
	d.CreatedAt = s
}

func (d *DbEmail) SetUpdatedAt(s string) {
	d.UpdatedAt = s
}

func (d *DbEmail) SetTryNum(s int32) {
	d.TryNum = s
}

func (d *DbWeChat) SetStatus(s int32) {
	d.Status = s
}

func (d *DbWeChat) SetResult(s int32) {
	d.ResultStatus = s
}

func (d *DbWeChat) SetCreatedAt(s string) {
	d.CreatedAt = s
}

func (d *DbWeChat) SetUpdatedAt(s string) {
	d.UpdatedAt = s
}

func (d *DbWeChat) SetTryNum(s int32) {
	d.TryNum = s
}
