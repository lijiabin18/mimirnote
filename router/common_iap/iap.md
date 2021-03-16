#### 作用

#### 层级

| 子模块        | 作用                  | 其他 |
| ------------- | --------------------- | ---- |
| `activity`    | 活动                  | <++> |
| `bootstrap`   | 启动类                | <++> |
| `cmd`         | cobra 命令行工具      | <++> |
| `config`      | dao/handler/models    | <++> |
| consts        | currency/country 常量 | <++> |
| `cron`        | 自动激活角标工具      | <++> |
| `gift`        | 礼物相关 handler      | <++> |
| `gift_pack`   | <++>                  | <++> |
| `i18n`        | 国际化                | <++> |
| `leaderboard` | <++>                  | <++> |
| `lottery`     | <++>                  | <++> |
| `market`      | <++>                  | <++> |
| `noble`       | <++>                  | <++> |
| `op`          | <++>                  | <++> |
| `prize`       | <++>                  | <++> |
| `reward`      | <++>                  | <++> |
| `rpc`         | <++>                  | <++> |
| `settlement`  | <++>                  | <++> |
| `subscriber`  | <++>                  | <++> |
| `test`        | <++>                  | <++> |
| `user`        | <++>                  | <++> |
| `utils`       | <++>                  | <++> |
| `wallet`      | <++>                  | <++> |
| `warehouse`   | <++>                  | <++> |
| `wish`        | <++>                  | <++> |
| <++>          | <++>                  | <++> |

```
├── activity
│   ├── consts
│   │   └── const.go
│   ├── cron
│   │   ├── agency_pk.go
│   │   ├── gift_pk.go
│   │   ├── host_pk.go
│   │   ├── magic_card.go
│   │   ├── recharge.go
│   │   └── valentine.go
│   ├── dao
│   │   ├── agency.go
│   │   ├── award_record.go
│   │   ├── displayid.go
│   │   ├── gift_package.go
│   │   ├── magic_card.go
│   │   ├── recharge_activity.go
│   │   ├── super_host.go
│   │   └── valentine.go
│   ├── handler
│   │   └── handler.go
│   ├── log
│   │   └── log.go
│   ├── model
│   │   └── model.go
│   ├── mq
│   │   └── magic_card.go
│   ├── serializer
│   │   ├── common.go
│   │   ├── magic_card.go
│   │   ├── new_year.go
│   │   └── recharge_activity.go
│   └── service
│       ├── magic_card.go
│       ├── new_year.go
│       ├── recharge_activity.go
│       └── top_user_recharge.go
├── bootstrap
│   ├── auto_active_usertool_cron.go
│   ├── basic.go
│   ├── grpc_server.go
│   ├── http_server.go
│   ├── kafka.go
│   ├── leaderboard_reward.go
│   ├── mq.go
│   ├── op_server.go
│   ├── recharge_kafka.go
│   ├── recharge_stats_kafka.go
│   ├── reconciliation.go
│   ├── refund_recovery.go
│   └── timer_billing.go
├── cmd
│   ├── auto_active_usertool_cron.go
│   ├── grpc_server.go
│   ├── http_server.go
│   ├── Jenkinsfile_http_canary
│   ├── kafka.go
│   ├── leaderboard_reward.go
│   ├── mq.go
│   ├── op.go
│   ├── recharge_kafka.go
│   ├── recharge_stats_kafka.go
│   ├── reconciliation.go
│   ├── refund_recovery.go
│   ├── root.go
│   └── timer_billing.go
├── config
│   ├── dao
│   │   └── ferris_user.go
│   ├── handler
│   │   ├── hander.go
│   │   ├── handler_v2.go
│   │   └── handler_v3.go
│   └── models
│       └── ferris_wheel_user.go
├── consts
│   ├── consts.go
│   └── country.go
├── cron
│   └── auto_active_usertool_cron.go
├── db.sql
├── Dockerfile.prod
├── gift
│   ├── api
│   │   ├── circulations.go
│   │   └── custom.go
│   ├── consts
│   │   └── consts.go
│   ├── control
│   │   ├── gift.go
│   │   └── price.go
│   ├── dao
│   │   ├── custom.go
│   │   ├── dao.go
│   │   └── price.go
│   ├── dto
│   │   └── gift.go
│   ├── handler
│   │   ├── gift.go
│   │   ├── gift_v2.go
│   │   └── gift_v3.go
│   ├── model
│   │   └── model.go
│   └── serializer
│       └── gift.go
├── gift_pack
│   ├── api
│   │   └── api.go
│   ├── cron
│   │   └── refund.go
│   ├── dao
│   │   └── dao.go
│   └── model
│       └── model.go
├── i18n
│   ├── gifts.go
│   ├── i18n.go
│   ├── transaction_title.go
│   ├── transaction_title.yaml
│   ├── transaction_type_game_coin_flow.go
│   └── transaction_type.go
├── leaderboard
│   ├── api
│   │   ├── blacklist.go
│   │   └── leaderboard.go
│   ├── consts
│   │   ├── consts.go
│   │   └── i18n.go
│   ├── couple
│   │   ├── couple.go
│   │   ├── cron.go
│   │   ├── rank.go
│   │   ├── rewards.go
│   │   └── valentine_couple.go
│   ├── cron
│   │   ├── family.go
│   │   ├── rewards.go
│   │   ├── room.go
│   │   └── user.go
│   ├── dao
│   │   ├── blacklist.go
│   │   ├── fan_contribution.go
│   │   └── rewards.go
│   ├── downloader
│   │   ├── downloader.go
│   │   └── model.go
│   ├── handler
│   │   └── leaderboard.go
│   ├── models
│   │   ├── couple.go
│   │   ├── fan_contribution.go
│   │   ├── group_push.go
│   │   ├── hide.go
│   │   ├── lang.go
│   │   ├── leaderboard_activity.go
│   │   ├── rewards.go
│   │   ├── room_blacklist.go
│   │   ├── room_leaderboard.go
│   │   ├── user_blacklist.go
│   │   ├── user_leaderboard.go
│   │   ├── version.go
│   │   └── weekly_gift.go
│   ├── newrank
│   │   ├── activity.go
│   │   ├── cron.go
│   │   ├── key.go
│   │   ├── new_year.go
│   │   ├── rank.go
│   │   └── util.go
│   ├── newserializer
│   │   ├── agency.go
│   │   ├── family.go
│   │   ├── gift.go
│   │   ├── room.go
│   │   ├── user.go
│   │   └── weekly_gift.go
│   ├── payload
│   │   └── blacklist.go
│   ├── rank
│   │   ├── area.go
│   │   ├── cron.go
│   │   ├── rank.go
│   │   ├── reward.go
│   │   ├── specific_gift.go
│   │   ├── static_leaderboard.go
│   │   └── weekly_gift.go
│   └── serializer
│       ├── blacklist.go
│       ├── leaderboard_activity.go
│       ├── room.go
│       ├── user.go
│       └── user_test.go
├── lottery
│   ├── dao
│   │   └── dao.go
│   ├── handler
│   │   └── handler.go
│   ├── log
│   │   └── logging.go
│   ├── model
│   │   └── model.go
│   └── serializers
│       └── serializers.go
├── main.go
├── market
│   ├── api
│   │   └── goods.go
│   ├── consts
│   │   └── consts.go
│   ├── dao
│   │   ├── dao.go
│   │   └── price.go
│   ├── dto
│   │   └── goods.go
│   ├── handler
│   │   ├── handler.go
│   │   ├── handler_v2.go
│   │   └── handler_v3.go
│   ├── model
│   │   └── goods_model.go
│   ├── payload
│   │   └── payload.go
│   ├── serializer
│   │   └── goods.go
│   └── service
│       ├── goods.go
│       ├── payload.go
│       └── purchase.go
├── noble
│   ├── api
│   │   └── inapp.go
│   ├── handler
│   │   └── handler.go
│   ├── kafka
│   │   └── kafka.go
│   └── model
│       └── config.go
├── op
│   ├── assets.go
│   ├── blacklist.go
│   ├── cron
│   │   ├── group_push.go
│   │   └── user_ban.go
│   ├── gift.go
│   ├── goods.go
│   ├── gp_reconciliation.go
│   ├── leaderboard.go
│   ├── noble.go
│   ├── payload
│   │   └── goods.go
│   ├── recharge.go
│   ├── serializer
│   │   ├── goods.go
│   │   └── withdraw_application.go
│   ├── transfer.go
│   ├── utils
│   │   ├── requests.go
│   │   ├── s3.go
│   │   ├── user.go
│   │   └── utils.go
│   ├── whitelist.go
│   └── withdraw.go
├── prize
│   ├── log
│   │   └── log.go
│   ├── model
│   │   ├── prize.go
│   │   └── prize_record.go
│   └── service
│       ├── pack
│       │   └── pack.go
│       └── prize.go
├── proxy
│   └── proxy.go
├── reward
│   ├── api
│   │   └── api.go
│   ├── consts
│   │   └── consts.go
│   ├── dao
│   │   ├── dao.go
│   │   └── raffle_ticket.go
│   ├── errors
│   │   └── errors.go
│   ├── handler
│   │   └── handler.go
│   ├── model
│   │   ├── giving_type.go
│   │   ├── raffle_ticket.go
│   │   └── reward.go
│   ├── response
│   │   └── user.go
│   ├── reward_service
│   │   └── reward.go
│   ├── serializer
│   │   ├── raffle_ticket.go
│   │   └── reward.go
│   ├── service
│   │   ├── cron.go
│   │   ├── giving_type.go
│   │   ├── lucky_gift.go
│   │   ├── param.go
│   │   ├── reward.go
│   │   ├── test_test.go
│   │   └── ticket_gift.go
│   └── utils
│       └── logging.go
├── rpc
│   ├── client
│   │   ├── client.go
│   │   ├── client_test.go
│   │   └── test_test.go
│   ├── errors
│   │   └── errors.go
│   ├── pb
│   │   ├── iap.pb.go
│   │   └── iap.proto
│   └── server
│       └── server.go
├── settlement
│   ├── api
│   │   ├── family.go
│   │   └── whitelist.go
│   ├── consts
│   │   └── consts.go
│   ├── cron
│   │   └── family.go
│   ├── dao
│   │   └── dao.go
│   ├── errors
│   │   └── error.go
│   ├── kafka
│   │   ├── consumer.go
│   │   └── dispatch_revenue.go
│   ├── log
│   │   └── log.go
│   ├── model
│   │   └── model.go
│   ├── payload
│   │   └── payload.go
│   └── serializer
│       └── serializer.go
├── subscriber
│   ├── kafka.go
│   └── vip.go
├── test
│   ├── models
│   │   ├── fixtures
│   │   │   ├── gift_order.yml
│   │   │   ├── gift_statistics.yml
│   │   │   ├── gift.yml
│   │   │   ├── order_scene.yml
│   │   │   ├── property_circulation.yml
│   │   │   ├── reward_info.yml
│   │   │   ├── reward_relation.yml
│   │   │   ├── sku.yml
│   │   │   ├── transaction_record.yml
│   │   │   └── user_assets.yml
│   │   ├── gift_test.go
│   │   ├── main_test.go
│   │   ├── models.go
│   │   ├── reward_test.go
│   │   ├── test_fixtures.go
│   │   ├── unit_tests.go
│   │   ├── wallet_assets_test.go
│   │   ├── wallet_sku_test.go
│   │   └── wallet_trans_test.go
│   └── test.sh
├── unit_test.Dockerfile
├── user
│   └── model
│       └── user.go
├── utils
│   ├── razer_gold.go
│   ├── time.go
│   └── util.go
├── wallet
│   ├── api
│   │   ├── escrow.go
│   │   ├── inapp.go
│   │   ├── inapp_test.go
│   │   ├── test_test.go
│   │   ├── timer_billing.go
│   │   ├── transaction_record.go
│   │   ├── transaction_record_test.go
│   │   └── transfer.go
│   ├── assets
│   │   ├── dao.go
│   │   ├── dao_test.go
│   │   └── transfer_dao.go
│   ├── assets_service
│   │   ├── asset_type
│   │   │   ├── asset_type.go
│   │   │   └── helper.go
│   │   ├── consts
│   │   │   ├── asset_log_direction.go
│   │   │   ├── asset_log_status.go
│   │   │   └── exchange_status.go
│   │   ├── dao.go
│   │   ├── dto
│   │   │   ├── asset_log.go
│   │   │   ├── deduct_info.go
│   │   │   ├── exchange.go
│   │   │   ├── read_tag.go
│   │   │   └── user_asset.go
│   │   ├── errors
│   │   │   └── errors.go
│   │   ├── internal
│   │   │   ├── config
│   │   │   │   └── exchange_pair_config.go
│   │   │   ├── dao
│   │   │   │   ├── exchange.go
│   │   │   │   ├── exchange.sql
│   │   │   │   ├── exchange_test.go
│   │   │   │   ├── sql.sql
│   │   │   │   ├── test_test.go
│   │   │   │   ├── user_asset_log.go
│   │   │   │   └── user_asset_log_test.go
│   │   │   ├── exchange_service_impl.go
│   │   │   ├── model
│   │   │   │   ├── exchange.go
│   │   │   │   ├── model.go
│   │   │   │   ├── shading.go
│   │   │   │   └── shading_test.go
│   │   │   └── user_assets_service_impl.go
│   │   ├── metrics
│   │   │   └── quota_limiter.go
│   │   ├── README.MD
│   │   ├── schedule
│   │   │   └── schedule.go
│   │   ├── test_test.go
│   │   ├── user_assets_service.go
│   │   ├── user_assets_service_test.go
│   │   └── utils
│   │       └── utils.go
│   ├── consts
│   │   └── consts.go
│   ├── cron
│   │   ├── delete_transaction.go
│   │   └── reconciliation.go
│   ├── errors
│   │   └── error.go
│   ├── escrow
│   │   ├── apple
│   │   │   ├── escrow.go
│   │   │   ├── local_validation.go
│   │   │   └── payload.go
│   │   ├── cashfree
│   │   │   ├── escrow.go
│   │   │   ├── payload.go
│   │   │   └── withdraw.go
│   │   ├── codapay
│   │   │   ├── escrow.go
│   │   │   ├── order_test.go
│   │   │   └── payload.go
│   │   ├── dlocal
│   │   │   ├── escrow.go
│   │   │   └── payload.go
│   │   ├── dlocal2
│   │   │   ├── escrow.go
│   │   │   └── payload.go
│   │   ├── dokypay
│   │   │   ├── escrow.go
│   │   │   └── payload.go
│   │   ├── fake_user.go
│   │   ├── gp
│   │   │   ├── escrow.go
│   │   │   └── payload.go
│   │   ├── oceanpay
│   │   │   ├── escrow.go
│   │   │   └── payload.go
│   │   ├── payermax
│   │   │   ├── escrow.go
│   │   │   ├── order_test.go
│   │   │   └── payload.go
│   │   ├── payload.go
│   │   ├── payssion
│   │   │   ├── escrow.go
│   │   │   └── payload.go
│   │   ├── processor
│   │   │   └── processor.go
│   │   ├── razer
│   │   │   ├── escrow.go
│   │   │   ├── order_test.go
│   │   │   └── payload.go
│   │   ├── razer_gold
│   │   │   └── escrow.go
│   │   └── tap
│   │       ├── escrow.go
│   │       └── payload.go
│   ├── handler
│   │   ├── handler.go
│   │   ├── handler_v2.go
│   │   └── payments.go
│   ├── identity
│   │   ├── dao.go
│   │   ├── identify.go
│   │   └── serializer.go
│   ├── ifsc
│   │   └── dao.go
│   ├── log
│   │   ├── bank_account.go
│   │   ├── log.go
│   │   ├── log_op.go
│   │   ├── metrics
│   │   │   └── order_metrics.go
│   │   ├── recharge_reward.go
│   │   └── reconciliation.go
│   ├── model
│   │   ├── fake_user.go
│   │   ├── identity.go
│   │   ├── ifsc.go
│   │   ├── im_assets.go
│   │   ├── income.go
│   │   ├── model.go
│   │   ├── payment.go
│   │   ├── razer_gold.go
│   │   ├── recharge_blacklist.go
│   │   ├── recharge.go
│   │   ├── recharge_reward.go
│   │   ├── refund_recover.go
│   │   ├── transfer_model.go
│   │   └── withdraw_application.go
│   ├── payload
│   │   ├── gp_rec.go
│   │   └── op.go
│   ├── refund
│   │   ├── cron.go
│   │   └── recovery.go
│   ├── service
│   │   ├── doAfterRecharge.go
│   │   ├── game
│   │   │   └── game.go
│   │   ├── income.go
│   │   ├── recharge_permission.go
│   │   ├── recharge_reward.go
│   │   ├── recharge_stats.go
│   │   ├── sku.go
│   │   ├── transaction_limit.go
│   │   └── transfer.go
│   ├── sku
│   │   ├── dao.go
│   │   └── serializer.go
│   ├── timer_billing
│   │   ├── dao.go
│   │   ├── dao_test.go
│   │   ├── payload.go
│   │   ├── payload_test.go
│   │   └── test_test.go
│   ├── transaction
│   │   ├── dao.go
│   │   ├── game.go
│   │   ├── serializer.go
│   │   ├── test_test.go
│   │   ├── transaction_listener.go
│   │   └── transaction_listener_test.go
│   ├── withdraw
│   │   └── dao.go
│   └── withdraw_service
│       ├── consts
│       │   └── consts.go
│       ├── internal
│       │   ├── dao
│       │   │   ├── db.sql
│       │   │   ├── withdraw.go
│       │   │   └── withdraw_test.go
│       │   └── impl
│       │       └── default.go
│       ├── model
│       │   └── withdraw.go
│       ├── scripts
│       │   └── clear_old_gold.go
│       ├── withdraw_service.go
│       └── withdraw_service_test.go
├── warehouse
│   ├── api
│   │   ├── api.go
│   │   └── vip.go
│   ├── consts
│   │   └── consts.go
│   ├── dao
│   │   └── dao.go
│   ├── errors
│   │   └── error.go
│   ├── handler
│   │   └── handler.go
│   ├── log
│   │   └── log.go
│   ├── model
│   │   └── models.go
│   ├── payload
│   │   └── payload.go
│   └── serializer
│       └── serializer.go
└── wish
    ├── api
    │   └── api.go
    ├── dao
    │   └── dao.go
    ├── handler
    │   └── handler.go
    ├── model
    │   └── model.go
    └── serializer
        └── serializer.go


```

<++>
