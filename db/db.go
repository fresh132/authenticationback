package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Не удалось подключиться к базе данных", err)
	}

	return db
}
