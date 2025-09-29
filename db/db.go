package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=2202 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Krasnoyarsk"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Не удалось подключиться к базе данных", err)
	}

	return db
}
