# kill 

```yaml
- name: envoy.filters.http.kill_request
  typed_config:
    "@type": type.googleapis.com/envoy.extensions.filters.http.kill_request.v3.KillRequest
     probability:
       numerator: 100
```


numerator: 默认为0, 杀死请求的百分比