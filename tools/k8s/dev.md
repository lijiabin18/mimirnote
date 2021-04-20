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

- chatroom-http
  kubectl set image deployments/chatroom-http chatroom=492666533052.dkr.ecr.ap-south-1.amazonaws.com/chatroom:git50590d0b7738ebc02d8c868150cee37a80a89166 -n default --record

- chatroom-op
  kubectl set image deployments/chatroom-op chatroom-op=492666533052.dkr.ecr.ap-south-1.amazonaws.com/chatroom:git50590d0b7738ebc02d8c868150cee37a80a89166 -n default --record

- chatroom-reward-rocket
  kubectl set image deployments/chatroom-reward-rocket server=492666533052.dkr.ecr.ap-south-1.amazonaws.com/chatroom:git50590d0b7738ebc02d8c868150cee37a80a89166 -n default --record
