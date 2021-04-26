#### 需求描述

- 发 feed 后给粉丝发推送

> 开发

- topic: sharemax_useraction

- step1: 封装 notification，然后发送。核心参数如下

  - 推送跳转：feed 详情页，Url，推送时用,notification/server/kafka.go 406
    - 在 common/notification/consts/deep_link.go 中配置:yoyo://feed?id={feedId}，需要 feedId
  - 推送图片：feed 封面，imageUrl，kafka 结构体传，notification/server/kafka.go 411
  - 推送文案：定义在 i18n 下，国际化，push 时使用。notification/server/kafka.go414
    - trigger:comment_feed_to_fans (文案详情：https://docs.google.com/spreadsheets/d/1pln6eO9RKxHRCIfoWO5Uioqd4Fm-JzFpVY6-JGbDeD8/edit#gid=965011138)

- step2:
  - 降频控制 ReachFreqencyControlLimit:(app,eventType,consts,userID)
  - 获取粉丝 getNeedNotiFans:(ctx,session,app,eventType,userID)
  - 数据过滤 filterExtraData()
  - ??? nconsts.GetDeepLinkURL(app,eventType,extra)
  - 为每个粉丝纷发具体内容(过滤 robot)，写入数据库
  - push，payload 是做什么的

#### 问题

1. 推送文案不全，少了 Mr/Kn/Ml 三个语种
2. consts.AppNameTogo 与 consts.AppNameAtlas 区别
3. 评论 togo/feed/server/feed_comment.go 454 HandleFeedCommentsRequest
4. 点赞 togo/feed/server/feed_likes.go 49 updateFeedActionLog

#### OP PUSH

```
POST /v1/op/push/system_account/notifications/

"noti_title":{"hi":"","ta":"","te":"","mr":"","mn":"","ml":""}//通知标题，map结构，key 为 lang, value 为具体title
"noti_noti_img_url":{"hi":"","ta":"","te":"","mr":"","mn":"","ml":""}//通知显示图片，map结构，key 为 lang, value 为具体title

    “noti_title”: "noti title chatroom",        // 通知标题
    “noti_img_url”: "",                         // 通知显示图片
```

- 架构，服务间数据流通；沟通
- push 交给别的团队，换方法。个人方向，iap、chatroom、usercenter。
