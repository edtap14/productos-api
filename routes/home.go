package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*Home es la ruta inicial de la api*/
func Home(c echo.Context)(err error){
	return c.String(http.StatusOK, "Hello world")
}