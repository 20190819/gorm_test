package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conn() (*gorm.DB, error) {
	dsn := "root:888888@tcp(localhost:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Printf("dsn: %v\n", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
