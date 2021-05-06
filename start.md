### 小结

- 耐心，耐心，耐心，多想，多用 google
- 向上管理。工作计划，学习内容不明确时，及时沟通，领导/负责人会给指点
- 合理安排工作，在保证基本工作任务的情况下，下班后补自己的不足，需要花时间知识总结，建立自己知识体系，因为工作中大多时候都是浅尝即止
- 需求不明确时，即时沟通，解决问题后再写代码
- 在看新的服务代码时，先了解目录结构，读别人的代码，模仿写。go 基础方面，go 语言推荐看下 'go 语言圣经'，'go 高级编程'两本书，拿不定地方自己写 demo 测
- 用到的数据库要的基本增删改查要掌握
- 环境问题较费时的，多请教，课下多了解下这方面的内容
- 工具熟练使用，ide，git，postman，wireshark

#### 常用库地址

| 服务                 | 命令                                                                                                | 备注 |
| -------------------- | --------------------------------------------------------------------------------------------------- | ---- |
| user_center 测试库   | mysql -uuser_center -pmrcd@123 user_center                                                          | <++> |
| iap 测试库           | mysql -uiap -pmrcd12345                                                                             | <++> |
| mongo 测试库         | mongo 172.31.20.188                                                                                 | <++> |
| redis 测试库         | redis-cli                                                                                           | <++> |
| user_center 线上读库 | mysql -umaster -pmrcd12345 -h userdb-cluster-1.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com | <++> |
| iap 线上库           | mysql -umaster -pmrcd12346 -h zhifu-dashang.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com    | <++> |
| mongo 线上库         | mongo 172.31.23.65:27017 && db.getMongo().setSlaveOk()                                              | <++> |
| redis_push_ro_addr   | redis-cli -h push-redis-ro.swlldj.ng.0001.aps1.cache.amazonaws.com -p 6379                          | <++> |
| redis_togo_ro_addr   | redis-cli -h togo-redis-ro.swlldj.ng.0001.aps1.cache.amazonaws.com -p 6379                          | <++> |
| redis_user_ro_addr   | redis-cli -h togo-redis-1-ro.swlldj.ng.0001.aps1.cache.amazonaws.com -p 6379                        | <++> |

#### 服务相关

- jenkins: https://jenkins.internal.newsdogapp.com/job/iap.server/

  - 本地建立 ssh 隧道，打开 9000 端口
  - 浏览器使用`SwitchOmega` 配置本地代理端口 9000

- [kibana](http://kibana.internal.newsdogapp.com:5601/app/kibana#/discover)

  - 本地建立 ssh 隧道，打开 9000 端口
  - 浏览器使用`SwitchOmega` 配置本地代理端口 9000

- [grafana](https://grafana.newsdogapp.com/d/C7JM8htik/golang-services?refresh=1m&orgId=1&var-namespace=All&var-service=notification-kafka&var-interval=1m)

  - 用户：公司邮箱

- [sentry](http://13.232.143.126:9000/auth/login/sentry/)
  - 用户名：个人公司邮箱，需要注册后才能登录

#### 测试环境生成自己的测试用户信息(token)

- 直接用线上的 token 访问测试地址，如果测试环境没有对应的用户，会自动创建同一个 id 的用户

- 查看自己的 id,did

```
//连线上读库，获取查自己用户信息
mysql -umaster -pmrcd12345 -h userdb-cluster-1.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com

//获取自己的did
select * from user_devices where user_id =<id>

//测试机生成自己的token
curl -H "Content-Type:application/json" -X POST -d '{"user_id": "42430217"}' http://usercenter-op.internal.newsdogapp.com/users/token/?app=funshare

//随便调用一个测试环境的用户中心接口，比如信息查询。将dud、aid、did和header中的token替换为自己的
curl -H 'Authorization: HIN eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X3R5cGUiOiJzbXMiLCJleHAiOjAsInVzZXJfaWQiOiIyMzY1NTQzMCIsInVzZXJuYW1lIjoiMTIwOSJ9.3iQfmx0IaOS484hIWXKsKy0RcB33gXcPOQIYav61pRU' -H 'Host: user.funshareapp.com' -H 'User-Agent: okhttp/4.6.0' --compressed 'http://user.funshareapp.com/v1/users/23655430/?os=android&timezone=GMT%2B08%3A00&channel=GooglePlay&ui_lang=en&imsi=&ntype=WIFI&pkg=com.fun.share&version=2.7.7&mac=ZYRYzR9NSH%2BRAEF4ig0fgdkM8x%2FsPACDdB63Wex3LmY%3D&vcode=16806&operator=unknown&os_v=29&country_code=IN&dud=23655430&imei=&android_id=kG%2F8KKuNPzluJ75Icr6oKA%3D%3D&lang=hi&brand=Redmi&device=Redmi+Note+8+Pro&aid=291ee42a-1303-40ca-ba0f-613605870a8e&did=291ee42a-1303-40ca-ba0f-613605870a8e'

//验证是否成功生成自己的测试用户信息
mysql -uuser_center -pmrcd@123 user_center

select user_id,name,created_at,country from funshare_user_accounts where user_id=42430217

```

##### 镜像构建

> 通知服务

- `docker build -t <tag> -f common/notification/Dockerfile.prod .` ,
- `docker run --rm -p10010:9080 <tag> http --mongo_addr=172.31.18.185:27017 --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --relationship_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"`

> iap

- `docker build -t iap.server:lijiabin -f common/iap/Dockerfile.prod .`
- 测试环境
  - docker run --rm -p 10010:9080 -v=/home/ubuntu/.aws/:/root/.aws/ -v=/home/ubuntu/.google/:/.google/ -v=/home/ubuntu/.dokypay/:/.dokypay/ -v=/home/ubuntu/.cashfree/:/.cashfree/ -v=/home/ubuntu/.oceanpay/:/.oceanpay/ -v=/home/ubuntu/.payssion/:/.payssion/ -v=/home/ubuntu/.dlocal/:/.dlocal/ -v=/home/ubuntu/.dlocal2/:/.dlocal2/ -v=/home/ubuntu/.tap/:/.tap/ iap.server:lijiabin http --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mysql_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/iap?charset=utf8mb4" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mongo_addr="172.31.20.188:27017" --mongo_server_use_tag="" --mongo_replicaset_name=""
- 线上环境
  - `docker run --rm -p 10010:9080 -v=/home/ubuntu/.aws/:/root/.aws/ -v=/home/ubuntu/.google/:/.google/ -v=/home/ubuntu/.dokypay/:/.dokypay/ -v=/home/ubuntu/.cashfree/:/.cashfree/ -v=/home/ubuntu/.oceanpay/:/.oceanpay/ -v=/home/ubuntu/.payssion/:/.payssion/ -v=/home/ubuntu/.dlocal/:/.dlocal/ -v=/home/ubuntu/.dlocal2/:/.dlocal2/ -v=/home/ubuntu/.tap/:/.tap/ iap.server:lijiabin http --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mysql_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/iap?charset=utf8mb4" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mongo_addr="172.31.20.188:27017" --mongo_server_use_tag="" --mongo_replicaset_name=""`

> usercenter

- `docker build -t ljb/user -f Dockerfile/usercenter.Dockerfile .`

- 单独起 op 服务：docker run -it --network=mrcd ljb/user op --env=test
- 单独起 grpc 服务：docker run -it --network=mrcd ljb/user grpc --env=test
- 单独起 op docker run -it --network=mrcd -p 10010:80 ljb/user op --env=test
- 给服务配置自己部署的 grpc : docker run -it --network=mrcd -p 10010:80 ljb/user op --env=test --grpc_user_addr=ljb_user-grpc.gateway.funshareapp.com:8090

> 部署 grpc

- build
  docker build -t lijiabin/user -f Dockerfile/usercenter.Dockerfile .

- grpc

  docker-compose -f test/compose/docker-compose.yml up -d lijiabin_user-grpc

- op

  docker run -it --network=mrcd -p 10010:80 lijiabin/user op --env=test --grpc_user_addr=lijiabin_user-grpc.gateway.funshareapp.com:8090 --mysql_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4" --mysql_ro_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4"

- http

  docker run -it --network=mrcd -p 10010:80 lijiabin/user http --env=test --grpc_user_addr=lijiabin_user-grpc.gateway.funshareapp.com:8090 --mysql_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4" --mysql_ro_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4"

> 聊天室

```
docker build -t chatroom:lijiabin -f Dockerfile/chatroom.Dockerfile .

docker run -it --rm -p10010:8081 -v=/home/ubuntu/.aws/:/root/.aws/ chatroom:lijiabin http \
--mongo_addr=172.31.18.185:27017 \
--user_center_backend=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--warm_redis=togo-redis.swlldj.ng.0001.aps1.cache.amazonaws.com:6379 \
--im_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--nats_backend=nats://ad84167c5f79611e8adf9026b3b9884f-a3cd0ab174ac96e9.elb.ap-south-1.amazonaws.com:80 \
--gowell_port=9017 \
--medal_service_grpc_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--iap_backend_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--op_mongo=172.31.20.218:27017 \
--audit_service_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--grpc_port=9012 \
--k8sESUrl=http://172.31.28.5:9200 \
--relationship_backend=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--noti_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--game_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--assignment_backend_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--rec_addr=172.31.27.51:53189 \
--uniform_msg_backend_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090

```

#### 其他命令

##### 压缩/解压

1. tar.gz
   压缩 tar zcvf 文件名.tar.gz 待压缩的文件名
   解压 tar zxvf 文件名.tar.gz

2. tar.xz
   压缩 tar -Jcvf xxx.tar.xz xxx
   解压 tar Jxvf xxx.tar.xz

3. 其他类型`man tar`

##### 文件上传/下载(scp)

1. 服务器下载文件

   - scp username@servername:/path/filename /tmp/local_destination

2. 服务器上传文件

   - scp /path/local_filename username@servername:/path

3. 复制整个文件夹上传到主机

   - scp -v -r diff root@192.168.1.104:/usr/local/nginx/html/webs

4. 从主机复制整个文件夹到本地
   - scp -r root@192.168.1.104:/usr/local/nginx/html/webs/diff .
5. 在两个远程主机间复制文件

   - scp root@192.168.1.104:/usr/local/nginx/html/webs/xx.txt root@192.168.1.105:/usr/local/nginx/html/webs/

6. 使用压缩传输，传输过程中被压缩，在目的主机上被解压
   - scp -vrC diff root@192.168.1.104:/usr/local/nginx/html/webs

##### log

- 测试机查看收集的日志
  - curl "op.audit.svid.in/v1/log/?app=funshare&resource_id=44937128&type=usercenter_phone_update"

##### rpc 改动

- 环境配置好，在 rpc 目录下执行`protoc --go_out=plugins=grpc:. iap.proto`

##### tmux

- 服务器环境需要把 tmux 学下，保存自己工作空间

> vi 模式

1. setw -g mode-keys vi

2. 进入 vi mode，`prefix + [` ，退出，`q`

###### 会话

- prefix
  |key|map|
  |----|----|
  |`prefix + d` |分离当前会话|
  |`prefix + s` |列出所有会话|
  |`prefix + $` |重命名当前会话|

###### 窗口

| key              | map                   |
| ---------------- | --------------------- |
| `prefix + c`     | 创建一个新窗口        |
| `prefix + p`     | 切换到上一个窗口      |
| `prefix + n`     | 切换到下一个窗口      |
| `prefix + <num>` | 切换到执行 num 的窗口 |
| `prefix + w`     | 从列表中选择窗口      |
| `prefix + ,`     | 窗口重命名            |

###### 窗格

| key          | map                          |
| ------------ | ---------------------------- |
| `prefix + %` | 划分左右窗格                 |
| `prefix + "` | 划分上下窗格                 |
| `prefix + ;` | 光标切换到上一个窗格         |
| `prefix + o` | 光标切换到下一个窗格         |
| `prefix + x` | 关闭当前窗格                 |
| `prefix + !` | 将当前窗格拆分独立           |
| `prefix + z` | 当前窗格全屏，再按一次回原样 |
| `prefix + q` | 显示窗格编号                 |
