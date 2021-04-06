# jwt 认证

获取jwks

```bash
$ wget https://raw.githubusercontent.com/istio/istio/release-1.9/security/tools/jwt/samples/jwks.json > jwks.json 

```

获取jwt token

```bash
export TOKEN=$(curl https://raw.githubusercontent.com/istio/istio/release-1.9/security/tools/jwt/samples/demo.jwt -s) 

```


* 请求当前服务

```bash
# 验证token是否存在

$ echo $TOKEN

$ curl localhost:8081/healthz -H "Authorization: Bearer $TOKEN" -v
```
