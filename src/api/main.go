package main

import (
	"biller-api/src/api/database"
	"biller-api/src/api/handler"
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
	r.GET("/product", handler.ProductGetAll)

	database.GetDB()
	err := database.DB.Ping()
	if err != nil {
		fmt.Println(fmt.Sprintf("COULD NOT CONNECT TO DATABASE: %s", err.Error()))
	}

	if err := r.Run(); err != nil {
		fmt.Println(fmt.Sprintf("error running server %s", err.Error()))
	}
}
