### 终端代理配置(Linux)

- 工具：`proxychains4`

> ArchLinux 配置

- step1: `yay -S proxychains4`

- step2: `sudo vim /etc/proxychains.conf`

```shell
[ProxyList]
# add proxy here ...
# meanwile
# defaults set to "tor"
socks4 	127.0.0.1 1087
socks5  127.0.0.1 1087
http 127.0.0.1 8087
```

- step3: `proxychains4 curl -I https://www.google.com` ，返回如下结果则终端配置代理成功

```
[proxychains] config file found: /etc/proxychains.conf
[proxychains] preloading /usr/lib/libproxychains4.so
[proxychains] DLL init: proxychains-ng 4.14
[proxychains] Strict chain  ...  127.0.0.1:1087  ...  127.0.0.1:1087  ...  127.0.0.1:8087  ...  www.google.com:443  ...  OK
HTTP/2 200
content-type: text/html; charset=ISO-8859-1
p3p: CP="This is not a P3P policy! See g.co/p3phelp for more info."
date: Thu, 18 Mar 2021 13:06:56 GMT
server: gws
x-xss-protection: 0
x-frame-options: SAMEORIGIN
expires: Thu, 18 Mar 2021 13:06:56 GMT
cache-control: private
set-cookie: 1P_JAR=2021-03-18-13; expires=Sat, 17-Apr-2021 13:06:56 GMT; path=/; domain=.google.com; Secure
set-cookie: NID=211=YfVL5RT-NQ6ZRYrF--ZbNe8AKU1efKkpKGlwhGESNIrD1cH0o8FJCFuZdKL7Bq971OVuvG11pGYQ1qb6T_Wb2wG50na6YvG2JnYjrjUeDW53jdKkvX83JtMmNaeNuFBDJ6xTG7QSeBXZO5xSz2See2IurmHVA2rZC9awUPZcz6I; expires=Fri, 17-Sep-2021 13:06:56 GMT; path=/; domain=.google.com; HttpOnly
alt-svc: h3-29=":443"; ma=2592000,h3-T051=":443"; ma=2592000,h3-Q050=":443"; ma=2592000,h3-Q046=":443"; ma=2592000,h3-Q043=":443"; ma=2592000,quic=":443"; ma=2592000; v="46,43"

```

> 应用

- 在需要代理的命令前加上 pproxychains4 即可通过代理访问
  - proxychains4 ssh -D 9000 user@host
