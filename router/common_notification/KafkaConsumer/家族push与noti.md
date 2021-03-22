#### topic

##### chatroom_family_event

1. 家族升级给家族成员发 `noti` 注意点：
   - 根据家族 id 获取家族信息，`roomclient.BatchGetFamily`
   - 家族推送内容是存在`general_notification` 这个表里的，多熟悉 mongo
   - (客户端) 入口`common/notification/handler/handler_v4.go :getNotificationSystemAccount` ,
     需要添加事件映射`common/notification/consts/sysem_account.go :SystemAccountEventTypeMapping 中的SystemAccountFamilyAssistant`;
     配置文案 Title`common/notification/dao/dao.go :TiTitleFormatter`，titleTmpl 是
     `i18n中国际化文案，来源于在common/notification/consts/copywriting.go :FunShareEventTypeNotiTitleMappingRandom配置文案`
2. 测试事项
   - 客户端部署，在`test/compose/docker-compose` 中添加自己 build 的镜像信息部署
   - 默认测试地址绑定的都是域名(domain)，通过`host domain`
     检查是否和自己单元测试文件配置地址一致，必要时需要使用类似于` viper.SetDefault("mongo_db_name", "epimetheus")` 指定使用的 mongo 数据库名
