package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	User      User    `gorm:"constraint:OnDelete:CASCADE;"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"constraint:OnDelete:CASCADE;"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}
