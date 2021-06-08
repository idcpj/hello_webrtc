
# turnserver 
https://github.com/coturn/coturn

webrtc,的访问流程为

1. 尝试直连.
2. 通过stun服务器进行穿透
3. 无法穿透则通过turn服务器中转.


## 快速入门
```
tar -zxvf turnserver.tar.gz
cd turnserver/
cp  turnserver.conf   /etc/
./turnserver
```

默认开放 3478 端口
iceServer 设置
```
{
    iceServers: [
        {
            url: 'turn:192.168.0.151',
            username: 'user',
            credential: '123'
        }
    ]
}
```

