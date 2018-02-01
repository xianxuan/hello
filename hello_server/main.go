package main

import (
	"hello/hello_server/cache"
	"hello/hello_server/model"
	"hello/hello_server/rpc_server"
)

func main() {
	model.Init()
	cache.Init()
	rpc_server.Init()
}
