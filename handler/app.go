package handler

import (
	"assigment2_aden/infra/database"
	orderpostgres "assigment2_aden/repository/order_repository/order_postgres"
	"assigment2_aden/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	database.InitDatabase()
	db := database.GetDatabaseInstance()
	orderRepo := orderpostgres.NewOrderPG(db)

	orderService := service.NewOrderService(orderRepo)

	orderHandler := NewOrderHandler(orderService)

	r := gin.Default()

	r.POST("/order", orderHandler.CreateOrder)
	r.GET("/order", orderHandler.GetOrders)
	r.GET("order/:orderID", orderHandler.GetAnOrderByID)
	r.PUT("order/:orderID", orderHandler.UpdateOrder)
	r.PATCH("order/:orderID", orderHandler.UpdateOrder)
	r.DELETE("order/:orderID", orderHandler.DeleteOrder)
	r.Run(":8080")

}
