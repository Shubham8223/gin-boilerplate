package services

import (
	"errors"

	"gin-boilerplate/config"
	"gin-boilerplate/models"
	"gin-boilerplate/schemas"
	"gin-boilerplate/utils"

	"gorm.io/gorm"
)

func CreateOrder(tx *gorm.DB, input *schemas.CreateOrderInput) (*models.Order, error) {
	order := &models.Order{}
	utils.StructToMapCreate(input, order)
	if err := tx.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func GetOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order
	if err := config.DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &order, nil
}

func GetOrderByUserId(userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := config.DB.Preload("User").Preload("Product").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, errors.New("orders for userId not found")
	}
	return orders, nil
}

func UpdateOrder(orderID uint, updatedData *schemas.UpdateOrderInput) error {
	updates := utils.StructToMapUpdate(updatedData)
	if len(updates) == 0 {
		return nil 
	}
	product,product_error := GetProductByID(updates["product_id"].(uint)) 
	if product_error != nil {
		return product_error				
	}	
	updates["total"] = product.Price * float64(updates["quantity"].(int))
	return config.DB.Model(&models.Order{}).Where("id = ?", orderID).Updates(updates).Error
}

func DeleteOrder(orderID uint) error {
	return config.DB.Delete(&models.Order{}, orderID).Error
}
