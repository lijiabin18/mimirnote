```
//user_auth 查询
select count(id) from user_authorizations where account_type='t_c';

select count(id) from user_authorizations where account_type='sms';

select a.id,a.user_id,a.account_type,a.account_id as aid,b.account_id as bid from user_authorizations as
a left join (select user_id,account_type,account_id from user_authorizations where account_type='t_c') as b on a.user_id=b.user_id where a.account_id != b.account_id and a.account_type='sms';
```

| 登录类型   | 账户数  |
| ---------- | ------- |
| sms        | 7035993 |
| truecaller | 1569951 |

- 其中，sms 账号与 truecaller 账号 信息不同的是 676 个，占所有 truecaller 账号的
  千分之四= 676 / 1569951
- 其中，sms 账号与 truecaller 账号 信息(号码)相同的有 240694 个，占所有 truecaller 账号的
  百分之十五= 240694 /1569951

```
gift_id:73 ,fail_country:[PK TR US ET CN BD MY]=======
gift_id:68 ,fail_country:[ET CN BD MY PK SG TR US]=======
gift_id:680 ,fail_country:[CN BD MY ET SG TR US]=======
gift_id:839 ,fail_country:[CN BD MY ET TR US]=======
gift_id:830 ,fail_country:[TR US CN BD MY ET]=======
gift_id:895 ,fail_country:[TR CN BD US ET]=======
gift_id:764 ,fail_country:[MY ET TR US CN BD]=======
gift_id:893 ,fail_country:[US CN BD MY ET TR]=======
gift_id:681 ,fail_country:[TR BD US ET SG]=======
gift_id:852 ,fail_country:[ET BD]=======
gift_id:851 ,fail_country:[ET TR CN US MY]=======
gift_id:897 ,fail_country:[MY ET TR US CN BD]=======
gift_id:377 ,fail_country:[PK SG TR US PH ET CN MY]=======
gift_id:400 ,fail_country:[PH ET PK SG US TR CN BD MY]=======
gift_id:358 ,fail_country:[PK SG US PH ET CN TR MY]=======
gift_id:118 ,fail_country:[CN MY PK TR US ET]=======
gift_id:86 ,fail_country:[PH ET CN TR MY PK SG US]=======
gift_id:899 ,fail_country:[MY ET TR US CN BD]=======
gift_id:409 ,fail_country:[BD TR US PH ET CN MY PK SG]=======
gift_id:67 ,fail_country:[CN PK US PH ET BD MY SG TR]=======
gift_id:374 ,fail_country:[PH ET CN MY PK TR US BD SG]=======
gift_id:82 ,fail_country:[US PH ET CN BD MY PK TR]=======
gift_id:221 ,fail_country:[MY PK SG US PH ET CN TR]=======
gift_id:330 ,fail_country:[US ET MY BD PK SG TR]=======
gift_id:623 ,fail_country:[SG TR US ET CN BD MY PK]=======
gift_id:370 ,fail_country:[PK SG US PH ET CN BD TR]=======
gift_id:88 ,fail_country:[BD TR PH SG US ET CN MY PK]=======
gift_id:371 ,fail_country:[PK SG US PH ET CN TR MY]=======
gift_id:120 ,fail_country:[TR MY PK SG US PH ET CN]=======
gift_id:84 ,fail_country:[US PH ET CN TR MY PK SG]=======
gift_id:373 ,fail_country:[BD MY PK TR US CN SG PH ET]=======
gift_id:78 ,fail_country:[BD PK SG TR ET CN MY US PH]=======
gift_id:77 ,fail_country:[SG TR US ET CN BD MY PK]=======
gift_id:622 ,fail_country:[TR PH ET CN BD MY PK SG US]=======
gift_id:368 ,fail_country:[US ET CN BD PK PH MY SG TR]=======
gift_id:99 ,fail_country:[PH CN ET MY PK SG TR US]=======
gift_id:322 ,fail_country:[CN PK SG US PH BD MY TR ET]=======
gift_id:145 ,fail_country:[MY PK TR ET PH CN SG]=======
gift_id:407 ,fail_country:[TR US PH CN BD MY PK ET]=======
gift_id:70 ,fail_country:[MY PK TR US PH ET CN BD]=======
gift_id:242 ,fail_country:[PH CN BD MY PK SG ET US]=======
gift_id:206 ,fail_country:[US PH ET CN SG MY PK TR]=======
gift_id:408 ,fail_country:[PK SG TR ET CN BD MY US PH]=======
gift_id:327 ,fail_country:[TR PH ET CN BD PK SG MY US]=======
gift_id:366 ,fail_country:[CN BD MY PK TR US ET SG PH]=======
gift_id:256 ,fail_country:[CN MY SG TR BD PK US PH ET]=======
gift_id:405 ,fail_country:[CN MY PK BD SG TR US PH ET]=======
gift_id:239 ,fail_country:[BD TR PK SG US PH ET CN]=======
gift_id:235 ,fail_country:[CN BD PK ET MY SG TR US PH]=======
gift_id:92 ,fail_country:[SG TR US PK BD MY PH ET CN]=======
gift_id:349 ,fail_country:[US PH ET SG TR MY PK]=======
gift_id:361 ,fail_country:[SG TR CN BD MY PK US PH ET]=======
gift_id:369 ,fail_country:[TR US PH SG BD MY PK ET CN]=======
gift_id:406 ,fail_country:[TR BD MY PK SG US PH ET]=======
gift_id:648 ,fail_country:[SG TR US PH ET BD MY PK]=======
gift_id:89 ,fail_country:[PK SG TR US PH ET BD MY]=======
gift_id:838 ,fail_country:[CN TR US ET]=======
gift_id:81 ,fail_country:[SG US PH BD MY PK TR ET CN]=======
gift_id:98 ,fail_country:[PK TR PH BD MY SG US ET CN]=======
gift_id:649 ,fail_country:[US PH ET CN SG MY PK TR]=======
gift_id:397 ,fail_country:[MY PK TR US PH ET SG]=======
gift_id:831 ,fail_country:[US CN BD MY ET TR]=======
gift_id:837 ,fail_country:[CN TR US ET]=======
gift_id:896 ,fail_country:[US CN BD MY ET TR]=======
gift_id:829 ,fail_country:[BD US ET TR CN]=======
gift_id:277 ,fail_country:[CN BD MY ET TR US]=======
gift_id:232 ,fail_country:[SG US PH ET CN TR MY PK]=======
gift_id:688 ,fail_country:[TR BD MY ET SG US]=======
gift_id:894 ,fail_country:[TR US CN BD MY ET]=======
gift_id:647 ,fail_country:[PK PH US ET CN TR MY]=======
gift_id:354 ,fail_country:[PH TR US ET MY PK
<++>
```

<++>
