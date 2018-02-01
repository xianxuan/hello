package main

import (
	"hello/hello_client/http_server"
	"hello/hello_client/rpc_client"
)

func main() {
	rpc_client.Init()
	http_server.Init()
}
