package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    uint    `json:"user_id" binding:"required" gorm:"not null"`
	User      User    `gorm:"constraint:OnDelete:CASCADE;"`
	ProductID uint    `json:"product_id" binding:"required" gorm:"not null"`
	Product   Product `gorm:"constraint:OnDelete:CASCADE;"`
	Quantity  int     `json:"quantity" binding:"required,gt=0" gorm:"not null"`
	Total     float64 `json:"total" binding:"required,gt=0" gorm:"not null"`
}
