package main

import (
	"context"

	"github.com/ironzhang/superlib/httputils/echoutil"
	"github.com/ironzhang/superlib/httputils/echoutil/echorpc"
	"github.com/labstack/echo"
)

func HandleEcho(ctx context.Context, in string, out *string) error {
	*out = in
	return nil
}

func main() {
	e := echo.New()
	e.HTTPErrorHandler = echoutil.HTTPErrorHandler
	e.POST("/echo", echorpc.HandlerFunc(HandleEcho))
	e.Start(":8000")
}
