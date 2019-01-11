/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : meta.go
#   Created       : 2019/1/10 16:16
#   Last Modified : 2019/1/10 16:16
#   Describe      :
#
# ====================================================*/
package meta

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"time"
	"uuabc.com/sendmsg/pkg/errors"
	"uuabc.com/sendmsg/pkg/utils"
)

// email begin...........

// TypeName 返回接口的类型
func (m *EmailProducer) TypeName() string {
	return "email"
}

// Delay 返回延迟发送的时间 毫秒 单位
func (m *EmailProducer) Delay() int64 {
	return delay(m.XUtcSendStamp)
}

// Validated 验证参数是否合法
func (m *EmailProducer) Validated() error {
	return nil
}

// Transfer 将必须的参数进行转换
func (m *EmailProducer) Transfer() {
	m.Id = uuid.NewV4().String()
	st := gbfToUTC(m.SendTime)
	m.SendTime = st.Format("2006-01-02T15:04:05")
	m.XUtcSendStamp = st.Unix()
}

// email end...........

// sms begin...........
// TypeName 返回接口类型
func (m *SmsProducer) TypeName() string {
	return "sms"
}

// Delay 返回延迟发送的时间 毫秒 单位
func (m *SmsProducer) Delay() int64 {
	return delay(m.XUtcSendStamp)
}

// Transfer 必要的参数转换
func (m *SmsProducer) Transfer() {
	m.Id = uuid.NewV4().String()
	st := gbfToUTC(m.SendTime)
	m.SendTime = st.Format("2006-01-02T15:04:05")
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
	if err := checkServer(m.Server); err != nil {
		return err
	}
	if err := checkContent(m.Content); err != nil {
		return err
	}
	if err := checkSmsTemplate(m.Template); err != nil {
		return err
	}
	if err := checkArguments(m.Arguments); err != nil {
		return err
	}
	if err := checkType(m.Type); err != nil {
		return err
	}
	return nil
}

// sms end...........

// wechat begin..........
// TypeName 接口的类型
func (m *WeChatProducer) TypeName() string {
	return "weixin"
}

// Delay 延迟发送的时间 毫秒 单位
func (m *WeChatProducer) Delay() int64 {
	return delay(m.XUtcSendStamp)
}

// Validated 参数是否合法
func (m *WeChatProducer) Validated() error {
	return nil
}

// Transfer 必要的参数合法
func (m *WeChatProducer) Transfer() {
	m.Id = uuid.NewV4().String()
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

func checkSendTime(s string) error {
	t, err := time.Parse("2006-01-02T15:04:05Z07:00", s)
	if t.Sub(time.Now()) > time.Hour*24*30 {
		return errors.ErrSendTimeTooLong
	}
	if err != nil {
		return errors.ErrTimeFormat
	}
	return nil
}

func checkServer(s int32) error {
	if _, ok := Server_name[s]; !ok || s == 0 {
		return errors.ErrMsgServerNotFound
	}
	return nil
}

func checkContent(s string) error {
	if s == "" {
		return errors.ErrMsgIsNil
	}
	return nil
}

func checkSmsTemplate(s string) error {
	if s == "" || !utils.ValidateTemplate(s) {
		return errors.ErrTemplateNo
	}
	return nil
}

func checkArguments(s string) error {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	if s == "" || err != nil {
		return errors.ErrTemplateParam
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
