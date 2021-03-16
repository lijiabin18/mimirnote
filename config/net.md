#### 1. Git 代理配置(socks5)

1. 确定本地配置的 socks5 端口。比如配置的是 1080
2. 为当前 git 仓库配置 socks5 代理

```
//http
git config --local http.proxy 'socks5://127.0.0.1:1087'
//https
git config --local https.proxy 'socks5://127.0.0.1:1087'

// 取消代理
git config --local --unset http.proxy
git config --local --unset https.proxy

// 参数说明
- local : 当前仓库，即当前仓库下的 (./git/config)
- global: 全局，即读/写 当前用户全局配置文件(~/.gitconfig)
- system: 读写系统全局的配置文件(/etc/giftconfig)
```
