/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : service.go
#   Created       : 2019/1/10 16:50
#   Last Modified : 2019/1/10 16:50
#   Describe      : 业务错误码
#
# ====================================================*/
package errors

var (
	ErrMsgCantEdit       = NewError(106001002, "此消息已发送,无法修改或取消")
	ErrMsgIsNil          = NewError(106001004, "消息内容为空")
	ErrMsgHasCancelled   = NewError(106001006, "此消息已取消发送,请勿重复操作")
	ErrMsgTypeNotFound   = NewError(106001007, "消息类型不存在")
	ErrTemplateNo        = NewError(106001008, "模板编号错误")
	ErrTemplateParam     = NewError(106001009, "模板参数错误")
	ErrTemplateContent   = NewError(106001010, "消息内容和模板不一致")
	ErrMsgServerNotFound = NewError(106001011, "短信服务商不存在")
	ErrPageInvalidate    = NewError(106001013, "查询页数不合法或超过范围(1~10)")
	ErrMsg1MinuteLimit   = NewError(106001014, "1分钟内短信发送数量到达上限")
	ErrMsg1HourLimit     = NewError(106001015, "1小时内短信发送数量到达上限")
	ErrMsg1DayLimit      = NewError(106001016, "24小时内短信发送数量到达上限")
)

var (
	ErrPhoneNumber     = NewError(100002006, "手机号码格式错误")
	ErrIDIsInvalid     = NewError(100002008, "查询id为空或格式错误")
	ErrTimeFormat      = NewError(100002007, "时间格式不符合ISO8601标准")
	ErrPlatNotFound    = NewError(106001003, "请求平台不存在")
	ErrPlatKeyIsNil    = NewError(106001012, "平台key内容为空")
	ErrSendTimeTooLong = NewError(106001023, "无法延迟发送间隔超过一个月的消息")
	ErrMsgNotFound     = NewError(106001024, "消息不存在")
)
