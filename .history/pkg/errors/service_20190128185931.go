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
	ErrPageInvalidate  = NewError(10001002, "查询页数不合法或超过范围(1~10)")
	ErrMsg1MinuteLimit = NewError(10001003, "1分钟内发送数量到达上限")
	ErrMsg1HourLimit   = NewError(10001004, "1小时内发送数量到达上限")
	ErrMsg1DayLimit    = NewError(10001005, "24小时内发送数量到达上限")
	ErrPhoneNumber     = NewError(10001006, "手机号码格式错误")

	ErrServerNotFound = NewError(10003001, "服务商不存在")
	ErrDestination    = NewError(10003002, "接收邮箱不合法")

	ErrTemplateIsExisted     = NewError(10004001, "模板SimpleID已存在")
	ErrTemplateSimpleInvalid = NewError(10004002, "SimpleID不能为空")
	ErrTemplateContentIsNil  = NewError(10004003, "模板内容为空")
	ErrTemplateTypeInvalid   = NewError(10004004, "模板类型不支持")

	ErrParam              = NewError(10009001, "参数格式错误")
	ErrTimeFormat         = NewError(10009002, "时间格式不符合ISO8601标准")
	ErrToUser             = NewError(10009003, "接收者为空")
	ErrArgumentsInvalid   = NewError(10009004, "参数和模板不匹配")
	ErrMsgTypeNotFound    = NewError(10009005, "消息类型不存在")
	ErrPlatNotFound       = NewError(10009006, "请求平台不存在")
	ErrPlatKeyIsNil       = NewError(10009007, "key内容为空")
	ErrMisMatch           = NewError(10009008, "请求接口和消息类型不匹配")
	ErrIDIsInvalid        = NewError(10009009, "查询id为空或格式错误")
	ErrMsgNotFound        = NewError(10009010, "消息不存在")
	ErrMsgHasCancelled    = NewError(10009011, "此消息已取消发送")
	ErrMsgCantEdit        = NewError(10009012, "此消息已无法修改")
	ErrSendTimeTooLong    = NewError(10009013, "无法延迟发送间隔超过一个月的消息")
	ErrMsgBusy            = NewError(10009014, "该消息正在处理")
	ErrMsgIsSameBefore    = NewError(10009015, "消息数据和原先一致")
	ErrFunctionNotSupport = NewError(10009016, "功能未实现")
)
