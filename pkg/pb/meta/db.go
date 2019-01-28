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

func (d *DbMsg) SetContent(s string) {
	d.Content = s
}

func (d *DbMsg) SetCreatedAt(s string) {
	d.CreatedAt = s
}

func (d *DbMsg) SetUpdatedAt(s string) {
	d.UpdatedAt = s
}

func (d *DbMsg) SetArguments(s string) {
	d.Arguments = s
}

func (d *DbMsg) SetTemplate(s string) {
	d.Template = s
}

func (d *DbMsg) SetSendTime(s string) {
	d.SendTime = s
}

func (d *DbMsg) SetTryNum(s int32) {
	d.TryNum = s
}

func (d *DbMsg) SetReason(s string) {
	d.Reason = s
}

func (d *DbMsg) SetVersion(s int32) {
	d.Version = s
}

func (m *DbMsg) SetStatus(s Status) {
	m.Status = s
}

func (m *DbMsg) SetResult(s Result) {
	m.ResultStatus = s
}

func (m *DbMsg) SetSendTo(s string) {
	m.SendTo = s
}
