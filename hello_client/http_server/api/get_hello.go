package api

import (
	"fmt"
	"github.com/labstack/echo"
	"hello/hello_client/rpc_client"
	"hello/hello_world"
	"net/http"

	"golang.org/x/net/context"
	"gopkg.in/square/go-jose.v1/json"
)

// e.GET("/users/:id", getUser)
func GetHello(c echo.Context) error {
	// User ID 来自于url `users/:id`
	id := c.Param("id")
	rsp, err := rpc_client.HelloClient.Hello(context.TODO(), &hello_world.HelloRequest{Name: id}) // 传入HelloWorldRequest对象作为调用RPC方法的参数'Hello'
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(rsp.Greeting)
	if bs, err := json.Marshal(&rsp); err != nil {
		panic(err)
		return c.String(http.StatusOK, "00")
	} else {
		return c.String(http.StatusOK, string(bs))
	}

}
