package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Conection función para conectar la base de datos
func Conection() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=52908 user=gorm dbname=comparacion password=Americas1487")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
}
