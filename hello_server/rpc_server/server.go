package rpc_server

import (
	"fmt"
	"github.com/micro/go-micro"
	"hello/hello_server/rpc_server/api"
	"hello/hello_world" // import proto生成的类
)

func Init() {
	service := micro.NewService(
		micro.Name("hello_world"), // 定义service的名称为hello_world
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init() // 初始化service

	hello_world.RegisterHelloWorldHandler(service.Server(), new(api.HelloWorld)) // 注册服务

	if err := service.Run(); err != nil {
		fmt.Println(err)
	} // 运行服务
}
