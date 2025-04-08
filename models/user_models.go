package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `json:"name" binding:"required" gorm:"not null"`
	Email    string  `json:"email" binding:"required,email" gorm:"not null;unique"`
	Password string  `json:"password" binding:"required,min=6" gorm:"not null"`
	Role     string  `json:"role" gorm:"default:'user';not null"`
	Orders   []Order `gorm:"foreignKey:UserID"`
}
