package main

import (
	"context"
	"fmt"

	"github.com/ironzhang/supernamego"
)

func main() {
	//	err := supernamego.AutoSetup()
	//	if err != nil {
	//		fmt.Printf("supernamego auto setup: %v\n", err)
	//		return
	//	}

	addr, cluster, err := supernamego.Lookup(context.Background(), "sns/https.myapp")
	if err != nil {
		fmt.Printf("supernamego lookup: %v\n", err)
		return
	}
	fmt.Printf("cluster=%s, address=%v\n", cluster, addr)
}
