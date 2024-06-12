# supernamego

## 1. Overview

supernamego is a Go client for sns, it supports service discovery like dns and dynamic configuration.

## 2. Quick Start

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

	addr, cluster, err := supernamego.Lookup(context.Background(), "sns/https.nginx")
	if err != nil {
		fmt.Printf("supernamego lookup: %v\n", err)
		return
	}
	fmt.Printf("cluster=%s, address=%v\n", cluster, addr)
}
```
