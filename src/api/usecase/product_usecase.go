package usecase

import (
	"biller-api/src/api/model"
	"biller-api/src/api/repository"
)

type Product struct{}

var productRepository repository.Product

func (p *Product) GetAll() interface{} {
	return productRepository.GetAll()
}

func (p *Product) Get(productID int) interface{} {
	return productRepository.Get(productID)
}

func (p *Product) Update(product model.Product) interface{} {
	return productRepository.Update(product)
}

func (p *Product) Delete(productID int) interface{} {
	return productRepository.Delete(productID)
}

func (p *Product) Create(product model.Product) interface{} {
	return productRepository.Create(product)
}
