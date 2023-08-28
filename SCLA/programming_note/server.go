package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users/:id", func(c echo.Context) error {
		userID := c.Param("id")
		return c.String(http.StatusOK, "User ID: "+userID)
	})

	e.Start(":8080")
}
