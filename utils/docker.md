#### docker 学习

- image：镜像，只读镜像

  - docker build 生成`可执行程序` 就是 image
  - docker pull 从 dock hub 拉取 image(也可以说是从 Registry 中拉取)

- dockfile:可以看作具体`image` 的源代码，则 docker 本身就可以看作编译器

- container:容器。容器是从镜像创建的运行实例。不同 container
  间是相互隔离的，容器在启动时候创建层可写层做为最上层(写时复制,copy on write)

  - docker run <image> 运行可执行程序

- docker Registry 仓库：是集中存放镜像文件的场所
  - 区分`仓库` 与 `Registry 仓库`：仓库注册服务器上往往存放着多个仓库，
    每个仓库中又包含了多个镜像，每个镜像有不同的标签(tag)

#### 命令

> 获取镜像

- docker search <sth> 查找镜像
- docker pull <domain>/<namespace>/<repo>:tag 拉取镜像到本地

> 查看镜像信息

- docker inspect <image_id>

- eg:

  - 查找<image_id>. `docker images | grep lijiabin`
  - 查看 image 信息. `docker inspect <image_id> | less`

  | repository(image) | TAG      | IMAGEG ID    | CREATED      | SIZE   |
  | ----------------- | -------- | ------------ | ------------ | ------ |
  | lijiabin/user     | latest   | b3ffbc71b630 | 18 hours ago | 51.8MB |
  | iap.server        | lijiabin | 91c9b8486ea8 | 41 hours ago | 54.9MB |

  > 删除镜像

- docker rmi <image>:<tag>

> 创建镜像

- `docker commit <options> <container-id> <repository:tag>`
  - `-a` --author:作者信息
  - `-m` --meassage:提交消息
  - `-p` --pause=true:提交时暂停容器运行

> 迁出镜像

- docker save -o <image>.tar <image>:<tag>
  - image 可以为标签(tag)或者 ID

> 载入镜像

- docker load --input <image>.tar 或 docker load <image>.tar

> 上传镜像

- docker push <domain> / <namespace>/<repo>:<tag>

#### 容器常用操作

> 查看容器信息

- docker ps -a | grep jiabin

| container_id | image         | command                | created      | status                  |
| ------------ | ------------- | ---------------------- | ------------ | ----------------------- |
| 0d61e1fb5550 | lijiabin/user | "/go/bin/main op --e…" | 18 hours ago | Exited (0) 18 hours ago |

> 删除容器

- docker rm <container_id>

> 查看容器

- docker create -ti <name> 创建容器

  - docker start <name> 启动创建的容器

- docker run <name> 创建并启动容器，若本地有镜像则利用本地的镜像创建容器

- 原理：

  - 检查本地是否存在镜像，若不存在则从公有仓库下载
  - 利用本地镜像创建并启动一个容器
  - 分配一个文件系统，并在只读的镜像层外才挂载一层可读写层
  - 从宿主机配置的网桥接口桥接一个虚拟接口到容器中去
  - 从地址池配置一个 IP 地址给容器
  - 执行用户指定的用户程序
  - 执行完毕后容器被终止

- docker run -it ubuntu /bin/bash

  - `-i` :让容器的标准输入保持打开
  - `-t` :让 docker 分配一个伪终端，并绑定到容器的标准输入上
  - `exit` :在容器中执行该指令则退出容器，退出后`容器自动进入终止态`

- docker run --name demo -d ubuntu，docker 容器以守护态在后台运行
  - `-d` 后跟 image 的名字
  - `--name` 后跟启动的 container 的名字，若不加该参数，则以会自动生成名字

> 终止容器

- docker stop <container_id>，当容器中的应用终止时，容器会自动终止
  - 查看终止的容器：docker ps -a
  - 查看运行的容器：docker ps
  - 重新启动容器：docker start <container_id>

> 进入容器

- 启动时进入容器：`docker run -it <image> -d /bin/bash`
- 进入正在运行的容器：`docker exec -it <container_id> bash`

> 删除容器

- docker rm container_id

> 导出和导入容器

- docker export <container_id>

  - eg: `docker export test_id > test.tar`

- docker import <file>

#### docker 网络访问

1. 默认情况下，容器能访问外部网络，但是外部网络不能访问容器。

   - 使用 `docker exec -it container_id bash` 进入容器
   - 进入容器后，使用 `iptables -t nat -L -n` ，可以查看容器内部网络状态

2. 通过将容器的端口映射到主机的端口，即能使用外部访问容器
   - 当容器运行一些网络应用时，要让外部访问这些应用，`需要在创建containter时，使用-P或-p 指定端口映射` ，`-p` 即为为外部访问提供的端口号
   - eg: `docker run --name docker_mysql -p 12345:3306 -e MYSQL_PASSWORK=<passwd> -d myslq:5.6.35`
     - -p，即将容器内部的 3306 映射为 12345 供外部访问
     - -d，以守护进程启动
   - 映射到指定地址`... -p 127.0.0.1:12345:3306....`
   - 映射到指定地址的任意端口：`...-p 127.0.0.1::3306`
   - 查看内部映射端口配置：docker port image_name port

#### 容器间相互通信

1. 方法 1：通过映射宿主机的端口实现容器互联
2. 方法 2：容器的连接(link)除了端口映射外，也是另一种可以与容器中的应用进行交互的方式

- (方法 2)示例

  - 第一步，创建一个数据库容器(mysqldb)

  ```shell
  docker run --name mysqldb -p 1234:3306 -e MYSQL_ROOT_PASSWORD=<passwd> -d mysql:5.6.35
  ```

  - 第二步，创建一个 Web 容器(tomcatApp)，并与数据库建立连接，使用`--link`，tomcat 容器创建好后,对外通过 MySQL 进行通信

  ```shell
  docker run  --name tomcatApp --link mysqldb:MySQL -d tomcat
  ```
