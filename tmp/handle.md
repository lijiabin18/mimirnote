#### 抓包

```shell
ip.src ==192.168.31.143 and ip.dst == 13.127.67.158 and http.request.method == "GET"

ip.src ==192.168.31.143 and ip.dst == 13.127.67.158 and http.request.method == "GET" and http.request.uri.path =="/v1/op/gifts/"

## 查看响应内容(响应体)具体操作
1. 选中请求
2. 追踪流
3. http

## 手动过滤(不推荐)
frame.len == 329 and tcp.stream eq 4

- frame.len 即为数据桢 的长度
- tcp.stream 的值 即当前桢 在tcp 层中对应的 stream.index 值

## 只查看tcp报文中发出去的字段(tcp payload)
tcp.flags.push eq 1

```

#### 线上 mysql

- run

```shell
	docker run --rm -p 10010:9080 -v=/home/ubuntu/.aws/:/root/.aws/ -v=/home/ubuntu/.google/:/.google/ -v=/home/ubuntu/.dokypay/:/.dokypay/ -v=/home/ubuntu/.cashfree/:/.cashfree/ -v=/home/ubuntu/.oceanpay/:/.oceanpay/  -v=/home/ubuntu/.tap/:/.tap/ iap.server:liuyu http --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"  --mysql_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/iap?charset=utf8mb4" --mysql_read_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com)/iap?charset=utf8mb4" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --noti_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --kafka_topic="reward_info" --mongo_addr="172.31.18.185:27017"
```

- addr

```shell
//iap
mysql_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-culfyfk1zivg.ap-south-1.rds.amazonaws.com:3306)/iap?charset=utf8mb4" --mysql_read_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com)/iap?charset=utf8mb4" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"
mysql -umaster -pmrcd12345 -h zhifu-dashang.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com iap
//user_center
mysql -umaster -pmrcd12345 -h userdb-cluster-1.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com user_center
```

#### notification

- ssh -D 9000 ubuntu@13.127.67.158
- docker build -t noti:lijiabin -f common/notification/Dockerfile.prod .
- docker build -t noti/lijiabin -f common/notification/Dockerfile.prod .
- docker run --rm -p10010:9080 noti:lijiabin http --mongo_addr=172.31.18.185:27017 --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --relationship_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"
- mongo 测试地址
- docker run --rm -p10010:9080 noti/lijiabin http --mongo_addr="172.31.20.188:27017" --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --relationship_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"

- docker run --rm -p 10010:9080 noti/lijiabin push-op -v=/home/ubuntu/.aws/:/root/.aws/ -v=/home/ubuntu/.google/:/.google/ -v=/home/ubuntu/.dokypay/:/.dokypay/ -v=/home/ubuntu/.cashfree/:/.cashfree/ -v=/home/ubuntu/.oceanpay/:/.oceanpay/ -v=/home/ubuntu/.payssion/:/.payssion/ -v=/home/ubuntu/.dlocal/:/.dlocal/ -v=/home/ubuntu/.dlocal2/:/.dlocal2/ -v=/home/ubuntu/.tap/:/.tap/ iap.server:lijiabin op --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mysql_dsn="iap:mrcd12345@tcp(172.31.20.188:3306)/iap?charset=utf8mb4" --redis_addr="172.31.20.188:6379" --lb_redis_addr="172.31.20.188:6379" --kafka_topic="reward_info" --kafka_restful_addr="172.31.20.188:9093" --kafka_host_port="172.31.20.188:9092" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --audit_service_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mongo_addr="172.31.20.188:27017" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"

#### 用户信息查询

- curl -H "Content-Type:application/json" -X POST -d '{"user_id": "17343746"}' http://usercenter-op.internal.newsdogapp.com/users/token/?app=funshare

#### iap

- 构建 image ，tag iap.server:lijiabin

```shell
docker build -t iap.server:lijiabin -f common/iap/Dockerfile.prod .

## 构建成功后的信息(contain_id and tag)
Successfully built 01369203425b
Successfully tagged iap.server:lijiabin
```

- 执行

```shell
//http
	docker run --rm -p 10010:9080 -v=/home/ubuntu/.aws/:/root/.aws/ -v=/home/ubuntu/.google/:/.google/ -v=/home/ubuntu/.dokypay/:/.dokypay/ -v=/home/ubuntu/.cashfree/:/.cashfree/ -v=/home/ubuntu/.oceanpay/:/.oceanpay/ -v=/home/ubuntu/.payssion/:/.payssion/ -v=/home/ubuntu/.dlocal/:/.dlocal/ -v=/home/ubuntu/.dlocal2/:/.dlocal2/ -v=/home/ubuntu/.tap/:/.tap/ iap.server:lijiabin http  --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"  --mysql_dsn="iap:mrcd12345@tcp(172.31.20.188:3306)/iap?charset=utf8mb4" --redis_addr="172.31.20.188:6379" --lb_redis_addr="172.31.20.188:6379" --kafka_topic="reward_info" --kafka_restful_addr="172.31.20.188:9093" --kafka_host_port="172.31.20.188:9092" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --audit_service_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mongo_addr="172.31.20.188:27017" --mongo_server_use_tag="" --mongo_replicaset_name=""

	//mysql 是线上

docker run --rm -p 10010:9080 -v=/home/ubuntu/.aws/:/root/.aws/ -v=/home/ubuntu/.google/:/.google/ -v=/home/ubuntu/.dokypay/:/.dokypay/ -v=/home/ubuntu/.cashfree/:/.cashfree/ -v=/home/ubuntu/.oceanpay/:/.oceanpay/ -v=/home/ubuntu/.payssion/:/.payssion/ -v=/home/ubuntu/.dlocal/:/.dlocal/ -v=/home/ubuntu/.dlocal2/:/.dlocal2/ -v=/home/ubuntu/.tap/:/.tap/ iap.server:lijiabin http --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mysql_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/iap?charset=utf8mb4" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mongo_addr="172.31.20.188:27017" --mongo_server_use_tag="" --mongo_replicaset_name=""
//备用域名，contain_name lijiabinhttp
docker run --rm -p 20000:9080 -v=/home/ubuntu/.aws/:/root/.aws/ -v=/home/ubuntu/.google/:/.google/ -v=/home/ubuntu/.dokypay/:/.dokypay/ -v=/home/ubuntu/.cashfree/:/.cashfree/ -v=/home/ubuntu/.oceanpay/:/.oceanpay/ -v=/home/ubuntu/.payssion/:/.payssion/ -v=/home/ubuntu/.dlocal/:/.dlocal/ -v=/home/ubuntu/.dlocal2/:/.dlocal2/ -v=/home/ubuntu/.tap/:/.tap/ iap.server:lijiabinhttp http --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mysql_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/iap?charset=utf8mb4" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mongo_addr="172.31.20.188:27017" --mongo_server_use_tag="" --mongo_replicaset_name=""


--user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"
--mysql_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/iap?charset=utf8mb4"
--chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"
--achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"

//待补充线上数据库
--redis_addr=
--kafka_topic=
--kafka_host_port=
--kafka_restful_addr=
--audit_service_addr=
--mongo_addr=

--mysql_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-culfyfk1zivg.ap-south-1.rds.amazonaws.com:3306)/iap?charset=utf8mb4" --mysql_read_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com)/iap?charset=utf8mb4" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"
	/// 线上数据库
//op
	docker run --rm -p 10010:9080 -v=/home/ubuntu/.aws/:/root/.aws/ -v=/home/ubuntu/.google/:/.google/ -v=/home/ubuntu/.dokypay/:/.dokypay/ -v=/home/ubuntu/.cashfree/:/.cashfree/ -v=/home/ubuntu/.oceanpay/:/.oceanpay/ -v=/home/ubuntu/.payssion/:/.payssion/ -v=/home/ubuntu/.dlocal/:/.dlocal/ -v=/home/ubuntu/.dlocal2/:/.dlocal2/ -v=/home/ubuntu/.tap/:/.tap/ iap.server:lijiabin op  --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"  --mysql_dsn="iap:mrcd12345@tcp(172.31.20.188:3306)/iap?charset=utf8mb4" --redis_addr="172.31.20.188:6379" --lb_redis_addr="172.31.20.188:6379" --kafka_topic="reward_info" --kafka_restful_addr="172.31.20.188:9093" --kafka_host_port="172.31.20.188:9092" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --audit_service_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mongo_addr="172.31.20.188:27017" --mongo_server_use_tag="" --mongo_replicaset_name=""
	//线上
docker run --rm -p 10010:9080 -v=/home/ubuntu/.aws/:/root/.aws/ -v=/home/ubuntu/.google/:/.google/ -v=/home/ubuntu/.dokypay/:/.dokypay/ -v=/home/ubuntu/.cashfree/:/.cashfree/ -v=/home/ubuntu/.oceanpay/:/.oceanpay/ -v=/home/ubuntu/.payssion/:/.payssion/ -v=/home/ubuntu/.dlocal/:/.dlocal/ -v=/home/ubuntu/.dlocal2/:/.dlocal2/ -v=/home/ubuntu/.tap/:/.tap/ iap.server:lijiabin op --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mysql_read_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com)/iap?charset=utf8mb4" --mysql_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/iap?charset=utf8mb4" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --mongo_addr="172.31.20.188:27017" --grpc_agency_addr="172.31.20.188:39867"  --mongo_server_use_tag="" --mongo_replicaset_name=""

http://staging13.miniviapp.com
```

- 测试机

| name                  | 地址                                                                                                         |
| --------------------- | ------------------------------------------------------------------------------------------------------------ |
| mysql_dsn             | --mysql_dsn="iap:mrcd12345@tcp(172.31.20.188:3306)/iap?charset=utf8mb4"                                      |
| redis_addr            | --redis_addr="172.31.20.188:6379" --lb_redis_addr="172.31.20.188:6379"                                       |
| lb_redis_addr         | --lb_redis_addr="172.31.20.188:6379"                                                                         |
| kafka_topic           | --kafka_topic="reward_info"                                                                                  |
| kafka_host_port       | --kafka_host_port="172.31.20.188:9092"                                                                       |
| kafka_restful_addr    | --kafka_restful_addr="172.31.20.188:9093"                                                                    |
| chatroom_grpc_addr    | --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"   |
| achi_addr             | --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"            |
| audit_service_addr    | --audit_service_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"   |
| mongo_addr            | --mongo_addr="172.31.20.188:27017"                                                                           |
| mongo_server_use_tag  | --mongo_server_use_tag=""                                                                                    |
| mongo_replicaset_name | --mongo_replicaset_name=""                                                                                   |
| user_center_backend   | --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"  |
| relationship_backend  | --relationship_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" |
| <++>                  | <++>                                                                                                         |
| <++>                  | <++>                                                                                                         |
| <++>                  | <++>                                                                                                         |
| <++>                  | <++>                                                                                                         |
| <++>                  | <++>                                                                                                         |
| <++>                  | <++>                                                                                                         |
| <++>                  | <++>                                                                                                         |
| <++>                  | <++>                                                                                                         |

- 线上数据库

| name                | addr                                                                                                                            |
| ------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| <++>                | <++>                                                                                                                            |
| mysql_dsn           | --mysql_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/iap?charset=utf8mb4"    |
| mysql_read_dsn      | --mysql_read_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com)/iap?charset=utf8mb4" |
| user_center_backend | --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"                     |
| chatroom_grpc_addr  | --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"                      |
| achi_addr           | --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"                               |
| noti_backend        | --noti_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"                            |
| kafka_topic         | --kafka_topic="reward_info" --mongo_addr="172.31.18.185:27017"                                                                  |
| <++>                | <++>                                                                                                                            |
| <++>                | <++>                                                                                                                            |

#### usercenter

> build

```
docker build -t lijiabin/user -f Dockerfile/usercenter.Dockerfile .

```

> run

```
docker run -it --network=mrcd lijiabin/user op --env=test
docker run -it --network=mrcd lijiabin/user grpc --env=test
docker run -it --network=mrcd -p 10010:80 lijiabin/user op --env=test
docker run -it --network=mrcd -p 10010:80 lijiabin/user op --env=test --grpc_user_addr=lijiabin_user-grpc.gateway.funshareapp.com:8090


docker run -it --network=mrcd lijiabin/user grpc --env=test --mysql_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4" --mysql_ro_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4"
```

- usercenter 使用线上数据库启动

```
docker run -it --network=mrcd -p 10010:80 lijiabin/user op --env=test --mysql_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4" --mysql_ro_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4"

docker run -it --network=mrcd -p 10010:80 lijiabin/user op --env=test
--mysql_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4"
--mysql_ro_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4"

docker run -it --network=mrcd -p 10010:80 lijiabin/user op --env=test
```

- grpc

```
- build
docker build -t lijiabin/user -f Dockerfile/usercenter.Dockerfile .

- grpc

docker-compose -f test/compose/docker-compose.yml up -d lijiabin_user-grpc

- op

docker run -it --network=mrcd -p 10010:80 lijiabin/user op --env=test --grpc_user_addr=lijiabin_user-grpc.gateway.funshareapp.com:8090 --mysql_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4" --mysql_ro_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4"

- http

docker run -it --network=mrcd -p 10010:80 lijiabin/user http --env=test --grpc_user_addr=lijiabin_user-grpc.gateway.funshareapp.com:8090 --mysql_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4" --mysql_ro_user_dsn="master:mrcd12345@tcp(userdb-cluster-1.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/user_center?charset=utf8mb4"


docker-compose -f test/compose/docker-compose.yml up chatroom_chatroom-http

docker-compose -f test/compose/docker-compose.yml up lijiabin_notification-http-server
```

<++>

> mysql

```shell


user_center:mrcd@123@tcp(mysql.gateway.funshareapp.com:3306)/user_center?charset=utf8mb4
mysql -uuser_center -pmrcd@123 user_center
```

#### chatroom_chatroom

```
docker run -it --rm -p10010:9080 -v=/home/ubuntu/.aws/:/root/.aws/ chatroom:lijiabin http \
--mongo_addr=172.31.18.185:27017 \
--user_center_backend=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--warm_redis=togo-redis.swlldj.ng.0001.aps1.cache.amazonaws.com:6379 \
--im_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--nats_backend=nats://ad84167c5f79611e8adf9026b3b9884f-a3cd0ab174ac96e9.elb.ap-south-1.amazonaws.com:80 \
--gowell_port=9017 \
--medal_service_grpc_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--iap_backend_addr=172.31.20.188:3333 \
--op_mongo=172.31.20.218:27017 \
--audit_service_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--grpc_port=9012 \
--k8sESUrl=http://172.31.28.5:9200 \
--relationship_backend=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--noti_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--game_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--assignment_backend_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--rec_addr=172.31.27.51:53189
<++>

im_addr
--gowell_port=9017 \

docker run -it --rm -p10010:9080 -v=/home/ubuntu/.aws/:/root/.aws/ chatroom:lijiabin http \
--ws_port=80 \
--nats_topic=websocket \
--rec_addr=172.31.27.51:53189 \
--mongo_addr=172.31.18.185:27017 \
--noti_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--audit_service_addr=test_audit-grpc.gateway.funshareapp.com:8090 \
--assignment_backend_addr=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--iap_backend_addr=test_iap-grpc.gateway.funshareapp.com:8090 \
--relationship_backend=ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090 \
--game_addr=test_game-grpc-server.gateway.funshareapp.com:8090 \
--user_center_backend=test_user-grpc.gateway.funshareapp.com:8090 \
--kafka_restful_addr=test.gateway.funshareapp.com:9093 \
--k8sESUrl=http://test.gateway.funshareapp.com:9200 \
--op_mongo=172.31.20.218:27017 \
--warm_redis=test.gateway.funshareapp.com:6379 \
--warm_ro_redis=test.gateway.funshareapp.com:6379 \
--kafka_addr=test.gateway.funshareapp.com:9092 \
--nats_backend=test.gateway.funshareapp.com:4222 \
--sqs_addr=http://test.gateway.funshareapp.com:9324
```

<++>
<++>

#### 小结

- httputils.GetStringParamOrDefault(r, "git_id", "") 这个是从 url 参数里取的
- mux.Vars(r)["gift_id"] 是 url 路径里取的
- 指针

#### 价格计算

1. goodsInfoIdPrices map 的 key 为 goods_info_id,value 为 goods_info_price 实体类。 goods_info_price 价格

```
- iap
//http
docker run --rm -p 10010:9080 -v=/home/ubuntu/.aws/:/root/.aws/ -v=/home/ubuntu/.google/:/.google/ -v=/home/ubuntu/.dokypay/:/.dokypay/ -v=/home/ubuntu/.cashfree/:/.cashfree/ -v=/home/ubuntu/.oceanpay/:/.oceanpay/  -v=/home/ubuntu/.tap/:/.tap/ iap.server:lijiabin http --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"  --mysql_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/iap?charset=utf8mb4" --mysql_read_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com)/iap?charset=utf8mb4" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --noti_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --topic_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --chromedp_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"  --kafka_topic="reward_info" --mongo_addr="172.31.18.185:27017"
//op
docker run --rm -p 10010:9080 -v=/home/ubuntu/.aws/:/root/.aws/ -v=/home/ubuntu/.google/:/.google/ -v=/home/ubuntu/.dokypay/:/.dokypay/ -v=/home/ubuntu/.cashfree/:/.cashfree/ -v=/home/ubuntu/.oceanpay/:/.oceanpay/  -v=/home/ubuntu/.tap/:/.tap/ iap.server:lijiabin op --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"  --mysql_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-culfyfk0zivg.ap-south-1.rds.amazonaws.com:3306)/iap?charset=utf8mb4" --mysql_read_dsn="master:mrcd12345@tcp(zhifu-dashang.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com)/iap?charset=utf8mb4" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --noti_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --topic_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --chromedp_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"  --kafka_topic="reward_info" --mongo_addr="172.31.18.185:27017"
<++>

curl -H "Content-Type:application/json" -X POST -d '{"user_id": "42430217"}' http://usercenter-op.internal.newsdogapp.com/users/token/?app=funshare
```

```
- 操作记录
curl "op.audit.svid.in/v1/log/?app=funshare&resource_id=44937128&type=usercenter_phone_update"
<++>
```

mysql -umaster -pmrcd12345 -h zhifu-dashang.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com -D iap -e 'select id,user_id,type,coin,created_at from transaction_record where user_id=31560915 and app="funshare" order by id desc limit 100;' > /tmp/liuyu/transaction.txt

mysql -umaster -pmrcd12345 -h zhifu-dashang.cluster-ro-culfyfk0zivg.ap-south-1.rds.amazonaws.com -D iap -e 'select distinct user_id,goods_id,expire_days from user_tool where expire_days >0 and usage_at < "2021-02-24 02:12:27" and usage_at > "2021-02-21 00:00:00" ' > /data/code/lijiabin/db.xls

scp ubuntu@13.127.67.158:/home/ubuntu/.cashfree/reward.csv /Users/liuyu/Desktop/reward.csv
scp ubuntu@13.127.67.158:/data/code/lijiabin/tool/elsticsearch/allType ./allType
scp ubuntu@13.127.67.158:/data/code/lijiabin/src/mimir_thl.tar.gz ./

scp ~/.config/ranger.tar.xz ubuntu@13.127.67.158:/home/ubuntu/.config/
<++>

<++>

#### chatroom

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
<++>

```
docker run --rm -p 10010:9080 -v=/home/ubuntu/.aws/:/root/.aws/ -v=/home/ubuntu/.google/:/.google/ -v=/home/ubuntu/.dokypay/:/.dokypay/ -v=/home/ubuntu/.cashfree/:/.cashfree/ -v=/home/ubuntu/.oceanpay/:/.oceanpay/ -v=/home/ubuntu/.payssion/:/.payssion/ -v=/home/ubuntu/.dlocal/:/.dlocal/ -v=/home/ubuntu/.dlocal2/:/.dlocal2/ -v=/home/ubuntu/.tap/:/.tap/ iap.server:lijiabin op --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"  --mysql_dsn="iap:mrcd12345@tcp(172.31.20.188:3306)/iap?charset=utf8mb4" --redis_addr="172.31.20.188:6379" --lb_redis_addr="172.31.20.188:6379" --kafka_topic="reward_info" --kafka_restful_addr="172.31.20.188:9093" --kafka_host_port="172.31.20.188:9092" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --audit_service_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --topic_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --chromedp_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --noti_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --grpc_agency_addr="172.31.20.188:54761" --mongo_addr="172.31.20.188:27017" --mongo_server_use_tag="" --mongo_replicaset_name=""
<++>
```

#### 更新 transfer_whitelist 数据 sql

```sql
update

```

```
docker run --rm -p 10086:9080 -v=/home/ubuntu/.aws/:/root/.aws/ -v=/home/ubuntu/.google/:/.google/ -v=/home/ubuntu/.dokypay/:/.dokypay/ -v=/home/ubuntu/.cashfree/:/.cashfree/ -v=/home/ubuntu/.oceanpay/:/.oceanpay/ -v=/home/ubuntu/.payssion/:/.payssion/ -v=/home/ubuntu/.dlocal/:/.dlocal/ -v=/home/ubuntu/.dlocal2/:/.dlocal2/ -v=/home/ubuntu/.tap/:/.tap/ iap.server:liuyu op --user_center_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090"  --mysql_dsn="iap:mrcd12345@tcp(172.31.20.188:3306)/iap?charset=utf8mb4" --redis_addr="172.31.20.188:6379" --lb_redis_addr="172.31.20.188:6379" --kafka_topic="reward_info" --kafka_restful_addr="172.31.20.188:9093" --kafka_host_port="172.31.20.188:9092" --chatroom_grpc_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --achi_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --audit_service_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --topic_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --chromedp_addr="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --noti_backend="ae2d427ccd07411e893fc02c564ef480-acb625e6c210e96b.elb.ap-south-1.amazonaws.com:8090" --grpc_agency_addr="172.31.20.188:54761" --mongo_addr="172.31.20.188:27017" --mongo_server_use_tag="" --mongo_replicaset_name=""
<++>
```

<++>
