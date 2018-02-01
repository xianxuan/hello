package http_server

import (
	"github.com/labstack/echo"
	"hello/hello_client/http_server/api"
	"net/http"
)

func Init() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	MountGroup(e)

	e.Logger.Fatal(e.Start(":1323"))
}
func MountGroup(e *echo.Echo) {
	MountApi(e.Group("/hello"))
}
func MountApi(group *echo.Group) {
	group.GET("/get_hello/:id", api.GetHello)
}
