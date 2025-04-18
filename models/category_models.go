package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string    `json:"name" gorm:"not null" binding:"required"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}
