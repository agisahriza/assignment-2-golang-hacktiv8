package controller

import (
	"assignment-2/database"
	"assignment-2/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db database.Database
}

func New(db database.Database) Controller {
	return Controller{
		db: db,
	}
}

func (c Controller) CreateOrderHandler(ctx *gin.Context) {
	var newOrder model.Order
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	orderResult, err := c.db.CreateOrder(newOrder)
	fmt.Println(orderResult.CustomerName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"message": "error create order",
		})
		return
	}

	ctx.JSON(http.StatusOK, orderResult)
}

func (c Controller) GetOrdersHandler(ctx *gin.Context) {
	orders, err := c.db.GetOrders()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": "500",
			"message": "error get data",
		})
	}

	ctx.JSON(http.StatusOK, orders)
}

func (c Controller) UpdateOrderHandler(ctx *gin.Context) {
	strOrderID := ctx.Param("orderID")
	intOrderID, err := strconv.Atoi(strOrderID)
	var updatedOrder model.Order
	
	if err != nil {
		fmt.Println("error convert string ID to uint ID: ", err)
		return
	}
	
	
	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	orderResult, err := c.db.UpdateOrder(uint(intOrderID), updatedOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": "404",
			"status": "fail",
			"message": err.Error(),
		})
		return	
	}

	ctx.JSON(http.StatusOK, orderResult)
}

func (c Controller) DeleteOrderHandler(ctx *gin.Context) {
	strOrderID := ctx.Param("orderID")
	intOrderID, err := strconv.Atoi(strOrderID)
	if err != nil {
		fmt.Println("error convert string ID to uint ID: ", err)
		return
	}

	err = c.db.DeleteOrder(uint(intOrderID))
	
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": "404",
			"status": "fail",
			"message": err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H {
		"message": fmt.Sprintf("order with id %v has been successfully deleted", intOrderID),
	})
}