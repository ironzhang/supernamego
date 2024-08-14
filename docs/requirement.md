# requirement

## 路由需求

1. 默认访问本机房的 default.k8s 集群（这一规则最好是可配置的）
2. 没有机房信息，则访问 dev.default.k8s 集群
4. 支持各种自定义的路由需求（暂定用脚本支持）
5. 同一个 node 上的 pod 共享一个 agent

