package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Mail       string `json:"mail" gorm:"unique"`
	Password   string `json:"password"`
	CreateTime int    `json:"createtime" gorm:"unique"`
}
