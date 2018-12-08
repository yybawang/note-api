package post

import (
	"model"

	"github.com/jinzhu/gorm"
)

type Posts struct {
	gorm.Model
	CateId  int
	Name    string
	Content string
}

/**
分类下记事本列表
*/
func List(cateId int) (list []Posts, err error) {
	db, err := model.Connect()
	if err != nil {
		return
	}
	db.Where("cate_id = ?", cateId).Find(&list)
	return
}

/**
技术文章详细内容
*/
func Detail(postId int) (detail Posts, err error) {
	db, err := model.Connect()
	if err != nil {
		return
	}
	db.Where("id = ?", postId).First(&detail)
	return
}
