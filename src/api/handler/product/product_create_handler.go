package handler

import (
	"biller-api/src/api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductCreate(c *gin.Context) {
	productUseCase := usecase.Product{}
	product, err := getProductFromContext(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, productUseCase.Create(product))
}
