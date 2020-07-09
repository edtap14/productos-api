package handlers

import (
	"github.com/edtap14/proyecto-productos-back/routes"
	"github.com/labstack/echo/v4"
)

/*Handlers es la función que nos permite realiar las rutas de las diferentes paginas*/
func Handlers() {
	e := echo.New()
	e.GET("/", routes.Home)
	e.POST("/login", routes.Login)
	e.POST("/product-search", routes.ProductSearch)
	e.Logger.Fatal(e.Start(":8080"))
}
