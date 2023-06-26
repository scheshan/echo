package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(connStr string) error {
	ins, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}
	db = ins
	return nil
}
