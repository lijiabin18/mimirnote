#### k8s 服务部署

1. jenkins build 镜像，获取 id
2. 查找 container name，kubectl describe deploy <POD>
3. 更新，`kubectl set image deployments/iap-http-server <CONTAINER_NAME> -n <NAMESPACE> --record`
4. 检查服务状态，kubectl describe pod chatroom 或 kubectl get pod -l app=chatroom

#### 服务示例

eb8381123fc1db81ed01b08c91db8650319eefc7

- iap_http
  kubectl set image deployments/iap-http-server server=492666533052.dkr.ecr.ap-south-1.amazonaws.com/iap:giteb8381123fc1db81ed01b08c91db8650319eefc7 -n common --record

- iap_grpc
  kubectl set image deployments/iap-grpc-server server=492666533052.dkr.ecr.ap-south-1.amazonaws.com/iap:git -n common --record

  75895b1444cfd36f0d0a58bd7e9dd61e7c81bfd7
  777a828bf94908aead10a6f1fd6495a4a66ed542

- iap_op
  kubectl set image deployments/iap-op-http-server server=492666533052.dkr.ecr.ap-south-1.amazonaws.com/iap:git777a828bf94908aead10a6f1fd6495a4a66ed542 -n common --record
  iap-op-http-server
  kubectl rollout undo deployments/iap-op-http-server

  e1693e95c84245a11029c8c6572d80c670acf948
  765a951e4a07782575ce2bfaa16b9eba51db8ec6
  kubectl rollout undo deployment/chatroom-http

- chatroom-http
  kubectl set image deployments/chatroom-http chatroom=492666533052.dkr.ecr.ap-south-1.amazonaws.com/chatroom:git765a951e4a07782575ce2bfaa16b9eba51db8ec6 -n default --record

- chatroom-op
  kubectl set image deployments/chatroom-op chatroom-op=492666533052.dkr.ecr.ap-south-1.amazonaws.com/chatroom:git765a951e4a07782575ce2bfaa16b9eba51db8ec6 -n default --record

- chatroom-event
  kubectl set image deployments/chatroom-event server=492666533052.dkr.ecr.ap-south-1.amazonaws.com/chatroom:gite1693e95c84245a11029c8c6572d80c670acf948 -n default --record

- chatroom-reward-rocket
  kubectl set image deployments/chatroom-reward-rocket server=492666533052.dkr.ecr.ap-south-1.amazonaws.com/chatroom:git765a951e4a07782575ce2bfaa16b9eba51db8ec6 -n default --record

- noti-op-http
  kubectl describe deployments/notification-op-http-server -n common
  kubectl set image deployments/notification-op-http-server server=492666533052.dkr.ecr.ap-south-1.amazonaws.com/notification:gitba6690cf2110ff8f6807615666bfe677c0719078 -n common --record
  492666533052.dkr.ecr.ap-south-1.amazonaws.com/notification:gitd8bf4879a228de8ed4676ffa9b50752c5fd8a14a
  kubectl rollout undo deployment/notification-op-http-server
- 回滚
