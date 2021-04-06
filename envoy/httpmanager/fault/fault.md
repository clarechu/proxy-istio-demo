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

fault.http.delay.fixed_delay_percent
标头匹配时将延迟的请求的百分比。默认为配置中指定的 delay_percent，否则为0。
仅当将过滤器配置为延迟时，此运行时密钥才可用。

fault.http.delay.fixed_duration_ms
延迟持续时间（以毫秒为单位）。如果未指定，将使用配置中指定的 fixed_duration_ms。
如果运行时和配置中都缺少此字段，则不会插入任何延迟。
仅当将过滤器配置为延迟时，此运行时密钥才可用。

fault.http.max_active_faults
Envoy将通过故障过滤器注入的所有类型的活动故障的最大数量。
可以在希望将故障100％注入的情况下使用此方法，但是用户希望避免出现过多的意外并发故障请求导致资源约束问题的情况。
如果未指定，将使用max_active_faults设置。

fault.http.rate_limit.response_percent
已注入响应速率限制错误的请求的百分比。默认为在百分比字段中设置的值。
仅当将过滤器配置为限制响应速率时，此运行时密钥才可用。

注意，特定下游群集的故障过滤器运行时设置会覆盖默认设置（如果存在）。
以下是下游特定的运行时密钥：

fault.http.<downstream-cluster>.abort.abort_percent

fault.http.<downstream-cluster>.abort.http_status

fault.http.<downstream-cluster>.delay.fixed_delay_percent

fault.http.<downstream-cluster>.delay.fixed_duration_ms

下游群集名称取自 HTTP x-envoy-downstream-service-cluster 标头。如果在运行时中找不到以下设置，则默认为全局运行时设置，而全局运行时设置默认为配置设置。