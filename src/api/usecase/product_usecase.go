package usecase

import (
	"biller-api/src/api/model"
	"biller-api/src/api/repository"
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"os"
	"time"
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
			path, err := generateImage(img.Base64)
			if err != nil {
				return fmt.Errorf("Error generate imagen: " + err.Error())
			}
			img.Path = path
			imgResp := imageRepository.Create(img)
			imgs = append(imgs, imgResp)
		}
	}
	productResp := productRepository.Create(product)
	productResp.Img = &imgs
	return productResp
}

func generateImage(b string) (string, error) {
	unbased, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		panic("Cannot decode b64")
	}

	r := bytes.NewReader(unbased)
	im, err := png.Decode(r)
	if err != nil {
		return "", fmt.Errorf("Error base64 reader: " + err.Error())
	}

	path := time.Now().Format(time.RFC3339) + ".png"

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return "", fmt.Errorf("Error open file: " + err.Error())
	}

	err = png.Encode(f, im)
	if err != nil {
		return "", fmt.Errorf("Error create file: " + err.Error())
	}
	return path, nil
}
