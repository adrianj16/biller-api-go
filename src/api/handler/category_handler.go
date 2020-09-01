package handler

import (
	"biller-api/src/api/model"
	"biller-api/src/api/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var categoryUseCase usecase.Category

func CategoryUpdate(c *gin.Context) {
	categoryID, err := getCategoryID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	category, err := getCategory(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	category.ID = categoryID

	c.JSON(http.StatusOK, categoryUseCase.Update(category))
}

func CategoryDelete(c *gin.Context) {
	categoryID, err := getCategoryID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, categoryUseCase.Delete(categoryID))
}

func CategoryCreate(c *gin.Context) {
	category, err := getCategory(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, categoryUseCase.Create(category))
}

func getCategoryID(c *gin.Context) (int, error) {
	categoryIDStr := c.Param("ID")
	if categoryIDStr == "" {
		return 0, fmt.Errorf("ProductID Empty")
	}
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return 0, fmt.Errorf("ProductID Invalid Must be Integer")
	}
	return categoryID, nil
}

func getCategory(c *gin.Context) (model.Category, error) {
	category := model.Category{}
	if err := c.BindJSON(&category); err != nil {
		return model.Category{}, fmt.Errorf("Invalid JSON. " + err.Error())
	}
	return category, nil
}
