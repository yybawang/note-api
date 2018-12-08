package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbr *gorm.DB
var err error

func init() {
	dbr, err = connect()
}

func connect() (db *gorm.DB, err2 error) {
	db2, err2 := gorm.Open("mysql", "root:123456@tcp(118.190.174.119:3307)/note?charset=utf8mb4&parseTime=True&loc=Local")
	return db2, err2
}

func Connect() (db *gorm.DB, err2 error) {
	return dbr, err
}
