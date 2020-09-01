package repository

import (
	"biller-api/src/api/database"
	model "biller-api/src/api/model"
	"fmt"
)

type Category struct{}

func (c *Category) GetByIdProduct(categoryID int) interface{} {
	categorys := []model.Category{}
	err := database.DB.Select(&categorys, "SELECT c.ID, c.Title, c.CategoryID FROM category AS c INNER JOIN product-category AS pc ON p.ProductID = %d AND c.ID = pt.CategoryID")
	if err != nil {
		fmt.Println(fmt.Sprintf("error running server %s", err.Error()))
		return nil
	}
	return categorys
}

func (c *Category) Get(categoryID int) model.Category {
	category := model.Category{}
	err := database.DB.Get(&category, fmt.Sprintf("SELECT c.ID, c.Title, c.CategoryID FROM category AS c WHERE c.ID=%d", categoryID))
	if err != nil {
		fmt.Println(fmt.Sprintf("error running server %s", err.Error()))
		return model.Category{}
	}
	return category
}

func (c *Category) Update(category model.Category) interface{} {
	query := fmt.Sprintf("UPDATE category AS c SET c.Title = '%s', c.CategoryID = %d WHERE c.ID = %d",
		category.Title,
		category.CategoryID,
		category.ID,
	)
	result := database.DB.MustExec(query)
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		fmt.Println("error in exec query")
		return nil
	}
	return c.Get(category.ID)
}

func (c *Category) Delete(categoryID int) interface{} {
	query := fmt.Sprintf("DELETE FROM category AS c WHERE c.ID = %d",
		categoryID,
	)
	result := database.DB.MustExec(query)
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		fmt.Println("error in exec query")
		return false
	}
	return true
}

func (c *Category) Create(category model.Category) model.Category {
	query := fmt.Sprintf("INSERT INTO category VALUES(0, '%s', %d)",
		category.Title,
		category.CategoryID,
	)
	result := database.DB.MustExec(query)
	lastId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("error in exec query")
		return model.Category{}
	}
	return c.Get(int(lastId))
}
