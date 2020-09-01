package usecase

import (
	"biller-api/src/api/model"
	"biller-api/src/api/repository"
)

type Tag struct{}

var tagRepository repository.Tag

func (t *Tag) Update(tag model.Tag) interface{} {
	return tagRepository.Update(tag)
}

func (t *Tag) Delete(tagID int) interface{} {
	return tagRepository.Delete(tagID)
}

func (t *Tag) Create(tag model.Tag) interface{} {
	return tagRepository.Create(tag)
}
