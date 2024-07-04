# supernamego 详细设计

[toc]

## 1. 概述

[SNS 总体设计](https://github.com/ironzhang/sns/blob/master/docs/design_cn.md) 如下图所示：

![](./diagram/architecture.png)

图1 SNS 总体设计

supernamego 作为 [SNS](https://github.com/ironzhang/sns) 的 Go 语言 SDK，其主要职责如下：

* 提供类 DNS 的服务发现功能
* 提供动态配置功能（TODO）

## 2. 方案设计

要实现类 DNS 的服务发现和动态配置这两大核心功能，SDK 需要做如下工作：

* 通过网络与 sns-agent 模块交互，发送要订阅的域名和配置空间。
* 通过文件读取订阅的域名地址及动态配置，并感知其变更。
* 实现路由策略模块，支持七层语义的流量调度。

除此之外，我们还有一些其他的设计考量：

* 提供类似 DNS /etc/resolv.conf 的配置文件，让 SDK 的某些参数配置化
* 负载均衡算法接口化，允许用户设置自定义的负载均衡算法
* 支持用户传入 IP:Port 做解析，方便测试

## 3. 模块设计

