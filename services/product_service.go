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

func GetProductByID(id uint) (*schemas.ResponseProductOutput, error) {
	var product schemas.ResponseProductOutput
	query := `
	SELECT p.name,p.price,c.name AS category,array_agg(o.id) AS order_ids
	FROM products p
	INNER JOIN orders o ON p.id = o.product_id
	INNER JOIN categories c ON p.category_id = c.id
	WHERE p.id = ?
	GROUP BY p.id, c.name`
	
	if err := config.DB.Raw(query, id).Scan(&product).Error; err != nil {
		return nil, err
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
