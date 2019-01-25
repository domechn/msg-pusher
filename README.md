# msg-pusher

### msg-pusher是用Golang编写的高性能消息推送平台

### 使用msg-pusher可以实现：

- 消息的实时推送
- 消息的定时推送

### 目前支持一下平台的推送：

- 阿里云短信
- 微信公众号
- 邮件服务

### 调用关系

### 服务集成

服务集成了 [prometheus](https://github.com/prometheus/prometheus)、[jaeger](https://github.com/jaegertracing/jaeger) ,你可以通过这些插件来观察msg-pusher的响应情况和性能状况

### 服务部署

msg-pusher依赖rabbit-mq、redis和mysql

> 源码部署

> docker部署


