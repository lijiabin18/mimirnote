#### iap.gift

> model

- Gift
- GiftList
- CustomizedGfit(ID;GiftID;RoomID;App)
- GiftTag(Lang;App;OpUser;UpdatedAt)
- TagCountry(ID;GiftID;TagId;APP;Country)
- GiftPrice
- GiftOrder(GiftId;Order;Scene;Country)
- OrderScene(Id;Name;App;CreateTime)
- GiftStatistics
- PropertyCirculation
- VipCustomGift(ModelGiftID;UserID;Level;App)
- GiftAndVipCustomGift(*Gift;*VipCustomGift)

> dao

- 对数据库 crud，接收 controller 传来的 session 进行操作
  - custom
  - price
  - dao

> dto

- 当 model 层的基本 struct 不满足使用，进行封装，添加格外属性

- 供 control(service)层调用的 `struct`

  - GiftDto(*Gift;*GiftPrice;\*[]GiftPrice)
  - GiftDtos type slice，唯一一个 method GetGifts():**[]\*model.Gift**

- 供 control(service)层调用的 `method`

  - GiftDto -> GetPrice
  - GiftDto -> FillCountryPrice
  - GiftDto -> Price

- 供 control(service)层调用的 `func`

  - NewGiftDto
  - NewGiftDtoByCountryPrice

- dao 与 control 中间层

> control

- 核心：

  - 初始化创建，全局唯一 service 层，供 gift 的 op 层服务
  - 接收 XxxPayload(已经接收 request) 与 appName ，返回调用者 model 实体

  ```golang
  func init() {
  GiftService = NewGiftControl()
  }

  type GiftControl struct {
  }

  func NewGiftControl() *GiftControl {
  	return &GiftControl{}
  }

  var GiftService *GiftControl
  ```

- struct
  - GiftControl，空 struct，唯一 service，内部封装 method，对 handler 提供 method
  - GiftPayload，供 handler 调用，进行 reques 与 model 封装
  - UpdateGiftPayload
  - Discount

> op

- 注册 handler，为每个不同的业务逻辑提供进行不同的控制

```golang
type GiftOpHandler struct {
	control *control.GiftControl
}

func NewGiftOpHandler() *GiftOpHandler {
	return &GiftOpHandler{
		control: control.NewGiftControl(),
	}
}
```

> handler

- 核心：AddRoutes

```golang
type GiftHandler struct {
	control *control.GiftControl
}

func NewGiftHandler() *GiftHandler {
	return &GiftHandler{
		control: control.NewGiftControl(),
	}
}

```

#### group 组分页

```
GET /v1/op/leaderboard/activity_goups/
**参数**
- app:""
- page:1
- page_num:5

**返回**

{
    "data": [
        {
            "id": "6037430b8ae4e100011a4c67",
            "name": "测试榜单增加语言",
            "app": "funshare"
        },
        {
            "id": "601916f0446be10001a61de2",
            "name": "test_自动发奖-copy2",
            "app": "funshare"
        },
        {
            "id": "600501bd6add2f0001fbaa1e",
            "name": "test_push_target",
            "app": "funshare"
        },
        {
            "id": "5ffd8c5438819d0001291d8c",
            "name": "test-xx",
            "app": "funshare"
        },
        {
            "id": "5ffc075538819d0001291ab5",
            "name": "test_自动发奖",
            "app": "funshare"
        }
    ]
}

```

#### 图片预览

```
URL /v1/api/s3file/
Method POST

params:
	- "file": //上传的具体文件
	- "bucket":"img.newsdog.today"
	- "path"://传当前日期
	- "content-type"://文件类型，若是预览图片，必须指定为 image/jpeg
	,image/png,image/gif 之一

resp:
{
    "s3path": "s3://img.newsdog.today/test/1615280928491546966.png",
    "url": "http://static2.miniviapp.com/test/1615280928491546966.png"
}
```

#### 礼物模块分区改版

> 创建礼物

- 改动：前端新增区域按钮开关，按区域划分 area 表示具体区域名；若不是区域划分，则 area 为空

```
URL /v1/op/gifts/
METHOD POST
params:
**post body:**
 - name
 - "price": 30,	// 折后价，即现价；当不打折时和 original_price 字段值相同
 - "original_price": 200, // 原价
 - "discount_start_time": 232132312,  // 折扣开始时间
 - "discount_end_time": 123213123123. // 折扣结束时间
 - currency // 货币类型 （金币coin, 水晶crystal）
 - icon
 - is_vol:false;默认为false，不是声音礼物，管理员设置为true时是声音礼物
 - order_scene // 排序场景 可选值: default| homepage |im .  不传默认为 “default”,
 - effects_icon
 - online_time
 - offline_time
 - tag_id
 - tag_online_time
 - tag_offline_time
 - scope ("global" or "customized")
 - room_ids // []string
 - order  // (旧)该礼物在特定国家的排序列表， map 结构  e.g. {"印度":1, "沙特": 2, "朝鲜": 3}
	- "area": //togo按大区划分为india，indonesia，other，middle_east
 - tag_country // tag 在特定国家的显示
 - "prices_info": [
        {
            "country": "india",
            "price": 30,	// 折后价，即现价；当不打折时和 original_price 字段值相同
            "original_price": 200, // 原价
            "discount_start_time": 232132312,  // 折扣开始时间
            "discount_end_time": 123213123123. // 折扣结束时间
        }
    ]

> 更新礼物

- 问题：
	- 大区划分，在gift_order表中，IN 与 ID与以前是一样的，所以为了区分，按大区划分，
	country字段的india 与 indonesia字段存原始国家名
	- 价格更新，以前只更新，india 与 middle_east，下个需求是否需要为其他大区也匹配对应价格更新操作
- 改动：前端新增区域按钮开关，按区域划分。注意，按区域划分和以前更新逻辑是分开的，按区域划分，
只能有目前支持的几个区域，不能是以前的多国家逻辑

> 查找所有礼物

- 沿用的以前的country参数，若是togo大区查找，则area字段传india，indonesia，other，middle_east
-
--------
前端页面调整
- 创建礼物时，country 选择 IN、ID、middle_east、other
- 待确定，最下分的是五个大区，添加礼物价格时，是否需要再添加大区，目前只有default-price、india-price、middle_east-price
```

```
> 调整接口
获取礼物	GET /v1/op/gifts/

					area中的参数，other → default

创建礼物	POST /v1/op/gifts/

					area中的参数，other → default

更新礼物  PUT /v1/op/gifts/{id}

					area中的参数，other → default

排序接口  POST /v1/op/gifts/sort/

					area中的参数，other → default

```

#### 礼物模块分区改版

> 调整接口

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

#### 福利

1. 4 月上班
2. 房补，直径 3 公里内。500
3. 餐补，无
4. 试用期，6 个月
5. 团队，后端岗位
6.

#### 负载的方案

> lvs.DR + keeypalive

> nginx

- 多进程，每个进程一个线程

#### 币商

- 背景：

  - 以前添加用户是一个 op 页面，创建订单(会指定汇率，及添加支付渠道)是另一个页面
  - 现在将链路缩短，在创建用户的时候完成两件事

- 问题：

- 币商：代理，币商要配置工会(agency)，身份勋章
- 身份：Coinser

- 转币/卖币：

  - 转币，第三方购买币商金币包
  - (新增)卖币，币商有卖币页，分销，且额度更大，app 默认使用中国银行，以前支付的时候，有些地区不支持，所以添加卖币和相关渠道，使用户能正常购币。

- 支付渠道

  - 创建用户的时候，默认配置大区支付渠道，默认是用户国家，支持 op 运营更改

  - ValidateTransferAuthority，校验币商，只有币商才能用。拓展表 新增字段，过期时间

- 支付渠道：payment channel

  - getLimitedPayments，后台给币商指定`支付渠道`
    - (线上)op 给币商转
    - (用户)购买金币包,agent(payment)

- 支付订单汇率：前端传来的值存到相关字段，不需要去考虑汇率换算。`和海龙自定义订单有交叉，`
- 支付渠道：币商向后台购买金币包时支付方式

  - 运营可以添加支付列表
  - 每个大区和国家有各自对应的默认支付渠道

- 其他：

  - 货币单位
  - 添加日期
  - 失效时间

- 自定义订单：白名单

- 币商 id 通过 op 加进去，运营会将用户标记为币商，再到自定义订单那里

- 以前。一个 op 是管理用户的；创建订单是另一个页面，支付渠道

```sql
-- 币商身份表，需要拓展(创建时间(add_date)，过期时间(expiration_date)，大区(area)，国家(country)，用户语言(lang)，身份(identity)，支付渠道(pay channel关联
agent表)，货币单位(money unit，与汇率相关))

CREATE TABLE `transfer_whitelist` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `app` char(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
alter table transfer_whitelist add `lang` string char(10) NOT NULL COMMENT '语言';
alter table transfer_whitelist add `area` string char(32) NOT NULL COMMENT '大区';
alter table transfer_whitelist add `country` string char(32) NOT NULL COMMENT '国家';
alter table transfer_whitelist add `identity` string char(10) NOT NULL COMMENT '身份标识';
alter table transfer_whitelist add `currency` string char(10) NOT NULL COMMENT '货币类型';
alter table transfer_whitelist add `rate` string decimal(6,5) NOT NULL COMMENT '货币汇率'; //请求json 中使用 float32 类型
alter table transfer_whitelist add `agency_id` string char(10) NOT NULL COMMENT '公会id';
alter table transfer_whitelist add `payment` string char(10) NOT NULL COMMENT '支付渠道';
alter table transfer_whitelist add `expire_time` datetime DEFAULT NULL COMMENT '过期时间';
alter table transfer_whitelist add `update_time` datetime DEFAULT NULL COMMENT '修改时间';
alter table transfer_whitelist add `create_time` datetime DEFAULT NULL COMMENT '创建时间';




-- 支付渠道
CREATE TABLE `agent` (
    `id` int not null AUTO_INCREMENT,
    `user_id` bigint,
    `app` char(10),
    `payment` varchar(255),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
ALTER TABLE `agent` ADD UNIQUE (`user_id`, `app`);

-- 区域_国家_支付渠道

CREATE TABLE `area_country_payment`(
	`id` int not null AUTO_INCREMENT,
	`area` char(10) not null,
	`country` char(20) not null,
	`payment` varchar(50) comment '支付渠道',
	`type` char(10) default null comment '支付类型',
	`create_time` datetime not null comment '创建时间',
	`update_time` datetime default null comment '更新时间',
	PRIMARY KEY(`id`),
)


--添加索引字段，country,type



-- model


```

### 检查用户

```
URL /v1/op/whitelist/transfer/account/check/
METHOD Get

- params
	- "app":
	- "user_id":123456789

- resp
{
	"status":"ok",
	"data":{
		"lang":"en",
		"country":"IN",
		"area":"india",
		"payment":{"NB", "PW", "UI"}//支付渠道
	}
}
```

### 添加用户

```
URL /v1/op/whitelist/transfer/account/
METHOD POST

- params
	- "app":

- body
{
	"user_id":12345678,
	"identity":"Coinseller"/"Big-user"
	"rate": //汇率，1 Coin= 多少金额
	"currency": //货币单位
	"exipre_at": //有效期
	"payment": ""//支付渠道
}

- resp
	- "status":"ok"

```

### 查找某个用户

```
URL /v1/op/whitelist/transfer/account/{id}/
Method Get

- params
	- "app":

- resp
{
	"status":"ok",
	"data":{
		"user_id":12345678,
		"identity":"Coinseller"/"Big-user"
		"area":
		"country":
		"lang":
		"rate": //汇率，1 Coin= 多少金额
		"currency": //货币单位
		"exipre_at": //有效期
		"payment": ""//支付渠道
	}
}
```

### 用户列表

```
URL /v1/op/whitelist/transfer/accounts/
Method Get

- params
	- "app":
	- "lang":
	- "area":
	- "country":
	- "identity": //两种身份coin_seller或者big_user

- resp
{
	"status":"ok"
	"data":[
		{
		"user_id":12345678,
		"identity":"Coinseller"/"Big-user"
		"area":
		"country":
		"lang":
		"rate": //汇率，1 Coin= 多少金额
		"currency": //货币单位
		"exipre_at": //有效期
		"payment": ""//支付渠道
		},
		...
	]
}

```

### 更新用户

```
URL /v1/op/whitelist/transfer/account/{id}/
METHOD PUT

- params
	- "app":

- body
{
	"identity":"Coinseller"/"Big-user"
	"rate": //汇率，1 Coin= 多少金额
	"currency": //货币单位
	"exipre_at": //有效期
	"payment": ""//支付渠道
}

- resp
	- "status":"ok"
```

### 删除用户

```
URL /v1/op/whitelist/transfer/account/{id}/
Method Delete

- params
	- "app":

- resp
{
	"status":"ok",
}

```

#### 配置支付渠道

```
URL /v1/op/transfer/payment/
METHOD POST

- params:
	- "app":

- body:
	- "area":
	- "country":
	- "payment":
	- "type": "" //为default 时，当前支付渠道为国家默认支付渠道

- resp:
	- "status":"ok"
```

#### 更新支付渠道

```
URL /v1/op/transfer/payment/{id}/
METHOD PUT

- params:
	- "app":

- body
	- "area":
	- "country":
	- "payment":
	- "type":

- resp:
	- "status":"ok"
```

#### 删除支付渠道

```
URL /v1/op/transfer/payment/{id}/
METHOD DELETE

- params:
	- "app":

- resp:
	- "status":"oK"

```

#### 支付渠道列表

```
URL /v1/op/transfer/payments/
METHOD GET

- params:
	- "app":
	- "area":
	- "country":
	- "type":""//值为default 时，返回默认支付渠道

- resp:
{
	"data":[
	{
		"id":
		"area":
		"country":
		"payment":
		"type":
	},

	]
}

```

#### 获取国家对应支付渠道配置列表

```
URL /v1/op/transfer/payment_config/
METHOD GET

- params:
	- "app":
	- "area":
	- "country":

- resp:
{
	"default_payment":["NB","PW"]
	"normal_payment":[]
}

```

<++>
