package usecase

import (
	"biller-api/src/api/model"
	"biller-api/src/api/repository"
	"biller-api/src/api/util"
	"fmt"
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
	imgs := []model.Image{}
	if product.Img != nil && len(*product.Img) > 0 {
		for _, img := range *product.Img {
			path, err := util.GenerateImage(img.Base64)
			if err != nil {
				return fmt.Errorf("Error generate imagen: " + err.Error())
			}
			img.Path = path
			imgResp := imageRepository.Create(img)
			imgs = append(imgs, imgResp)
		}
	}

	tags := []model.Tag{}
	if product.Tags != nil && len(*product.Tags) > 0 {
		for _, tag := range *product.Tags {
			tagResp := tagRepository.Create(tag)
			tags = append(tags, tagResp)
		}
	}

	categories := []model.Category{}
	if product.Categories != nil && len(*product.Categories) > 0 {
		for _, category := range *product.Categories {
			categoryResp := categoryRepository.Create(category)
			categories = append(categories, categoryResp)
		}
	}

	productResp := productRepository.Create(product)
	productResp.Img = &imgs
	productResp.Tags = &tags
	productResp.Categories = &categories
	return productResp
}
