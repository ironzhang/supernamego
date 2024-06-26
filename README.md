English | [中文](./README_CN.md)

# supernamego

## Overview

supernamego is a Go client for sns, it supports service discovery like dns.

## Quick Start

### Requirements

* go version >= 1.22.3
* A working docker environment

### Installation

see [sns Installation](https://github.com/ironzhang/sns/tree/master?tab=readme-ov-file#installation)

### Examples

The following code shows how to use supernamego to resolve sns domain names.

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
