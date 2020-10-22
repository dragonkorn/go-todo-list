package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Ping - Pong
func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "Pong")
}
