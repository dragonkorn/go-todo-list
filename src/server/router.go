package server

import (
	"net/http"

	"github.com/dragonkorn/go-todo-list/src/controllers"
	"github.com/labstack/echo"
)

func getRouter(router *echo.Echo) *echo.Echo {
	router.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Hello world</h1>")
	})

	router.File("/index", "public/templates/index.html")

	router.GET("/ping", controllers.Ping)

	return router
}
