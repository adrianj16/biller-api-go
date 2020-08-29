package repository

import (
	"biller-api/src/api/database"
	model "biller-api/src/api/model/product"
	"fmt"
)

type Product struct{}

func (p *Product) GetAll() interface{} {
	products := []model.Product{}
	err := database.DB.Select(&products, "SELECT ID, Title, Description, (New = b'1') AS New, (MultiProduct = b'1') AS MultiProduct FROM product")
	if err != nil {
		fmt.Println(fmt.Sprintf("error running server %s", err.Error()))
		return nil
	}
	return products
}
