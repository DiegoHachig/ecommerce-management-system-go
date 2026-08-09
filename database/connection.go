package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB almacena la instancia global de conexión
// a la base de datos utilizada por la aplicación.
var DB *gorm.DB

// ConnectDB establece la conexión con MySQL
// utilizando GORM como gestor ORM.

func ConnectDB() {
	// DSN contiene la información necesaria
	// para conectarse al servidor MySQL.
	dsn := "root:C0nsult0r$2026ec@tcp(127.0.0.1:3306)/ecommerce_db?charset=utf8mb4&parseTime=True&loc=Local"
	// Se establece la conexión con la base de datos.
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
	// Si la conexión falla se detiene la ejecución
	// mostrando el error encontrado.
	if err != nil {
		panic(err)
	}
	// Se guarda la instancia de conexión para ser
	// reutilizada por los diferentes módulos.
	DB = db

	fmt.Println("Base de datos conectada correctamente")
}
