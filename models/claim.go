package models

import (
	"github.com/dgrijalva/jwt-go"
)

/*Claim es la estructura usada para procesar el JWT*/
type Claim struct {
	Email string `json:"email"`
	ID    int    `json:"_id,omitempty"`
	jwt.StandardClaims
}
