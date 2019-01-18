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

func (d *DbEmail) SetContent(s string) {
	d.Content = s
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

func (d *DbEmail) GetSendTo() string {
	return d.Destination
}

func (d *DbEmail) SetSendTo(s string) {
	d.Destination = s
}

func (d *DbEmail) SetArguments(s string) {
	d.Arguments = s
}

func (d *DbEmail) SetTemplate(s string) {
	d.Template = s
}

func (d *DbEmail) SetSendTime(s string) {
	d.SendTime = s
}

func (d *DbEmail) SetTryNum(s int32) {
	d.TryNum = s
}

func (d *DbEmail) SetReason(s string) {
	d.Reason = s
}

func (d *DbWeChat) SetContent(s string) {
	d.Content = s
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

func (d *DbWeChat) GetSendTo() string {
	return d.Touser
}

func (d *DbWeChat) SetSendTo(s string) {
	d.Touser = s
}

func (d *DbWeChat) SetArguments(s string) {
	d.Arguments = s
}

func (d *DbWeChat) SetTemplate(s string) {
	d.Template = s
}

func (d *DbWeChat) SetSendTime(s string) {
	d.SendTime = s
}

func (d *DbWeChat) SetTryNum(s int32) {
	d.TryNum = s
}

func (d *DbWeChat) SetReason(s string) {
	d.Reason = s
}

func (d *DbSms) SetContent(s string) {
	d.Content = s
}

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

func (d *DbSms) GetSendTo() string {
	return d.Mobile
}

func (d *DbSms) SetSendTo(s string) {
	d.Mobile = s
}

func (d *DbSms) SetArguments(s string) {
	d.Arguments = s
}

func (d *DbSms) SetTemplate(s string) {
	d.Template = s
}

func (d *DbSms) SetSendTime(s string) {
	d.SendTime = s
}

func (d *DbSms) SetTryNum(s int32) {
	d.TryNum = s
}

func (d *DbSms) SetReason(s string) {
	d.Reason = s
}
