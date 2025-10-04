package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `json:"id" gorm:"unique"`
	Mail       string    `json:"mail" gorm:"unique"`
	Password   string    `json:"password"`
	CreateTime time.Time `json:"createtime"`
}

type PasswordChangeRequest struct {
	NewPassword string `json:"newpassword" binding:"required,min=8"`
}

type PasswordMailRequest struct {
	Mail     string `json:"mail" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
