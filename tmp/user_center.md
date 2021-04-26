#### 用户信息修改 后台化 #7451

##### 新增 修改用户绑定手机号

> 根据 user_id 查询用户信息

- 查询信息包括
  - user_id
  - country
  - language
  - 所有登录方式(apple 账号，谷歌邮箱，facebook 账号，电话，twitter 账号)

> 根据 user_id 查询用户信息

```
url /users/{user_id}/base/
method GET
params user_id

response:
{
    "data": {
        "bind_accounts": {
            "apple_id": "",
            "apple_name": "",
            "facebook_id": "",
            "facebook_name": "",
            "g_mail": "",
            "phone": "",
            "twitter_id": "",
            "twitter_name": ""
        },
        "country": "IN",
        "language": "te",
        "user_id": 39786842
    },
    "msg": "ok"
}
```

> 绑定用户手机号

```
url /users/bind/
method PUT
req_body:
{
        "source": "sms",
        "source_id":"0861570",
        "user_id": "39786842",
        "cause":"test bind user account "

}

response:
	//绑定失败，手机号被绑定了
	{
    "err_code": 92033,
    "err_msg": "This phone has been bound to the account."
	}
	//绑定成功
	{
		"msg": "ok"
	}
```

<++>

```
url /users/{user_id}/info/
Method GET
urlParams user_id

response:
	"user_id": ,
	"country":"",
	"language":"",
	"bind_accounts":
	{
"apple_name" :"",
"apple_id" :"",
"g_mail" :"",
"facebook_name": "",
"facebook_id" :"",
"twitter_name":"",
"twitter_id":"",
"phone":"",
	}

```

[](++) <++>

---

##### 靓号回收

- 若当前靓号 status=0，即未使用状态，只能点击删除；若当前靓号 status=1，则即能删除也能取消

```
POST /display_id/recall/

body:
	- "user_id":42430217,
  - "remove":false,//true时，删除靓号记录

**resp**
{status: 'OK'}
```

---

##### 靓号记录

```
GET /display_id/records/

params:
	- "app":"funshare",
	- "user_id":
	- "page":
	- "limit":

**resp**
  "data": [
        {
            "created_at": 1614856383,
            "display_id": 428428,
            "status": true,
            "user_id": 42430217
        },
        {
            "created_at": 1614840970,
            "display_id": 1112387,
            "status": true,
            "user_id": 36485522
        },
        {
            "created_at": 1614818318,
            "display_id": 1112386,
            "status": true,
            "user_id": 43077517
        },
        {
            "created_at": 1614797246,
            "display_id": 4443786,
            "status": true,
            "user_id": 27178299
        }
				]
				}

```

---

##### 靓号能否使用

```
URL /display_id/${display_id}/usable/

path_params: "display_id": //不能为空

params:
	- "user_id": //默认为空

resp:
{
    "data": {
        "usable": true
    },
    "msg": "ok"
}
```
