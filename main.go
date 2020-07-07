package main

import (
	"net/http"

	"github.com/edtap14/proyecto-productos-back/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	e.POST("/login", handlers.Login)
	e.Logger.Fatal(e.Start(":8080"))
}
