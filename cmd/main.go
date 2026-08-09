package main

import (
	"fmt"

	"github.com/DiegoHachig/ecommerce-management-system-go/database"
	"github.com/DiegoHachig/ecommerce-management-system-go/internal/inventory"
	"github.com/DiegoHachig/ecommerce-management-system-go/internal/orders"
	"github.com/DiegoHachig/ecommerce-management-system-go/internal/products"
	"github.com/DiegoHachig/ecommerce-management-system-go/internal/reports"
	"github.com/DiegoHachig/ecommerce-management-system-go/internal/users"
)

type inventoryAdapter struct {
	*inventory.InventoryRecord
}

func (a *inventoryAdapter) GetProductID() string {
	return fmt.Sprintf("%d", a.InventoryRecord.GetProductID())
}

func main() {

	database.ConnectDB()
	// AutoMigrate crea o actualiza automáticamente
	// las tablas necesarias en MySQL.
	database.DB.AutoMigrate(
		&products.ProductEntity{},
		&users.UserEntity{},
	)
	repo := &products.MySQLRepository{}

	service := products.NewProductService(repo)

	product, err := products.NewProduct(
		1,
		"Laptop Lenovo",
		850,
		10,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = service.RegisterProduct(*product)

	if err != nil {
		fmt.Println(err)
		return
	}

	list, _ := repo.GetAll()

	products.ShowProducts(list)

	fmt.Println("=== USUARIOS ===")

	userRepo := &users.MySQLRepository{}
	userService := users.NewUserService(userRepo)

	user, err := users.NewUser(
		1,
		"Diego Hachig",
		"diego@email.com",
		"1234567890",
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = userService.RegisterUser(*user)
	if err != nil {
		fmt.Println(err)
		return
	}

	userList, _ := userRepo.GetAll()
	users.ShowUsers(userList)

	fmt.Println("=== PEDIDOS ===")

	orderRepo := &orders.MemoryRepository{}

	orderService := orders.NewOrderService(
		orderRepo,
	)

	order, err := orders.NewOrder(
		1,
		1,
		850,
		"Pendiente",
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = orderService.CreateOrder(
		*order,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	orderList, _ := orderRepo.GetAll()

	orders.ShowOrders(orderList)

	fmt.Println("=== INVENTARIO ===")

	inventoryRepo := &inventory.MemoryRepository{}

	inventoryService := inventory.NewInventoryService(
		inventoryRepo,
	)

	item, err := inventory.NewInventory(
		1,
		10,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = inventoryService.AddInventory(*item)

	if err != nil {
		fmt.Println(err)
		return
	}

	inventoryList, _ := inventoryRepo.GetAll()

	inventory.ShowInventory(
		inventoryList,
	)
	fmt.Println("=== REPORTES ===")

	reportService := reports.NewReportService()

	reportService.GenerateSalesReport(850)

	reportService.GenerateInventoryReport(10)
}
