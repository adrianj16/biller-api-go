package repository

import (
	"biller-api/src/api/database"
	model "biller-api/src/api/model"
	"fmt"
)

type Image struct{}

func (i *Image) GetByIdProduct(productID int) interface{} {
	images := []model.Image{}
	err := database.DB.Select(&images, "SELECT i.ID, i.Path, i.Order FROM image AS i INNER JOIN product-image AS pi ON pi.ProductID = %d AND pi.ImageID = i.ID")
	if err != nil {
		fmt.Println(fmt.Sprintf("error running server %s", err.Error()))
		return nil
	}
	return images
}

func (i *Image) Get(imageID int) model.Image {
	image := model.Image{}
	err := database.DB.Get(&image, fmt.Sprintf("SELECT i.ID, i.Path, i.Order FROM image AS i WHERE i.ID=%d", imageID))
	if err != nil {
		fmt.Println(fmt.Sprintf("error running server %s", err.Error()))
		return model.Image{}
	}
	return image
}

func (i *Image) Update(image model.Image) interface{} {
	query := fmt.Sprintf("UPDATE image SET Order = %d WHERE ID = %d",
		image.Order,
		image.ID,
	)
	result := database.DB.MustExec(query)
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		fmt.Println("error in exec query")
		return nil
	}
	return i.Get(image.ID)
}

func (i *Image) Delete(imageID int) interface{} {
	query := fmt.Sprintf("DELETE FROM image WHERE ID = %d",
		imageID,
	)
	result := database.DB.MustExec(query)
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		fmt.Println("error in exec query")
		return false
	}
	return true
}

func (i *Image) Create(image model.Image) model.Image {
	query := fmt.Sprintf("INSERT INTO image VALUES(0, '%s', %d)",
		image.Path,
		image.Order,
	)
	result := database.DB.MustExec(query)
	lastId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("error in exec query")
		return model.Image{}
	}
	return i.Get(int(lastId))
}
