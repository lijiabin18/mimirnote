#### 礼物模块分区改版

- 2021-03-16,[issue](https://www.google.com/url?q=https://github.com/tmrwh/NewsDog/issues/7636&sa=D&source=editors&ust=1615881193450000&usg=AFQjCNHkwoiHPi7EXANND3ettNXYjfH9-g)

> 调整接口参数

- [获取礼物](https://github.com/tmrwh/NewsDog/wiki/iap-op-%E5%90%8E%E5%8F%B0-api#%E7%A4%BC%E7%89%A9%E5%88%97%E8%A1%A8) GET /v1/op/gifts/

  - area 中的参数，other → default

- [创建礼物](https://github.com/tmrwh/NewsDog/wiki/iap-op-%E5%90%8E%E5%8F%B0-api#%E5%88%9B%E5%BB%BA%E7%A4%BC%E7%89%A9) POST /v1/op/gifts/

  - area 中的参数，other → default

- [更新礼物](https://github.com/tmrwh/NewsDog/wiki/iap-op-%E5%90%8E%E5%8F%B0-api#%E6%9B%B4%E6%96%B0%E7%A4%BC%E7%89%A9) PUT /v1/op/gifts/{id}

  - area 中的参数，other → default

- [排序接口](https://github.com/tmrwh/NewsDog/wiki/iap-op-%E5%90%8E%E5%8F%B0-api#%E6%8E%92%E5%BA%8F%E6%8E%A5%E5%8F%A3) POST /v1/op/gifts/sort/

  - area 中的参数，other → default

- [礼物查找](https://github.com/tmrwh/NewsDog/wiki/iap-op-%E5%90%8E%E5%8F%B0-api#id%E6%89%B9%E9%87%8F%E6%90%9C%E7%B4%A2) POST /v1/op/gifts/gifts/

  - area 中的参数，other → default

- [礼物 id 查找](https://github.com/tmrwh/NewsDog/wiki/iap-op-%E5%90%8E%E5%8F%B0-api#%E5%8D%95%E4%B8%AA%E7%A4%BC%E7%89%A9%E8%AF%A6%E6%83%85) GET /v1/op/gifts/{id}

  - area 中的参数，other → default

#### 礼物大区补充

1. 创建礼物，单大区变多大区。area → []area
2. 编辑礼物，同时支持多大区编辑。id → []area
3. 查找礼物，返回多大区信息。

> 待改善

- h.control.UpdateCustomizedGift(giftId, p.RoomIDs,
  appName)，自定义礼物更新操作是否需要根据大区来

-
