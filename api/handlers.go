package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Listado de productos",
	})
}

func CreateProduct(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{
		"message": "Producto creado",
	})
}

func GetUsers(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Listado de usuarios",
	})
}

func CreateUser(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado",
	})
}

func GetOrders(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Listado de pedidos",
	})
}

func CreateOrder(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{
		"message": "Pedido creado",
	})
}

func GetInventory(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Listado de inventario",
	})
}

func CreateInventory(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{
		"message": "Inventario registrado",
	})
}
