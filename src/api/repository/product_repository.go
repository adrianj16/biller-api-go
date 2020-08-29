package repository

import (
	"biller-api/src/api/database"
	model "biller-api/src/api/model"
	"biller-api/src/api/util"
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

func (p *Product) Get(productID int) interface{} {
	product := model.Product{}
	err := database.DB.Get(&product, fmt.Sprintf("SELECT ID, Title, Description, (New = b'1') AS New, (MultiProduct = b'1') AS MultiProduct FROM product WHERE ID=%v", productID))
	if err != nil {
		fmt.Println(fmt.Sprintf("error running server %s", err.Error()))
		return nil
	}
	return product
}

func (p *Product) Update(product model.Product) interface{} {
	query := fmt.Sprintf("UPDATE product SET Title = '%s', Description = '%s', New = %d, MultiProduct = %d WHERE ID = %d",
		product.Title,
		product.Description,
		util.BoolToInt(product.New),
		util.BoolToInt(product.Multiproduct),
		product.ID,
	)
	result := database.DB.MustExec(query)
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		fmt.Println("error in exec query")
		return nil
	}
	return p.Get(product.ID)
}

func (p *Product) Delete(productID int) interface{} {
	query := fmt.Sprintf("DELETE FROM product WHERE ID = %d",
		productID,
	)
	result := database.DB.MustExec(query)
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		fmt.Println("error in exec query")
		return false
	}
	return true
}

func (p *Product) Create(product model.Product) interface{} {
	query := fmt.Sprintf("INSERT INTO product(Title, Description, New, MultiProduct) VALUES('%s', '%s', %d, %d)",
		product.Title,
		product.Description,
		util.BoolToInt(product.New),
		util.BoolToInt(product.Multiproduct),
	)
	result := database.DB.MustExec(query)
	lastId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("error in exec query")
		return nil
	}
	return p.Get(int(lastId))
}
