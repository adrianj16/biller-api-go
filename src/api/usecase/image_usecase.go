package usecase

import (
	"biller-api/src/api/model"
	"biller-api/src/api/repository"
)

type Image struct{}

var imageRepository repository.Image

func (i *Image) Update(image model.Image) interface{} {
	return imageRepository.Update(image)
}

func (i *Image) Delete(imageID int) interface{} {
	return imageRepository.Delete(imageID)
}

func (i *Image) Create(image model.Image) interface{} {
	return imageRepository.Create(image)
}
