package usecase

import (
	"biller-api/src/api/model"
	"biller-api/src/api/repository"
)

type Category struct{}

var categoryRepository repository.Category

func (t *Category) Update(category model.Category) interface{} {
	return categoryRepository.Update(category)
}

func (t *Category) Delete(categoryID int) interface{} {
	return categoryRepository.Delete(categoryID)
}

func (t *Category) Create(category model.Category) interface{} {
	return categoryRepository.Create(category)
}
