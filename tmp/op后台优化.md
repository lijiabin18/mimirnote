#### op 后台道具 操作页优化

> 道具/勋章/尾灯/礼物 新增备注功能

- 礼物

  - 更新礼物 POST /v1/op/gifts/{id}/
  - body 新增参数 note
  - 返回 值新增 note

  - 创建礼物 POST /v1/op/gifts/
  - body 新增参数 note
  - 返回 值新增 note

- 道具

  - 更新礼物 POST /v1/op/goods/{id}/
  - body 新增参数 note

  - 创建礼物 POST /v1/op/goods/
  - body 新增参数 note

- 勋章

  - 更新勋章 PUT /v1/medal/{id}/
  - body 新增参数 note

  - 创建勋章 POST /v1/medal/
  - body 新增参数 note

- 尾灯

  - 更新尾灯 PUT /v1/taillight/{id}/
  - body 新增参数 note

  - 创建尾灯 POST /v1/taillight/
  - body 新增参数 note

> 道具新增按名称模糊搜索框

- GET /v1/op/goods/
- 新增请求参数，goods_name

> 道具/勋章 新增 All 类型

- 前端改，GET /v1/op/goods/

  - 该请求的 status 字段，新增 "all" 值

- 前端改，GET /v1/medal/
  - 前端传的 status 字段, 将原来的 0,1 改为("online","offline","all")

> 道具新增价格列

- GET /v1/op/goods/

- 新增返回值，price

> 新增价格区间筛选框

- GET /v1/op/goods/
  - 请求参数：
  - low int
  - high int
