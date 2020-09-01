package handler

import (
	"biller-api/src/api/model"
	"biller-api/src/api/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var imageUseCase usecase.Image

func ImageUpdate(c *gin.Context) {
	imageID, err := getImageID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	image, err := getImage(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	image.ID = imageID

	c.JSON(http.StatusOK, imageUseCase.Update(image))
}

func ImageDelete(c *gin.Context) {
	imageID, err := getImageID(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, imageUseCase.Delete(imageID))
}

func ImageCreate(c *gin.Context) {
	image, err := getImage(c)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, imageUseCase.Create(image))
}

func getImageID(c *gin.Context) (int, error) {
	imageIDStr := c.Param("ID")
	if imageIDStr == "" {
		return 0, fmt.Errorf("ProductID Empty")
	}
	imageID, err := strconv.Atoi(imageIDStr)
	if err != nil {
		return 0, fmt.Errorf("ProductID Invalid Must be Integer")
	}
	return imageID, nil
}

func getImage(c *gin.Context) (model.Image, error) {
	image := model.Image{}
	if err := c.BindJSON(&image); err != nil {
		return model.Image{}, fmt.Errorf("Invalid JSON. " + err.Error())
	}
	return image, nil
}
