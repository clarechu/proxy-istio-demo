# 健康检查

```bash
      health_checks:
        timeout: 2s
        interval: 5s
        interval_jitter: 1s
        unhealthy_threshold: 1
        healthy_threshold: 3
        no_traffic_interval: 60s
        event_log_path: /dev/stdout
        http_health_check:
          path: /healthz
```