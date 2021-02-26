package main

import (
	"biller-api/src/api/database"
	handler "biller-api/src/api/handler/product"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//Product
	r.GET("/product", handler.ProductGetAll)
	r.GET("/product/:ID", handler.ProductGet)
	r.PUT("/product/:ID", handler.ProductUpdate)
	r.DELETE("/product/:ID", handler.ProductDelete)
	r.POST("/product", handler.ProductCreate)

	database.GetDB()
	err := database.DB.Ping()
	if err != nil {
		fmt.Println(fmt.Sprintf("COULD NOT CONNECT TO DATABASE: %s", err.Error()))
	}

	if err := r.Run(); err != nil {
		fmt.Println(fmt.Sprintf("error running server %s", err.Error()))
	}
}
