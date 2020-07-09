package routes

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/edtap14/proyecto-productos-back/models"
	"github.com/labstack/echo/v4"
)

/*Login es la ruta para el registro de usuarios*/
func Login(c echo.Context) (err error) {
	u := new(models.Usuario)
	if err = c.Bind(u); err != nil {
		return
	}
	// Throws unauthorized error
	if u.Nombre != "TeamBorregos" || u.Password != "1234" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Team Borrego"
	claims["email"] = "team.borrego@team.borrego"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
