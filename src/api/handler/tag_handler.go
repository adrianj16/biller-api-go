package handler

import (
	"biller-api/src/api/model"
	"biller-api/src/api/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var tagUseCase usecase.Tag

func TagUpdate(c *gin.Context) {
	tagID, err := getTagID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	tag, err := getTag(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	tag.ID = tagID

	c.JSON(http.StatusOK, tagUseCase.Update(tag))
}

func TagDelete(c *gin.Context) {
	tagID, err := getTagID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, tagUseCase.Delete(tagID))
}

func TagCreate(c *gin.Context) {
	tag, err := getTag(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, tagUseCase.Create(tag))
}

func getTagID(c *gin.Context) (int, error) {
	tagIDStr := c.Param("ID")
	if tagIDStr == "" {
		return 0, fmt.Errorf("ProductID Empty")
	}
	tagID, err := strconv.Atoi(tagIDStr)
	if err != nil {
		return 0, fmt.Errorf("ProductID Invalid Must be Integer")
	}
	return tagID, nil
}

func getTag(c *gin.Context) (model.Tag, error) {
	tag := model.Tag{}
	if err := c.BindJSON(&tag); err != nil {
		return model.Tag{}, fmt.Errorf("Invalid JSON. " + err.Error())
	}
	return tag, nil
}
