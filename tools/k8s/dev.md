#### k8s 服务部署

1. jenkins build 镜像，获取 id
2. 查找 container name，kubectl describe deploy <POD>
3. 更新，`kubectl set image deployments/iap-http-server <CONTAINER_NAME> -n <NAMESPACE> --record`
4. 检查服务状态，kubectl describe pod chatroom 或 kubectl get pod -l app=chatroom

#### 服务示例

- iap_http
  kubectl set image deployments/iap-http-server server=492666533052.dkr.ecr.ap-south-1.amazonaws.com/iap:git50590d0b7738ebc02d8c868150cee37a80a89166 -n common --record

- iap_grpc
  kubectl set image deployments/iap-grpc-server server=492666533052.dkr.ecr.ap-south-1.amazonaws.com/iap:git50590d0b7738ebc02d8c868150cee37a80a89166 -n common --record

- iap_op
  kubectl set image deployments/iap-op-http-server server=492666533052.dkr.ecr.ap-south-1.amazonaws.com/iap:git2ef11f75eface9f033315414d3f48fddc3d29ebd -n common --record
  iap-op-http-server
  kubectl rollout undo deployment/iap-op-http-server

- chatroom-http
  kubectl set image deployments/chatroom-http chatroom=492666533052.dkr.ecr.ap-south-1.amazonaws.com/chatroom:git50590d0b7738ebc02d8c868150cee37a80a89166 -n default --record

- chatroom-op
  kubectl set image deployments/chatroom-op chatroom-op=492666533052.dkr.ecr.ap-south-1.amazonaws.com/chatroom:git50590d0b7738ebc02d8c868150cee37a80a89166 -n default --record

- chatroom-reward-rocket
  kubectl set image deployments/chatroom-reward-rocket server=492666533052.dkr.ecr.ap-south-1.amazonaws.com/chatroom:git50590d0b7738ebc02d8c868150cee37a80a89166 -n default --record

- noti-op-http
  kubectl describe deployments/notification-op-http-server -n common
  kubectl set image deployments/notification-op-http-server server=492666533052.dkr.ecr.ap-south-1.amazonaws.com/notification:gitba6690cf2110ff8f6807615666bfe677c0719078 -n common --record
  492666533052.dkr.ecr.ap-south-1.amazonaws.com/notification:gitd8bf4879a228de8ed4676ffa9b50752c5fd8a14a
  kubectl rollout undo deployment/notification-op-http-server
- 回滚
