# 限流策略配置

```yaml
                  - name: envoy.filters.http.local_ratelimit
                    typed_config:
                      "@type": type.googleapis.com/envoy.extensions.filters.http.local_ratelimit.v3.LocalRateLimit
                      stat_prefix: http_local_rate_limiter
                      token_bucket:
                        max_tokens: 1
                        tokens_per_fill: 1
                        fill_interval: 10s
                      filter_enabled:
                        runtime_key: local_rate_limit_enabled
                        default_value:
                          numerator: 100
                          denominator: HUNDRED
                      filter_enforced:
                        runtime_key: local_rate_limit_enforced
                        default_value:
                          numerator: 100
                          denominator: HUNDRED
                      response_headers_to_add:
                        - append: true
                          header:
                            key: x-local-rate-limit
                            value: 'true'
```

上面的配置主要是讲解

```bash
max_tokens: 1
tokens_per_fill: 1
fill_interval: 10s
```

10s之内只允许一个请求