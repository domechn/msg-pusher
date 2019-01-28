/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : meta.go
#   Created       : 2019/1/28 13:26
#   Last Modified : 2019/1/28 13:26
#   Describe      :
#
# ====================================================*/
package meta

import (
	"github.com/domgoer/msg-pusher/pkg/errors"
	"github.com/domgoer/msg-pusher/pkg/utils"
	"github.com/satori/go.uuid"
	"time"
)

const (
	ISO8601Layout = "2006-01-02T15:04:05Z07:00"
)

// Transfer 必要的参数转换
func (m *MsgProducer) Transfer(setID bool) {
	if setID {
		m.Id = uuid.NewV4().String()
	}
	var st time.Time
	st, m.SendTime = gbfToUTC(m.SendTime)
	m.XUtcSendStamp = st.Unix()
}

// Delay 返回延迟发送的时间 毫秒 单位
func (m *MsgProducer) Delay() int64 {
	return delay(m.XUtcSendStamp)
}

// Validated 验证参数时候合法
func (m *MsgProducer) Validated() error {
	if err := checkSubId(m.SubId); err != nil {
		return err
	}
	if err := checkType(m.Type); err != nil {
		return err
	}
	if err := checkServer(m.Server); err != nil {
		return err
	}
	if err := checkSendTo(m.Type, m.SendTo); err != nil {
		return err
	}
	if err := checkSendTime(m.SendTime); err != nil {
		return err
	}

	return nil
}

func checkSubId(s string) error {
	if s == "" {
		return errors.ErrPlatKeyIsNil
	}
	return nil
}

func checkSendTo(t int32, s string) error {
	switch t {
	case int32(Sms):
		if !utils.ValidatePhone(s) {
			return errors.ErrPhoneNumber
		}
	case int32(Email):
		if !utils.ValidateEmailAddr(s) {
			return errors.ErrDestination
		}
	}
	return nil
}

func checkSendTime(s string) error {
	if s == "0" {
		return nil
	}
	t, err := time.Parse(ISO8601Layout, s)
	if t.Sub(time.Now()) > time.Hour*24*30 {
		return errors.ErrSendTimeTooLong
	}
	if err != nil {
		return errors.ErrTimeFormat
	}
	return nil
}

func checkType(s int32) error {
	if _, ok := Type_name[s]; !ok {
		return errors.ErrMsgTypeNotFound
	}
	return nil
}

func checkServer(s int32) error {
	if _, ok := Server_name[s]; !ok {
		return errors.ErrServerNotFound
	}
	return nil
}

func gbfToUTC(s string) (time.Time, string) {
	if s == "0" {
		utc := time.Now().UTC()
		return utc, utc.Format(ISO8601Layout)
	}
	st, _ := time.Parse(ISO8601Layout, s)
	sts := st.UTC()
	return sts, s
}

func delay(begin int64) int64 {
	ns := time.Now().Unix()
	d := (begin - ns) * 1000
	if d < 0 {
		d = 0
	}
	return d
}

func (m *MsgProducer) SetSendTo(s string) {
	m.SendTo = s
}

func (m *MsgProducer) ValidateEdit() error {
	return nil
}
