package usecase

import (
	"biller-api/src/api/repository"
)

type Product struct{}

func (p *Product) GetAll() interface{} {
	productRepository := repository.Product{}
	return productRepository.GetAll()
}
