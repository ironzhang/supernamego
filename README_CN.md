[English](./README.md) | 中文

# supernamego

## 概述

supernamego 是 sns 的 Go SDK 客户端，它提供类似 DNS 的服务发现功能。

## 快速开始

### 要求

* go version >= 1.22.3
* A working docker environment

### 安装

参见 [sns Installation](https://github.com/ironzhang/sns/tree/master?tab=readme-ov-file#installation)

### Examples

以下代码展示了如何使用 supernamego 来解析 sns 域名。

```
package main

import (
	"context"
	"fmt"

	"github.com/ironzhang/supernamego"
)

func main() {
	err := supernamego.AutoSetup()
	if err != nil {
		fmt.Printf("supernamego auto setup: %v\n", err)
		return
	}

	addr, cluster, err := supernamego.Lookup(context.Background(), "sns/https.myapp")
	if err != nil {
		fmt.Printf("supernamego lookup: %v\n", err)
		return
	}
	fmt.Printf("cluster=%s, address=%v\n", cluster, addr)
}
```
