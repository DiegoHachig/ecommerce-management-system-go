package main

import (
	"fmt"

	"github.com/DiegoHachig/ecommerce-management-system-go/api"
)

func main() {

	router := api.SetupRouter()

	fmt.Println("API ejecutándose en http://localhost:8080")

	router.Run(":8080")
}
