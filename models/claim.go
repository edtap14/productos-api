package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// Claim es la estructura usada para procesar el JWT
type Claim struct {
	gorm.Model
	Email string `json:"email"`
	ID    int    `json:"_id,omitempty"`
	jwt.StandardClaims
}
