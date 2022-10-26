package main

import (
	"assignment-2/controller"
	"assignment-2/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Start()

	if err != nil {
		fmt.Println("error start database :", err)
	}

	ctl := controller.New(db)

	router := gin.Default()

	router.GET("/orders", ctl.GetOrdersHandler)
	router.POST("/orders", ctl.CreateOrderHandler)
	router.PUT("/orders/:orderID", ctl.UpdateOrderHandler)
	router.DELETE("/orders/:orderID", ctl.DeleteOrderHandler)

	router.Run("localhost:8080")
}