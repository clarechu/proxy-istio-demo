# envoy 配置详解

流量管理 流量分配

```yaml
route:
  weighted_clusters:
    runtime_key_prefix: routing.traffic_split.helloworld
    clusters:
    - name: helloworld_v1
      weight: 50
    - name: helloworld_v2
      weight: 50

cds 配置

clusters:
  - name: helloworld_v1
    connect_timeout: 15s
    type: STRICT_DNS
    lb_policy: ROUND_ROBIN
    #http2_protocol_options: {}
    load_assignment:
      cluster_name: helloworld_v1
      endpoints:
        - lb_endpoints:
            - endpoint:
                address:
                  socket_address:
                    address: 127.0.0.1
                    port_value: 8080
  - name: helloworld_v2
    connect_timeout: 0.25s
    type: STRICT_DNS
    lb_policy: ROUND_ROBIN
    http2_protocol_options: {}
    load_assignment:
      cluster_name: helloworld_v2
      endpoints:
        - lb_endpoints:
            - endpoint:
                address:
                  socket_address:
                    address: 127.0.0.1
                    port_value: 8084
```

