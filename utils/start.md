#### 常用服务地址

| 服务               | 命令                                       | 备注 |
| ------------------ | ------------------------------------------ | ---- |
| user_center 测试库 | mysql -uuser_center -pmrcd@123 user_center | <++> |
| iap 测试库         | mysql -uiap -pmrcd12345                    | <++> |
| <++>               | <++>                                       | <++> |
| <++>               | <++>                                       | <++> |
| <++>               | <++>                                       | <++> |
| <++>               | <++>                                       | <++> |
| <++>               | <++>                                       | <++> |

#### 部署相关

> 常用 docker 命令

-

> 个人设备信息查询

- 测试机执行：`curl -H "Content-Type:application/json" -X POST -d '{"user_id": "17343746"}' http://usercenter-op.internal.newsdogapp.com/users/token/?app=funshare`

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

<++>

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

- curl "op.audit.svid.in/v1/log/?app=funshare&resource_id=44937128&type=usercenter_phone_update"

##### rpc 改动

- protoc --go_out=plugins=grpc:. iap.proto

##### 服务相关

- jenkins: https://jenkins.internal.newsdogapp.com/job/iap.server/

- kibana：http://kibana.internal.newsdogapp.com:5601/app/kibana#/discover

- grafana：https://grafana.newsdogapp.com/d/C7JM8htik/golang-services?refresh=1m&orgId=1&var-namespace=All&var-service=notification-kafka&var-interval=1m
