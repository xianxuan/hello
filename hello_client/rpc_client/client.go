package rpc_client

import (
	"fmt"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
	"hello/hello_world"
)

var (
	HelloClient hello_world.HelloWorldClient
)

func Init() {
	service := micro.NewService(
		micro.Name("hello_world"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init()

	greeter := hello_world.NewHelloWorldClient("hello_world", service.Client()) // 创建服务hello_world的client对象, 以便调用其中定义的RPC方法'Hello'

	rsp, err := greeter.Hello(context.TODO(), &hello_world.HelloRequest{Name: "Alice"}) // 传入HelloWorldRequest对象作为调用RPC方法的参数'Hello'
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
	HelloClient = hello_world.NewHelloWorldClient("hello_world", service.Client())
}
