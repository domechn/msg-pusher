/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : msg.go
#   Created       : 2019/1/28 11:37
#   Last Modified : 2019/1/28 11:37
#   Describe      :
#
# ====================================================*/
package db

import (
	"context"
	"github.com/domgoer/msg-pusher/pkg/pb/meta"
)

// UpdateAndInsertMsgBatch 批量执行修改,如果不存在就插入
func UpdateAndInsertMsgBatch(ctx context.Context, ds []*meta.DbMsg) error {
	var args []interface{}
	for _, s := range ds {
		args = append(args,
			s.Id,
			s.SubId,
			s.Content,
			s.SendTo,
			s.Type,
			s.Template,
			s.Arguments,
			s.Server,
			s.Reserved,
			s.SendTime,
			s.TryNum,
			s.Status,
			s.ResultStatus,
			s.Reason,
			s.Version)
	}
	err := batch(ctx,
		"msg",
		[]string{"id", "sub_id", "content", "send_to", "type", "template", "arguments", "server", "reserved", "send_time", "try_num", "status", "result_status", "reason", "version"},
		args...,
	)
	return err
}

// DetailByToAndPage 根据接收人的标识分页查询
func DetailByToAndPage(ctx context.Context, to string, page int) ([]*meta.DbMsg, error) {
	var res []*meta.DbMsg
	size := (page - 1) * 10
	err := list(ctx,
		&res,
		"DetailByToAndPage",
		"SELECT * FROM msg WHERE to=? LIMIT ?,10",
		to,
		size)
	return res, err
}

// DetailByKey 根据key分页查询
func DetailByKey(ctx context.Context, key string, page int) ([]*meta.DbMsg, error) {
	var res []*meta.DbMsg
	size := (page - 1) * 10
	err := list(ctx,
		&res,
		"DetailByToAndPage",
		"SELECT * FROM msg WHERE key=? LIMIT ?,10",
		key,
		size)
	return res, err
}

// WaitingMsgByKey 根据key值返回所有待发送的消息的id
func WaitingMsgByKey(ctx context.Context, key string) ([]string, error) {
	var res []string
	err := list(ctx,
		"WaitingMsgByKey",
		"SELECT id FROM msg WHERE key=? AND status = 0",
		key)
	return res, err
}
