package server

import (
	"github.com/dragonkorn/go-todo-list/src/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Init initial server
func Init() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	db.Connect()

	e = getRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}
