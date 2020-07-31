package handlers

import (
	"github.com/edtap14/proyecto-productos-back/routes"
	"github.com/labstack/echo/v4"
)

// Handlers es la función que nos permite realizar las rutas de las diferentes páginas
func Handlers() {
	e := echo.New()
	e.GET("/", routes.Home)
	e.POST("/login", routes.Login)
	e.GET("/search", routes.Search)
	e.Logger.Fatal(e.Start(":3030"))
}
