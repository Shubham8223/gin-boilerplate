package services

import (
	"errors"

	"gin-boilerplate/config"
	"gin-boilerplate/models"
	"gin-boilerplate/schemas"
	"gin-boilerplate/utils"

	"gorm.io/gorm"
)

func CreateProduct(tx *gorm.DB, input *schemas.CreateProductInput) (*models.Product, error) {
	product := &models.Product{}
	utils.StructToMapCreate(input, product)
	if err := tx.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, errors.New("no products found")
	}
	return products, nil
}

func GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := config.DB.Preload("Orders").Preload("Category").First(&product, id).Error; err != nil {
		return nil, errors.New("product not found")
	}
	return &product, nil
}

func UpdateProduct(id uint, updatedData *schemas.UpdateProductInput) error {
	updates := utils.StructToMapUpdate(updatedData)
	if len(updates) == 0 {
		return nil 
	}
	return config.DB.Model(&models.Product{}).Where("id = ?", id).Updates(updates).Error
}

func DeleteProduct(id uint) error {
	return config.DB.Delete(&models.Product{}, id).Error
}
