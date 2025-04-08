package services

import (
	"errors"

	"gin-boilerplate/config"
	"gin-boilerplate/models"
	"gin-boilerplate/schemas"
	"gin-boilerplate/utils"

	"gorm.io/gorm"
)

func CreateCategory(tx *gorm.DB, input *schemas.CreateCategoryInput) (*models.Category, error) {
	category := &models.Category{}
	utils.StructToMapCreate(input, category)

	if err := tx.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return nil, errors.New("no categories found")
	}
	return categories, nil
}

func GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := config.DB.Preload("Products").First(&category, id).Error; err != nil {
		return nil, errors.New("category not found")
	}
	return &category, nil
}

func UpdateCategory(categoryID uint, updatedData *schemas.UpdateCategoryInput) error {
	updates := utils.StructToMapUpdate(updatedData)
	if len(updates) == 0 {
		return nil 
	}
	return config.DB.Model(&models.Category{}).Where("id = ?", categoryID).Updates(updates).Error
}

func DeleteCategory(categoryID uint) error {
	return config.DB.Delete(&models.Category{}, categoryID).Error
}
