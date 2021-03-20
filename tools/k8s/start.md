#### 基础

- 如何在`k8s` 部署一个应用？
- 如何使用`secret` `configmap`
- 如何挂载和使用`pv`
- 如何配置`ingress controller` 让你的集群里的服务暴露出去
- 如何配置`coredns` 实现集群之间通信
- 如何配置私有仓库

#### 探索

- webhook 的实现机制
- list-watch 是如何实现的
- k8s 的多集群管理设计
- k8s 是如何做到高可用的
- k8s 的网络模型

#### 源码

- kube-scheduler 的一主多从是如何实现的
- kubectl 的源码实现
- deploy controller 的实现细节是什么
- api-server 的高可用的实现机制

#### 自我实现

- 开源备份还原工具 velero 源码学习
- 安全引擎 kyverno 源码学习
- 自己写一个 controller，实现 helm 包的自动运维
- 自己实现一个 webhook 实现环境变量的自动导入
