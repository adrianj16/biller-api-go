package repository

import (
	"biller-api/src/api/database"
	model "biller-api/src/api/model"
	"fmt"
)

type Tag struct{}

func (t *Tag) GetByIdProduct(tagID int) interface{} {
	tags := []model.Tag{}
	err := database.DB.Select(&tags, "SELECT t.ID, t.Title, t.TagID FROM tag AS t INNER JOIN product-tag AS pt ON p.ProductID = %d AND t.ID = pt.TagID")
	if err != nil {
		fmt.Println(fmt.Sprintf("error running server %s", err.Error()))
		return nil
	}
	return tags
}

func (t *Tag) Get(tagID int) model.Tag {
	tag := model.Tag{}
	err := database.DB.Get(&tag, fmt.Sprintf("SELECT t.ID, t.Title, t.TagID FROM tag AS t WHERE t.ID=%d", tagID))
	if err != nil {
		fmt.Println(fmt.Sprintf("error running server %s", err.Error()))
		return model.Tag{}
	}
	return tag
}

func (t *Tag) Update(tag model.Tag) interface{} {
	query := fmt.Sprintf("UPDATE tag AS t SET t.Title = '%s', t.TagID = %d WHERE t.ID = %d",
		tag.Title,
		tag.TagID,
		tag.ID,
	)
	result := database.DB.MustExec(query)
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		fmt.Println("error in exec query")
		return nil
	}
	return t.Get(tag.ID)
}

func (t *Tag) Delete(TagID int) interface{} {
	query := fmt.Sprintf("DELETE FROM tag AS t WHERE t.ID = %d",
		TagID,
	)
	result := database.DB.MustExec(query)
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		fmt.Println("error in exec query")
		return false
	}
	return true
}

func (t *Tag) Create(tag model.Tag) model.Tag {
	query := fmt.Sprintf("INSERT INTO tag VALUES(0, '%s', %d)",
		tag.Title,
		tag.TagID,
	)
	result := database.DB.MustExec(query)
	lastId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("error in exec query")
		return model.Tag{}
	}
	return t.Get(int(lastId))
}
