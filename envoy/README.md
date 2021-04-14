# envoy 配置详解

## 背景:

我们知道 istio 的数据面 默认使用的是`envoy`。
但是 对 envoy 了解的人比 istio 的人数少之又少,
主要的原因是envoy的文档复杂 不好读懂, 还有就是需要的专业知识比较多, 导致了解envoy 的人数比较少。
该文档主要把envoy中的主要功能 及测试的yaml 贴出来 让大家更加方便的了解envoy的运行过程, 及使用教程。
有不对的地方麻烦大家多多指正。

## 概念:

### xds: 包含 LDS/EDS/CDS/RDS

* LDS: 配置详细信息实际的源头是来自xDS API 中 Listener

* EDS: 配置详细信息实际的源头是来自xDS API 中  Endpoint 


* CDS: 群集(cluster)是一组类似的上游主机(hosts)，接受来自Envoy的流量。
集群允许均衡服务集的负载平衡，以及更好的基础架构弹性。
在Envoy的白话中，“群集(cluster)”是一组命名的主机(host)/端口(port)，Envoy将通过群集来实现流量的负载均衡。可以将cluster称为服务，微服务或API。Envoy将定期轮询CDS端点以获取群集配置。

* RDS: 路由是一组将虚拟主机(virtual hosts)与群集(cluster)匹配的规则。
虚拟主机 = 域名 + 网址 `virtual hosts = Domain + Path`
也就是说：Route的功能是将以 “Domain + Path” 形式展示的虚拟主机映射到集群。

### 名词解释:
* Host：能够进行网络通信的实体（手机，服务器等上的应用程序）。在本文档中，主机是逻辑网络应用程序。一个物理硬件可以在其上运行多个主机，只要它们中的每一个都可以独立寻址即可。

* Downstream：下游主机连接到Envoy，发送请求并接收响应。

* Upstream：上游主机从Envoy接收连接和请求，并返回响应。

* Listener：侦听器是可以由下游客户端连接到的命名网络位置（例如，端口，Unix域套接字等）。Envoy公开了下游主机连接到的一个或多个侦听器。

* Cluster：群集是Envoy连接到的一组逻辑相似的上游主机。Envoy通过服务发现来发现集群的成员。它可以选择通过主动的运行状况检查来确定集群成员的运行状况。Envoy将请求路由到的群集成员由负载平衡策略确定。

* Mesh：一组主机以进行协调以提供一致的网络拓扑。在本文档中，“ Envoy网格”是一组Envoy代理，它们构成了由许多不同服务和应用程序平台组成的分布式系统的消息传递基础。

* Runtime configuration：与Envoy一起部署的带外实时配置系统。可以更改配置设置，这将影响操作，而无需重新启动Envoy或更改主要配置。

## 安装envoy

测试步骤之前我们需要安装envoy

macos 

```bash
brew install envoy
```

windows

```bash

docker pull envoyproxy/envoy-windows-dev:latest
docker run --rm envoyproxy/envoy-windows-dev:latest --version

```

其他的操作系统请参考 官方文档

[Installing Envoy](https://www.envoyproxy.io/docs/envoy/latest/start/install)

## 配置简介

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

