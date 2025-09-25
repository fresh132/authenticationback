package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `json:"id" gorm:"unique"`
	Mail     string `json:"mail" gorm:"unique"`
	Password string `json:"password"`
}
