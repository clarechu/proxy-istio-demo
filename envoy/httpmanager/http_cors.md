# 设置完成后 请求

在浏览器console 中请求

从中遇到的问题及解决方案

```bash
https://github.com/envoyproxy/envoy/issues/11776
```

```bash
fetch('http://localhost:8081/healthz').then(function(response) {
    console.log(response.text());
})
```

![img.png](img.png)