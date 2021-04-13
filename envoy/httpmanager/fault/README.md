# 故障注入

abort

如果指定，则过滤器将根据对象中的值 终止请求

```yaml
abort:
  http_status: 504
  percentage:
    numerator: 50
    denominator: HUNDRED
```

numerator: 如果标头匹配，将中止的请求的百分比。默认为config中指定的 abort_percent。如果配置不包含 中止块，则abort_percent默认为0。
http_status: HTTP状态代码，用作标头匹配时将中止的请求的响应状态代码。默认为配置中指定的HTTP状态代码。如果配置不包含中止块，则 http_status默认为0。
grpc_status: gRPC状态代码，将用作请求的响应状态代码，如果标头匹配，则将中止该请求。默认为配置中指定的gRPC状态代码。
denominator: 枚举类型 
```bash
    # **Example**: 1/100 = 1%.
    HUNDRED = 0;
    #  **Example**: 1/10000 = 0.01%.
    TEN_THOUSAND = 1;
    #  **Example**: 1/1000000 = 0.0001%.
    MILLION = 2;
```


```yaml
delay:
  fixed_delay:
    seconds: 5
  percentage:
    numerator: 10
```
numerator: 标头匹配时将延迟的请求的百分比。默认为配置中指定的 delay_percent，否则为0。

fixed_delay: 延迟持续时间（以毫秒为单位）。如果未指定，将使用配置中指定的 fixed_duration_ms。 
如果运行时和配置中都缺少此字段，则不会插入任何延迟。

max_active_faults: Envoy将通过故障过滤器注入的所有类型的活动故障的最大数量。
可以在希望将故障100％注入的情况下使用此方法，但是用户希望避免出现过多的意外并发故障请求导致资源约束问题的情况。
如果未指定，将使用max_active_faults设置。
