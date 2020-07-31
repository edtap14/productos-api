package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Usuario es el modelo de usuario de la base de datos
type Usuario struct {
	gorm.Model
	ID              int       `json:"id"  gorm:"PRIMARY_KEY"`
	Nombre          string    `json:"nombre"`
	Apellidos       string    `json:"apellidos"`
	FechaNacimiento time.Time `json:"fechaNacimiento"`
	Email           string    `json:"email" gorm:"type:varchar(100);unique_index"`
	Password        string    `json:"password" gorm:"not null"`
}
