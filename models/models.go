package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `json:"id" gorm:"unique"`
	Mail       string    `json:"mail" gorm:"unique"`
	Password   string    `json:"password"`
	CreateTime int       `json:"createtime" gorm:"unique"`
}
