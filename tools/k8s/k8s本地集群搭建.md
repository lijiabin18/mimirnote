### docker 安装

1. 更换国内软件源，推荐中国科技大学的源，稳定速度快（可选）

```
sudo cp /etc/apt/sources.list /etc/apt/sources.list.bak
sudo sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
sudo apt update
```

2. 安装所需的包

```
sudo apt install apt-transport-https ca-certificates software-properties-common curl

```

3. 添加 GPG 密钥，并添加 Docker-ce 软件源(阿里源)

```
curl -fsSL https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu \
$(lsb_release -cs) stable"

```

4. 更新软件包缓存，sudo apt update

- Q1: apt update 报错，因为之前安装 docker 配置了官方源，并不能使用
  - exception:Failed to fetch https://apt.dockerproject.org/repo/dists/ubuntu..
  - answer: sudo rm /etc/apt/sources.list.d/docker\*

5. 安装 docker-ce 并启动 docker 服务

```
sudo apt install docker-ce
sudo systemctl enable docker
sudo systemctl start docker
```

6. 测试 docker

```
sudo docker run hello-world
```

7. 添加当前用户到 docker 用户组，可以不用 sudo 运行 docker（可选）

```
sudo groupadd docker
sudo usermod -aG docker $USER
```

8. 测试添加用户组（可选）

```
docker run hello-world
```

### 安装 kubeadm

1. 关闭 swap，`swapoff -a`

2. 安装 kubeadm

```
- 下载GPG并添加GPG
sudo wget https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg

sudo apt-key add apt-key.gpg

- 写入文件

sudo tee /etc/apt/sources.list.d/kubernetes.list <<EOF
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF

- 更新缓存

sudo apt-get update

- 安装kubelet-1.18.3-00 kubeadm=1.18.3-00 kubectl=1.18.3-00

// 无效。sudo apt-get install -y kubelet-1.18.3-00 kubeadm=1.18.3-00 kubectl=1.18.3-00
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl

- 检查安装k8s需要的镜像有哪些
kubeadm config images list

k8s.gcr.io/kube-apiserver:v1.20.5
k8s.gcr.io/kube-controller-manager:v1.20.5
k8s.gcr.io/kube-scheduler:v1.20.5
k8s.gcr.io/kube-proxy:v1.20.5
k8s.gcr.io/pause:3.2
k8s.gcr.io/etcd:3.4.13-0
k8s.gcr.io/coredns:1.7.0

- pull，镜像需要配置代理，否则无法拉取
- ubuntu 使用 qv2ray 代理
	- 安装qv2ray:sudo snap install qv2ray
	- 配置v2ray-core: https://github.com/v2fly/v2ray-core/releases/tag/v4.36.2
	下载 v2ray-linux-64.zip ，然后在home目录下创建执行`mkdir bin/vcore/`
	将压缩文件解压到 该目录下，并在qv2ray中指定 v2ray-core 文件位置
	bin/vc
	- (可选)配置终端代理：sudo apt install tor && apt install pproxychains4 && sudo nvim
	/etc/proxychains.conf 添加代理

- docker pull 配置代理
sudo mkdir -p /etc/systemd/system/docker.service.d && vim /etc/systemd/system/docker.service.d/proxy.conf

[Service]
Environment="HTTPS_PROXY=http://127.0.0.1:8889/" "NO_PROXY=localhost,127.0.0.1,registry.docker-cn.com,hub-mirror.c.163.com"

```

<++>

3. 初始化

```
- sudo kubeadm init --kubernetes-version=1.20.5 --pod-network-cidr 10.244.0.0/16
	- pod-network-cidr: ip 地址通信范围，10.244.0.0/16
	- image-repository: 若本地没镜像时，拉取镜像的仓库

```

4. 访问集群

```
- mkdir -p ~/.kube && sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
&& sudo chown $(id -u):$(id -g) $HOME/.kube/config && kubectl get pod -A

- 查询本机 k8s 进程。`ps -aux | grep kube`

```

5. 安装 calico 网络插件(解决 dns pending 问题)
   - 这个时候 coredns 的 pending 状态解除

```
sudo kubectl apply -f https://docs.projectcalico.org/v3.11/manifests/calico.yaml

```

export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at
https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root
kubeadm join 192.168.0.102:6443 --token ls3teb.3go4xl5v16o0st4h \
--discovery-token-ca-cert-hash sha256:b196ab80e32c9481fe95532f72e81250ce3d57671c9a9e83cc6478b
96b88c1e3

6. 部署一个应用，发现一直处于
   pending。这是因为我们的新建集群目前还只有一个 master 节点，而 k8s 集群默认情况下是不会把 pod 调度到
   master 节点的。我们可以解除这这种限制：

   kubectl taint nodes --all node-role.kubernetes.io/master-

7. 部署应用

- kubectl apply -f test-pod.yaml //应用信息 kind name command
- kubectl get pod //获取 pod 镜像
- kubectl logs <name> //查看镜像部署情况
- kubectl top nodes //查看内存或 cpu 使用量，通过 metrics-server 查看
