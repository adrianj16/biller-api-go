package handler

import (
	"biller-api/src/api/model"
	"biller-api/src/api/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var productUseCase usecase.Product

func ProductGetAll(c *gin.Context) {
	c.JSON(http.StatusOK, productUseCase.GetAll())
}

func ProductGet(c *gin.Context) {
	productID, err := getProductID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, productUseCase.Get(productID))
}

func ProductUpdate(c *gin.Context) {
	productID, err := getProductID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	product, err := getProduct(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	product.ID = productID

	c.JSON(http.StatusOK, productUseCase.Update(product))
}

func ProductDelete(c *gin.Context) {
	productID, err := getProductID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, productUseCase.Delete(productID))
}

func ProductCreate(c *gin.Context) {
	product, err := getProduct(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, productUseCase.Create(product))
}

func getProductID(c *gin.Context) (int, error) {
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

func getProduct(c *gin.Context) (model.Product, error) {
	product := model.Product{}
	if err := c.BindJSON(&product); err != nil {
		return model.Product{}, fmt.Errorf("Invalid JSON. " + err.Error())
	}
	return product, nil
}
