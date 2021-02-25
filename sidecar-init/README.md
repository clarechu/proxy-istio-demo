# macos 反向代理设置

在macos上面的用法

[iptables](https://penglei.github.io/post/transparent_proxy_on_macosx/)

```bash
$
$
$
$

```



在linux上面的用法

```bash
$ iptables -t nat -A PREROUTING -p tcp --dport 8080 -j REDIRECT --to-ports 8888
```
