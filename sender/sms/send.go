/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : send.go
#   Created       : 2019/1/16 16:44
#   Last Modified : 2019/1/16 16:44
#   Describe      :
#
# ====================================================*/
package sms

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/pkg/send/sms"
	"uuabc.com/sendmsg/pkg/utils"
	"uuabc.com/sendmsg/sender/pub"
	"uuabc.com/sendmsg/storer/cache"
)

func (r *Receiver) check(data []byte, msg pub.Messager) error {
	id := string(data)
	if err := r.checkID(id); err != nil {
		return err
	}
	if err := r.checkStatus(id, msg); err != nil {
		return err
	}
	if err := r.checkSendTime(msg); err != nil {
		return err
	}
	return nil
}

func (r *Receiver) checkID(id string) error {
	if err := utils.ValidateUUIDV4(id); err != nil {
		logrus.WithFields(logrus.Fields{
			"data": id,
			"type": r.queueName,
		}).Error("从mq中接收到的数据不符合要求")
		return err
	}
	return nil
}

func (r *Receiver) checkStatus(id string, msg pub.Messager) error {
	b, err := cache.BaseDetail(context.Background(), id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"data":  id,
			"error": err,
			"type":  r.queueName,
		}).Error("从cache中获取数据失败")
		return err
	}
	if err := msg.Unmarshal(b); err != nil {
		logrus.WithFields(logrus.Fields{
			"data":  string(b),
			"error": err,
			"type":  r.queueName,
		}).Error("数据转码异常")
		return err
	}
	status := msg.GetStatus()
	if status != int32(meta.Status_Wait) {
		logrus.WithFields(logrus.Fields{
			"data":   string(b),
			"status": meta.Status_name[status],
			"error":  err,
			"type":   r.queueName,
		}).Error("数据状态不是待发送状态，可能已被处理")
		return pub.ErrMsgHasDealed
	}
	return nil
}

// checkValidate 消息的发送时间是否符合要求，超过当前时间15分钟就不发送，或者发送时间未到
func (r *Receiver) checkSendTime(msg pub.Messager) error {
	now := time.Now().UTC()
	sendT, err := time.Parse("2006-01-02T15:04:05Z", msg.GetSendTime())
	if err != nil {
		return err
	}
	if sendT.Add(pub.DefaultExpirationTime).Before(now) {
		return pub.ErrMsgIsExpiration
	}
	if sendT.After(now) {
		return pub.ErrMsgDeliveryNotArrived
	}
	return nil
}

func (r *Receiver) send(msg pub.Messager) pub.RetryFunc {
	return func(count int) error {
		smsMsg := msg.(*meta.DbSms)
		if count > pub.TryNum {
			smsMsg.SetStatus(int32(meta.Status_Final))
			smsMsg.SetResult(int32(meta.Result_Fail))
			smsMsg.SetTryNum(pub.TryNum)
			updateDbAndCache(smsMsg)
			return pub.ErrTooManyTimes
		}

		var res *sms.Response

		err := pub.SmsClient.Send(sms.NewRequest(
			smsMsg.Mobile,
			"UUabc",
			smsMsg.Template,
			smsMsg.Arguments,
			"12345",
		), func(rs interface{}) {
			if rs != nil {
				res = rs.(*sms.Response)
			}
		})
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"res": res,
			}).Errorf("短信发送失败，error: %v", err)
			return pub.ErrMsgSendFailed
		}

		smsMsg.SetStatus(int32(meta.Status_Final))
		smsMsg.SetResult(int32(meta.Result_Success))
		smsMsg.SetTryNum(int32(count))
		// 更新数据库和缓存，如果出错打印日志，不做错误处理
		updateDbAndCache(smsMsg)

		return nil
	}
}

func updateDbAndCache(msg *meta.DbSms) {
	var err error
	logrus.WithFields(logrus.Fields{
		"id":    msg.Id,
		"data":  msg,
		"error": err,
	}).Errorf("消息发送成功但是更新缓存和数据库时发生错误，请手动修改")

}
