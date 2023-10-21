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

func EatAll(e Eater) {
	e.PutIn()
	e.Chew()
	e.Swallow()
}

	e.Start(":8080")
}
