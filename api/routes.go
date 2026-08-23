package api

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/products", GetProducts)
	r.POST("/products", CreateProduct)

	r.GET("/users", GetUsers)
	r.POST("/users", CreateUser)

	r.GET("/orders", GetOrders)
	r.POST("/orders", CreateOrder)

	r.GET("/inventory", GetInventory)
	r.POST("/inventory", CreateInventory)

	return r
}
