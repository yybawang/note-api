package category

import (
	"model"

	"github.com/jinzhu/gorm"
)

type PostCategory struct {
	gorm.Model
	ParentId int
	Name     string
}

func List() (category []PostCategory, err error) {
	db, err := model.Connect()
	if err != nil {
		return
	}
	db.Find(&category)
	return
}
