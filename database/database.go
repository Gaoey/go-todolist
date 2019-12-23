package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open("mysql", "root:password@/todolist?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
