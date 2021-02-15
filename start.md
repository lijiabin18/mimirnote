#### 启动概述

##### 预备知识

> 库

- [viper](https://zhuanlan.zhihu.com/p/103522811) ，配置文件

  - 监听配置文件修改，修改时自动加载新配置
  - 从环境变量、命令行选项和 io.Reader 中读取配置
  - 从远程配置系统中读取和监听修改，如 etcd/Consul
  - 代码逻辑中显示设置键值

- cobra ，命令行程序库，可以用来编写命令行程序
- Prometheus ，Kubernetes 集群的监控系统

##### 启动入口 main

> op server 为例

```go
bootstrap.HTTPServer(bootstrap.GetOpBootstrap())

```

- 创建 HTTPServer, 是在 Bootstrap 中配置
  - 读取配置信息，初始化 mq、mysql、mongo、room、usercenterclient、kafka、redis)
  - 初始化服务器，配置 MonitoringMiddleware 与 MongoMiddleware 中间件

> 初始化 server

- basicInit()

  - initInternalWithArgs() 初始化 sentry
  - InitializeServerWithHealthySignalAndArgs() 创建一个`server` 并运行，
    设置请求 Header(Content-Type)、监控(Prometheus)，链路追踪等

- 整个初始化过程是异步的，通过一个名为`healthy` 的 有缓冲 channel
  实现，所有相关服务初始化成功后，向 `healthy` 写入 true，
  无异常后，则一个模块的服务器创建完毕。接下来就是为这个服务器注册路由

```go
func main() {
	flag.Parse()
	cmd.Execute(
		cmd.HttpServerCmd,
		cmd.GrpcServerCmd,
		cmd.KafkaConsumerCmd,
		cmd.RechargeKafkaConsumerCmd,
		cmd.HttpServerOpCmd,
		cmd.ReconciliationCmd,
		cmd.TimerBillingWorkerCmd,
		cmd.LeaderBoardRewardCmd,
		cmd.RefundRecoveryCmd,
		cmd.MqCmd,
		cmd.RechargeStatsKafkaConsumerCmd,
		cmd.AutoActiveUserToolCronJob,
	)
}

```
