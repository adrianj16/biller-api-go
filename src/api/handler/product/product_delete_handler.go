package handler

import (
	"biller-api/src/api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductDelete(c *gin.Context) {
	productUseCase := usecase.Product{}
	productID, err := getProductIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, productUseCase.Delete(productID))
}
