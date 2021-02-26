package handler

import (
	"biller-api/src/api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductUpdate(c *gin.Context) {
	productUseCase := usecase.Product{}
	productID, err := getProductIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	product, err := getProductFromContext(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	product.ID = productID

	c.JSON(http.StatusOK, productUseCase.Update(product))
}
