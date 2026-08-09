package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "root:C0nsult0r$2026ec@tcp(127.0.0.1:3306)/ecommerce_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("Base de datos conectada correctamente")
}
