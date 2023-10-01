package handler

import (
	datatransferobject "assigment2_aden/data_transfer_object"
	"assigment2_aden/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	OrderService service.OrderService
}

func NewOrderHandler(OrderService service.OrderService) orderHandler {
	return orderHandler{
		OrderService: OrderService,
	}

}

func (oh *orderHandler) CreateOrder(ctx *gin.Context) {

	var newOrderRequest datatransferobject.NewOrderRequest

	err := ctx.ShouldBindJSON(&newOrderRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message ": "Invalid json request",
		})
		return
	}

	err = oh.OrderService.CreateOrder(newOrderRequest)

	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Suksesss",
	})
}

func (oh *orderHandler) GetOrders(ctx *gin.Context) {
	response, err := oh.OrderService.GetOrders()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (oh *orderHandler) GetAnOrderByID(ctx *gin.Context) {
	OrderID := ctx.Param("orderID")

	ConvOrderID, err := strconv.Atoi(OrderID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := oh.OrderService.GetOrderByID(ConvOrderID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (oh *orderHandler) UpdateOrder(ctx *gin.Context) {

	OrderID := ctx.Param("orderID")

	ConvOrderID, err := strconv.Atoi(OrderID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var newOrderRequest datatransferobject.NewOrderRequest

	err = ctx.ShouldBindJSON(&newOrderRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message ": "Invalid json request",
		})
		return
	}

	err = oh.OrderService.UpdateOrder(ConvOrderID, newOrderRequest)

	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return

	}
	response, err := oh.OrderService.GetOrderByID(ConvOrderID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Code": http.StatusOK,

		"data": response.Data,
	})
}

func (oh *orderHandler) DeleteOrder(ctx *gin.Context) {

	OrderID := ctx.Param("orderID")

	ConvOrderID, err := strconv.Atoi(OrderID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = oh.OrderService.DeleteOrder(ConvOrderID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprintf("order with id %d not found, or delete success", ConvOrderID),
	})
}
