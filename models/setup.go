package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	conn := "root:root@tcp(127.0.0.1:3306)/churchdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
