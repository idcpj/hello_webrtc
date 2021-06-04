- [webrtc](#webrtc)
    * [快速开始](#----)
    * [注意:](#---)


## webrtc

### 视频会议
启动 service
```
cd app
go run main.go
```

访问浏览器

```
https://127.0.0.1:8083/web/?uid=1&roomid=1
https://127.0.0.1:8083/web/?uid=2&roomid=1
```

### 注意:
如果使用局域网ip,需要配置https

方式一:app/main.go 中以 整数的方式启动

方式二: apache 配置
1. 打开模块 "mod_proxy.so,mod_proxy_http.so,mod_rewrite.somod_ssl.so,mod_proxy_wstunnel.so"
2. 在httpd-ssl.conf 配置 https 的站点
```
RewriteEngine On
RewriteCond %{HTTP:UPGRADE} ^WebSocket$ [NC,OR]
RewriteCond %{HTTP:CONNECTION} ^Upgrade$ [NC]
 
RewriteRule ^/wss(.*)    ws://192.168.0.130:8083/wss$1 [P,L]


ProxyPass  / http://127.0.0.1:8083/
ProxyPassReverse / http://127.0.0.1:8083/
```


### 远程协助

desktop 用于被控端
crl_desktop 用于控制端
```
https://127.0.0.1:8083/web/desktop.html?uid=31&roomid=33
https://127.0.0.1:8083/web/ctl_desktop.html?uid=32&roomid=33
```