package api

import (
	"golang.org/x/net/context"
	"hello/hello_world"
)

type HelloWorld struct{}

func (g *HelloWorld) Hello(ctx context.Context, req *hello_world.HelloRequest, rsp *hello_world.HelloResponse) error {
	rsp.Greeting = "Hello World: " + req.Name
	return nil
} // 实现hello_world service中Hello方法
