package handler

import (
	"biller-api/src/api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductGetAll(c *gin.Context) {
	productUseCase := usecase.Product{}
	c.JSON(http.StatusOK, productUseCase.GetAll())
}
