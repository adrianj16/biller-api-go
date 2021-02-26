package handler

import (
	"biller-api/src/api/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getProductIDFromContext(c *gin.Context) (int, error) {
	productIDStr := c.Param("ID")
	if productIDStr == "" {
		return 0, fmt.Errorf("ProductID Empty")
	}
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return 0, fmt.Errorf("ProductID Invalid Must be Integer")
	}
	return productID, nil
}

func getProductFromContext(c *gin.Context) (model.Product, error) {
	product := model.Product{}
	if err := c.BindJSON(&product); err != nil {
		return model.Product{}, fmt.Errorf("Invalid JSON. " + err.Error())
	}
	return product, nil
}
