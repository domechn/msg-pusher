/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : meta.go
#   Created       : 2019/1/10 16:16
#   Last Modified : 2019/1/10 16:16
#   Describe      : 用于检查参数是否合法，注意：本层不检查 模板和模板参数，
#					具体这两项检查需要结合业务在service层做.
#
# ====================================================*/
package meta

// TODO snedtime geshibudui
import (
	"time"

	"github.com/json-iterator/go"
	"github.com/satori/go.uuid"
	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/pkg/utils"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	layout = "2006-01-02T15:04:05Z"
)

// email begin...........
// Delay 返回延迟发送的时间 毫秒 单位
func (m *EmailProducer) Delay() int64 {
	return delay(m.XUtcSendStamp)
}

// Validated 验证参数是否合法
func (m *EmailProducer) Validated() error {
	if err := checkPlatform(m.Platform); err != nil {
		return err
	}
	if err := checkPlatformKey(m.PlatformKey); err != nil {
		return err
	}
	if err := checkDestination(m.Destination); err != nil {
		return err
	}
	if err := checkSendTime(m.SendTime); err != nil {
		return err
	}
	if err := checkEmailServer(m.Server); err != nil {
		return err
	}
	if err := checkType(m.Type); err != nil {
		return err
	}
	return nil
}

func (m *EmailProducer) ValidateEdit() error {
	if m.Destination != "" {
		if err := checkDestination(m.Destination); err != nil {
			return err
		}
	}
	if m.SendTime != "" {
		if err := checkSendTime(m.SendTime); err != nil {
			return err
		}
	}
	return nil
}

// Transfer 将必须的参数进行转换
func (m *EmailProducer) Transfer(setID bool) {
	if setID {
		m.Id = uuid.NewV4().String()
	}
	st := gbfToUTC(m.SendTime)
	m.SendTime = st.Format(layout)
	m.XUtcSendStamp = st.Unix()
}

// email end...........

// sms begin...........
// Delay 返回延迟发送的时间 毫秒 单位
func (m *SmsProducer) Delay() int64 {
	return delay(m.XUtcSendStamp)
}

// Transfer 必要的参数转换
func (m *SmsProducer) Transfer(setID bool) {
	if setID {
		m.Id = uuid.NewV4().String()
	}
	st := gbfToUTC(m.SendTime)
	m.SendTime = st.Format(layout)
	m.XUtcSendStamp = st.Unix()
}

// Validated 验证参数时候合法
func (m *SmsProducer) Validated() error {
	if err := checkPlatform(m.Platform); err != nil {
		return err
	}
	if err := checkPlatformKey(m.PlatformKey); err != nil {
		return err
	}
	if err := checkMobile(m.Mobile); err != nil {
		return err
	}
	if err := checkSendTime(m.SendTime); err != nil {
		return err
	}
	if err := checkSmsServer(m.Server); err != nil {
		return err
	}
	if err := checkType(m.Type); err != nil {
		return err
	}
	return nil
}

func (m *SmsProducer) ValidateBatch() error {
	if err := checkPlatformKey(m.PlatformKey); err != nil {
		return err
	}
	if err := checkMobile(m.Mobile); err != nil {
		return err
	}
	if err := checkSendTime(m.SendTime); err != nil {
		return err
	}
	if err := checkSmsServer(m.Server); err != nil {
		return err
	}
	if err := checkType(m.Type); err != nil {
		return err
	}
	return nil
}

func (m *SmsProducer) ValidateEdit() error {
	if m.Mobile != "" {
		if err := checkMobile(m.Mobile); err != nil {
			return err
		}
	}
	if m.SendTime != "" {
		if err := checkSendTime(m.SendTime); err != nil {
			return err
		}
	}
	return nil
}

// sms end...........

// wechat begin..........
// Delay 延迟发送的时间 毫秒 单位
func (m *WeChatProducer) Delay() int64 {
	return delay(m.XUtcSendStamp)
}

// Validated 参数是否合法
func (m *WeChatProducer) Validated() error {
	if err := checkPlatform(m.Platform); err != nil {
		return err
	}
	if err := checkToUser(m.Touser); err != nil {
		return err
	}
	if err := checkSendTime(m.SendTime); err != nil {
		return err
	}
	if err := checkType(m.Type); err != nil {
		return err
	}
	return nil
}

func (m *WeChatProducer) ValidateEdit() error {
	if m.Touser != "" {
		if err := checkToUser(m.Touser); err != nil {
			return err
		}
	}
	if m.SendTime != "" {
		if err := checkSendTime(m.SendTime); err != nil {
			return err
		}
	}
	return nil
}

// Transfer 必要的参数合法
func (m *WeChatProducer) Transfer(setID bool) {
	if setID {
		m.Id = uuid.NewV4().String()
	}
	sendTime := gbfToUTC(m.SendTime)
	m.SendTime = sendTime.Format("2006-01-02T15:04:05")
	m.XUtcSendStamp = sendTime.Unix()
}

// wechat end........

func checkPlatform(s int32) error {
	if _, ok := PlatForm_name[s]; !ok || s == 0 {
		return errors.ErrPlatNotFound
	}
	return nil
}

func checkPlatformKey(s string) error {
	if s == "" {
		return errors.ErrPlatKeyIsNil
	}
	return nil
}

func checkMobile(s string) error {
	if !utils.ValidatePhone(s) {
		return errors.ErrPhoneNumber
	}
	return nil
}

func checkToUser(s string) error {
	if s == "" {
		return errors.ErrToUser
	}
	return nil
}

func checkDestination(s string) error {
	if !utils.ValidateEmailAddr(s) {
		return errors.ErrDestination
	}
	return nil
}

func checkSendTime(s string) error {
	if s == "0" {
		return nil
	}
	t, err := time.Parse("2006-01-02T15:04:05Z07:00", s)
	if t.Sub(time.Now()) > time.Hour*24*30 {
		return errors.ErrSendTimeTooLong
	}
	if err != nil {
		return errors.ErrTimeFormat
	}
	return nil
}

func checkSmsServer(s int32) error {
	if _, ok := SmsServer_name[s]; !ok || s == 0 {
		return errors.ErrSmsServerNotFound
	}
	return nil
}

func checkEmailServer(s int32) error {
	if _, ok := EmailServer_name[s]; !ok || s == 0 {
		return errors.ErrEmailServerNotFound
	}
	return nil
}

func checkType(s int32) error {
	if _, ok := Message_name[s]; !ok || s == 0 {
		return errors.ErrMsgTypeNotFound
	}
	return nil
}

func gbfToUTC(s string) time.Time {
	if s == "0" {
		return time.Now().UTC()
	}
	st, _ := time.Parse("2006-01-02T15:04:05Z07:00", s)
	sts := st.UTC()
	return sts
}

func delay(begin int64) int64 {
	ns := time.Now().Unix()
	d := (begin - ns) * 1000
	if d < 0 {
		d = 0
	}
	return d
}

func (m *EmailProducer) GetSendTo() string {
	return m.Destination
}

func (m *EmailProducer) SetSendTo(s string) {
	m.Destination = s
}

func (m *WeChatProducer) GetSendTo() string {
	return m.Touser
}

func (m *WeChatProducer) SetSendTo(s string) {
	m.Touser = s
}

func (m *SmsProducer) GetSendTo() string {
	return m.Mobile
}

func (m *SmsProducer) SetSendTo(s string) {
	m.Mobile = s
}
