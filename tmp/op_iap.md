#TODO

> Feed 测试: <22-01-21,li> >
> giving icon 测试: <28-01-21, yourname> >
> k8s 部署测试环境: <28-01-21, li> >
> 更新数据库: <28-01-21, yourname> >
> 更新数据库: 更新数据库 <28-01-21, yourname> >

#### skill

- 查找功能相关表，可以直接看 xorm 的 model

#### tag 删除

DELETE /v1/op/gifts/tag/{tag_id}/
删：
url: /v1/op/gifts/tag/{tag_id}/?app=
method: DELETE
resp: ok

**参数** - app:
**返回值** - 无

#### 多礼物配置角标

POST /v1/op/gifts/gifts/tag/{tag_id}

**参数**

- url_params:

  - app
  - tag_id

- post_body

  - ids: [1, 2, 3] //gift_id
  - tag_online_time//timestamp
  - tag_offline_time
  - tag_id
  - tag_country: ["IN","SA"]

**返回值** - 无

#### question

1. api.go 中有 status 不同状态的 func，我新增加了 all 类型，需要在这里添加吗
2. 页面传过来的 status 是数值还是字符串("online","offline")，已解决
3. 数据库表单新增`note` 字段，但是查询的时候没有字段显示
4. 尾灯后台访问不了
5. 数据问题，goods_kind 表的 实体类中的 price
   被删除了，数据库中该字段有一半为 0 值

#### 声音礼物

> 问题：

- 声音礼物样例图和 app 自己后台不一样
  - category 控制
- gift 的 type 字段应用场景(reward → lucky,ticket, rpc,leaderboard → cp)
- gift 的 category(vip,custom,appointment,hot,soft) 和 app 中显示的分类不一致
- CustomizeGift 是什么场景,dao.CreateCustomizedGiftForRooms(session, gift.Id, roomIDs, app)
- order 与 orderScene(name,app,create_time) 应用场景

> 需求分析：

1. 后台，添加字段 ,isVol(bool)；前台，添加 vol_gift(默认为 0)
2. 创建/更新礼物时，获取到 vol_gift 字段为 1(默认为 0)时，获取请求中的 effectiveIcon
   中的信息，写入数据库
3. 查询操作，不需要额外添加控制，直接返回结果
4. 删除操作，若礼物是声音礼物，则修改数据库 isVol 与 effectiveIcon 字段

- 疑问：

  - 需求描述中，角标与声音礼物标识共存，后台通过 isVol 字段判断是声音礼物，两种标识
    共存应该是前端调整，后台需要在角标做其他控制吗
  - 自定义异常信息

> 需求文档

- url v3/gifts/

- method GET
- 请求参数中新增字段

  > 声音礼物 op

- 改动

  - 新增 is_vol(bool)字段
  - 声音礼物角标由 vol_gift_tag 字段表示
  - 声音礼物的声音文件由 effects_icon 字段表示，声音文件格式为 .svga 格式

- 创建礼物

  - url /v1/op/gifts
  - method POST
  - post_body:
    - is_vol: bool //默认 false，为 true 时礼物为声音礼物

- 更新礼物：

  - url /v1/op/gifts/{id}
  - method POST
  - post_body:
    - is_vol: bool //默认 false，为 true 时礼物为声音礼物

- 礼物列表:
  - url /v1/op/gifts/
  - method GET
  - response:
    - 新增 is_vol，
    - 新增 vol_gift_tag//声音角标

#### feed 推送

> 问题：

1. feed 是什么
2. 触发动作，发布成功、评论、点赞，不同状态的动作是在哪里控制,监听的

3. feed 相关 http router

   - togo/main/http_route.go 注册路由 HandleFeedCreateRequest
   - feed.go 中 feed create handler
     - HandleFeedCreateRequest → feedCreateHandler.go 408 line push
     - common/notification/server/kafka.go 处理 topic

4. 参考其他推送基本流程是如下，准备工作??，push 时的方法太复杂，没看懂
   - step1: 创建结构体接收传来的内容
   - step2: 准备工作
   - step3: 定义 notification 结构体
   - step4: version、push(核心)

##### 前端相关页面

- 礼物前端其他请求需要调整显示效果，新增了备注栏

  - "getGiftsOp" GET /v1/op/gifts/
  - "getGiftsByIDOp" POST /v1/op/gifts/gifts/
  - "updateGiftsByTag" POST /v1/op/gifts/gifts/tag/
  - "getGiftOp" GET /v1/op/gifts/{id}

- 道具前端调整，新增备注功能
  - "getGoods" GET /v1/op/goods/

#### 赠送道具优化(自定义激活时间/有效期)

- 接口文档

```
url /v1/op/goods/{goods_id:[0-9]+}/circulation/
method: post
req_params: goods_id
path_params: app
post_body:

{
    "user_ids":[39786842],
    "goods_type":"mount",
    "is_activated":false,//不自动激活，到指定激活时间后激活，必须填false
    "expire_days":10,//道具有效期
    "active_days":15,//道具激活时间，激活时间为0时，则根据有效期duration来激活，expire_days无效
}

response:
{
	"status": ok
}
//查询 sql
select property_id,receiver_id,property_type,property_count from property_circulation where property_id=616 and receiver_id = 39786842;
select property_id,receiver_id,property_type,property_count from property_circulation where receiver_id = 39786842;

```

- 线上数据库更新

```sql
alter table user_tool add `expire_days` int(10) DEFAULT 0;

```

- 待测试接口

39786842
29830391

```
-- 送礼接口(http)
POST /v1/reward/

POST /v1/reward/inprivate/


-- 自定义道具接口(op)

POST /v1/goods/{goods_id:[0-9]+}/circulation/

```

<++>

#### [礼物角标优化](https://github.com/tmrwh/NewsDog/issues/7515)

> 变更，新增加一张表 lang_info

```sql
CREATE TABLE `tag_info`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `tag_id` bigint(20) NOT NULL ,
    `lang` char(16) NOT NULL,
    `hint` varchar 128 default "",
    `hint_url` varchar(128) default "",
    PRIMARY key (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

```

> 需求描述

- 背景：以前的场景是，默认一个 tag 所有语言通用，即 tip text 和 tip link 是一致 的
- 新需求：同一个角标，根据语言不同，每个`tip text` 和 `tip link` 不同

> 方案

- gift_tag 表，添加 extra 字段，键值对形式存储不同语言的 `hint` 与 `hint url`

```
extra:
	[]map[string]*dto.langTag:
		m1:
			hint:
			hinturl:

		m2:
			hint:
			hinturl:

```

> 涉及到的改动

- tag 增删改查
- gift 创建时配置 tag 的逻辑

> 获取所有角标

```
url:/v1/op/gifts/tag/list/?lang=&app=
method:GET
resp:
{
"status":"ok"
"data":[
{
"id":int64,
"title":,
"icon":,
"hint": "",
"hint_url": "",
"lang":,
"app":,
}
]
}

```

> 创建 tag

- 默认创建为 en 类型

```
url: /v1/op/gifts/tag/?lang=&app=
params: lang= app=
method: POST
post_data
{
"title":
"icon":
"op_user":
"hint": "",
"hint_url": "",
}
resp:
{
"id":int64,
"title":,
"icon":,
"lang":,
"app":,
"hint": "",
"hint_url": "",
}
```

> 删除 tag

- 代码结构稍做修改

```
url: /v1/op/gifts/tag/{tag_id}/?app=
method: DELETE
resp:

```

> 按语言获取 tag 列表

- 改动，传语言参数 lang，默认展示为 en 类型 tag

```
url:/v1/op/gifts/tag/list/?lang=&app=
method:GET
req_params:
	- lang (hi,ta,te,mr,ar,en,id,kn,ml 9 种类型中的一种，若不填则默认为en)
	- app
resp:
{
"status":"ok"
"data":[
{
"id":int64,
"title":,
"icon":,
"hint": "",
"hint_url": "",
"lang":,
"app":,
}
]
}
```

> 获取单个 tag 的详情(当前改动需和前端/需求沟通)

- 改动部分，现在返回当前 tag 的所有语言类型的信息

```
url:/v1/op/gifts/tag/{tag_id}/?lang=&app=
method: GET
req_params:
	- app
	- 不传lang值，直接返回当前tag的所有语言的详情

resp: 需要改
{
"id":int64,
"title":,
"icon":,
"lang":,
"app":,
"hint": "",
"hint_url": "",
}

```

#### [【OP】活动模板 新增 团队榜单 功能组件](https://github.com/tmrwh/NewsDog/issues/7626)

- 问题：

  - 团队排行是怎么弄的。是数据杨松那边去处理吗，http
    端的`common/iap/leaderboard/handler/leaderboard.go :66 rankActivity`
    ，获取 rank 排名时是在 reids 是中取的，这块是数据处理的吗 - getRank()，key 与 type 关联

  - datasource 应该是我和前端确定？
  - extra
    字段可以存我想存的数据吗，团队榜单仅是针对礼物的吗，那我可以参考礼物榜单吗

- 需求文档
- 后台活动接口变化(v1/op/activitys 增删改查)

  - LBActivityType 新增 `team_send` 与`team_receive`
  - extra map 新 1 个字段(teamIcon)，存团队 icon

- 需求会议

```
- 一个队伍就是一场活动，
- mongo 创建队伍表(集合),队伍id
	关联到leaderboard，用数组teamList存。运营传用户id / 语言

- team 排名榜，team 成员(国家分、语言分、手动传用户ID)？
- 成员为国家时，需要国家排行榜
	- icon ：一个team
	- name ：一个
	- 组队：？，按国家来不同的team会国家相同吗
- 子榜 ？
- 语言pk

<++>
```

<++>
