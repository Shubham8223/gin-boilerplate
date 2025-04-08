package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string    `json:"name" binding:"required" gorm:"not null"`
	Price      float64   `json:"price" binding:"required" gorm:"not null"`
	CategoryID uint      `json:"category_id" binding:"required" gorm:"not null"`
	Category   Category  `gorm:"constraint:OnDelete:CASCADE;"`
	Orders     []Order   `gorm:"foreignKey:ProductID"`
}
