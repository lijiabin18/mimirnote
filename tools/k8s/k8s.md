#### 单集群架构

- master 核心组件(Kubernetes Control Plane)

  - `kube-api-server` : `k8s 请求入口服务` ，对外提供`系统调用` ,k8s 资源调用
  - `Scheduler` ：`k8s 所有 work node 的监控器`，当用户部署服务时，Scheduler 选择合适的`work node 来部署`
  - `etcd` : k8s 存储单元，只有**kubu-api-server** 会直接和 etcd 打交道。
  - `kube-controller manager` ：管理所有资源

- Kubenetes Nodes
  - kubectl: 通过监听 kube-api-server 来`创建/删除` 对应的 pod
  - kube-proxy: 根据 kubectl 创建/删除 的 pod
    来 创建(计算)当前节点上的网络路由规则，从而实现 pod 或 service 的网络
    流量流转规则
  - 容器引擎：docker 或 contained 等引擎，这些引擎是与 k8s
    解耦的，k8s 通过 cri 接口 和运营时的引擎打交道

#### k8s 架构 与 linux

- `linux 子系统 进程` 与 `controller` ：

  - 不同的 linux 子系统 可以执行/分配不同的资源(不同的进程可以实现不同的功能)。
  - k8s 中 不同的 controller 可以监听不同的资源，并且实现资源的一致性

- `linux Scheduler` 与 `k8s Scheduler`
  - linux 通过 Scheduler 来调度不同的进程
  - k8s 的 Scheduler 调用每个设备上的`操作系统`，pod 。k8s
    由很多个设备组成，每个设备是由一个操作系统组成，所以 `k8s 的 Scheduler 层级更高`

#### k8s 高可用

- 高可用 架构设计？
- 多 master 间如何实现负载均衡(api-sever)，使用什么算法实现 `etcd 一致性`？(api server 通过 balancer 实现，etcd 通过 raft 算法来通信实现，etcd 主从选举实现高可用，leader 内部的`scheduler controller cluster` 通过一主多从(1 master node，n work node)实现)
  - 主从架构(etcd cluster)。当某个 `leader` (master node) 故障时，会通过选举选一个 `从master 成为 leader` 保证高可用
  - 可以在 kube-system 命令空间下看到名字为 kube-scheduler 的 endpoint
  - 当某个 master node 设置了 `leader-elect` 参数时，在启动时会尝试去获取 \*\*\*\*
    leader 节点的身份。当多个 master-node 设置了 `leader-elect`
    参数时，会在启动时创建名字为 kube-scheduler 的 endpoint，可以在`kube-system` 中看到。
    成为 leader 后，需要定期更新自己的 endpoint 信息，来维护自己 leader 的身份，而其他从节点需要定期检查 endpoint 的信息，如果在某个时间段内没有更新这，就说明原来的 leader 故障，从节点可以通过选举 成为 新的 leader
