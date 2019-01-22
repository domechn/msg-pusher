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
	ErrSmsServerNotFound = NewError(106001001, "短信服务商不存在")
	ErrPageInvalidate    = NewError(106001002, "查询页数不合法或超过范围(1~10)")
	ErrMsg1MinuteLimit   = NewError(106001003, "1分钟内短信发送数量到达上限")
	ErrMsg1HourLimit     = NewError(106001004, "1小时内短信发送数量到达上限")
	ErrMsg1DayLimit      = NewError(106001005, "24小时内短信发送数量到达上限")
	ErrPhoneNumber       = NewError(106001006, "手机号码格式错误")

	ErrEmailServerNotFound = NewError(106003001, "邮件服务商不存在")
	ErrDestination         = NewError(106003002, "接收邮箱不合法")

	ErrTemplateIsExisted     = NewError(106004001, "模板SimpleID已存在")
	ErrTemplateSimpleInvalid = NewError(106004002, "SimpleID不能为空")
	ErrTemplateContentIsNil  = NewError(106004003, "模板内容为空")
	ErrTemplateTypeInvalid   = NewError(106004004, "模板类型不支持")

	ErrParam              = NewError(106009001, "参数格式错误")
	ErrTimeFormat         = NewError(106009002, "时间格式不符合ISO8601标准")
	ErrToUser             = NewError(106009003, "接收者为空")
	ErrArgumentsInvalid   = NewError(106009004, "参数和模板不匹配")
	ErrMsgTypeNotFound    = NewError(106009005, "消息类型不存在")
	ErrPlatNotFound       = NewError(106009006, "请求平台不存在")
	ErrPlatKeyIsNil       = NewError(106009007, "平台key内容为空")
	ErrMisMatch           = NewError(106009008, "请求接口和消息类型不匹配")
	ErrIDIsInvalid        = NewError(100009009, "查询id为空或格式错误")
	ErrMsgNotFound        = NewError(106009010, "消息不存在")
	ErrMsgHasCancelled    = NewError(106009011, "此消息已取消发送")
	ErrMsgCantEdit        = NewError(106009012, "此消息已无法修改")
	ErrSendTimeTooLong    = NewError(106009013, "无法延迟发送间隔超过一个月的消息")
	ErrMsgBusy            = NewError(106009014, "该消息正在处理")
	ErrMsgIsSameBefore    = NewError(106009015, "消息数据和原先一致")
	ErrFunctionNotSupport = NewError(106009016, "功能未实现")
)
