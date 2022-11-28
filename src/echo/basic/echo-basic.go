package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/name/:username", func(c echo.Context) error {
		username := c.Param("username")
		return c.String(http.StatusOK, "You reached with "+username)
	})

	e.Logger.Fatal(e.Start(":5500"))
}
