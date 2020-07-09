package models

import (
	"time"
)

/*Usuario es el modelo de usuario de la base de datos*/
type Usuario struct {
	ID              int       `json:"id"`
	Nombre          string    `json:"name,omitempty"`
	Apellidos       string    `json:"apellidos,omitempty"`
	FechaNacimiento time.Time `json:"fechaNacimiento,omitempty"`
	Email           string    `json:"email,omitempty"`
	Password        string    `json:"password,omitempty"`
}
