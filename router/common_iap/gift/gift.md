#### 应用场景

- router 入口为 `op/gift.go`

#### 包结构

| 层           | 作用                                                | 其他                 |
| ------------ | --------------------------------------------------- | -------------------- |
| `api`        | ?                                                   | circulation/custom   |
| `consts`     | scope/type                                          | <++>                 |
| `control`    | 介与 dao 与 router 中间，解耦，同 java 中的 service | gift/price           |
| `dao`        | 数据库操作                                          | custom/dao/price     |
| `dto`        | ?                                                   | gift                 |
| `handler`    | http 服务，为客户端提供服务                         | gift/gift_v2/gift_v3 |
| `model`      | 实体类                                              | <++>                 |
| `serializer` | 序列化                                              | <++>                 |

```shell

├── api
│   ├── circulations.go
│   └── custom.go
├── consts
│   └── consts.go
├── control
│   ├── gift.go
│   └── price.go
├── dao
│   ├── custom.go
│   ├── dao.go
│   └── price.go
├── dto
│   └── gift.go
├── handler
│   ├── gift.go
│   ├── gift_v2.go
│   └── gift_v3.go
├── model
│   └── model.go
└── serializer
    └── gift.go

```
