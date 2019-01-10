/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : config.go
#   Created       : 2019/1/7 14:56
#   Last Modified : 2019/1/7 14:56
#   Describe      :
#
# ====================================================*/
package mail

// Config 用于连接邮件服务器的配置
type Config struct {
	// 邮件服务的地址，host:port
	ServerAddr string
	// 登陆邮件服务的用户名
	Username string
	// 登陆密码
	Password string
	// 邮件服务的域名
	Host string
}
