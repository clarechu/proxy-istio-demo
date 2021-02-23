# proxy-istio-demo

![img.png](img.png)

## demo
一个demo的项目
启动了一个端口 `8080` 主要用来测试

```bash
$ curl http://localhost:8080\?a\=b -v


*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET /?a=b HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< hello world: demo
< Date: Mon, 22 Feb 2021 15:25:31 GMT
< Content-Length: 11
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
{"up":true}* Closing connection 0

```

响应为

```bash

body

{"up":true}

header:

hello world: demo

```



## proxy 代理服务


## sidecar-init


使用iptables的命令行 将流量打到demo服务的流量 重定向到proxy 服务


