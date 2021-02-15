#### Gift 表间的关系

- gift

| Field               | Type         | Desc                               |
| ------------------- | ------------ | ---------------------------------- |
| id                  | bigint(20)   | <++>                               |
| name                | varchar(255) | <++>                               |
| lang                | char(16)     | <++>                               |
| vip                 | int(10)      | <++>                               |
| category            | varchar(32)  | <++>                               |
| price               | bigint(20)   | 当前价格，如果是有折扣则显示折扣价 |
| is_discount         | tinyint(1)   | 是否有折扣                         |
| is_vol              | tinyint(1)   | <++>                               |
| currency            | varchar(32)  | <++>                               |
| contribution        | bigint(20)   | <++>                               |
| status              | int(11)      | <++>                               |
| app                 | char(32)     | <++>                               |
| type                | varchar(32)  | lucky/cp/ticket/ac_cp              |
| icon                | char(128)    | <++>                               |
| note                | varchar(255) | <++>                               |
| effects_icon        | char(128)    | 声音文件，svg 文件                 |
| online_time         | datetime     | <++>                               |
| offline_time        | datetime     | <++>                               |
| create_time         | datetime     | <++>                               |
| update_time         | datetime     | <++>                               |
| expire_time         | datetime     | <++>                               |
| tag_id              | bigint(20)   | <++>                               |
| tag_online_time     | datetime     | <++>                               |
| tag_offline_time    | datetime     | <++>                               |
| scope               | char(32)     | global/customized/vip_custom       |
| original_price      | bigint(20)   | 原始价格                           |
| discount_start_time | datetime     | <++>                               |
| discount_end_time   | datetime     | <++>                               |

- customized_gift，与 roomid 绑定

| Field   | Type       | Desc |
| ------- | ---------- | ---- |
| id      | bigint(20) | <++> |
| gift_id | bigint(20) | <++> |
| room_id | char(30)   | <++> |
| app     | char(32)   | <++> |

- tag_country，一个礼物对应一个 tag，多个 country

| Field   | Type       | Desc |
| ------- | ---------- | ---- |
| id      | bigint(20) | <++> |
| gift_id | bigint(20) | <++> |
| tag_id  | bigint(20) | <++> |
| app     | char(10)   | <++> |
| country | char(32)   | <++> |

- order_scenes，??? 礼物订单

| Field       | Type         | Desc |
| ----------- | ------------ | ---- |
| id          | bigint(20)   | <++> |
| name        | varchar(255) | <++> |
| app         | char(32)     | <++> |
| create_time | datetime     | <++> |

- gift_price，礼物是有折扣(IsDiscount 为 true)时，折扣信息插入该表

| Field               | Type        | Desc |
| ------------------- | ----------- | ---- |
| id                  | bigint(20)  | <++> |
| gift_id             | bigint(20)  | <++> |
| price               | int(11)     | <++> |
| original_price      | int(11)     | <++> |
| country             | varchar(32) | <++> |
| discount_start_time | datetime    | <++> |
| discount_end_time   | datetime    | <++> |
| create_time         | datetime    | <++> |
| update_time         | datetime    | <++> |
