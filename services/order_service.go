package services

import (
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

func GetOrderByID(orderID uint) (*schemas.ResponseOrderOutput, error) {
	var order schemas.ResponseOrderOutput
	query := `
		SELECT u.name AS user, p.name AS product, o.quantity, o.total
		FROM orders o
		INNER JOIN users u ON u.id = o.user_id
		INNER JOIN products p ON p.id = o.product_id
		WHERE o.id = ?
		`
    if err := config.DB.Raw(query, orderID).Scan(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func GetOrderByUserId(userID uint) (*schemas.ResponseOrderPerUserOutput, error) {
	var rows []schemas.ResponseOrderOutput
	query := `
		SELECT u.name AS user, p.name AS product, o.quantity, o.total
		FROM orders o
		INNER JOIN users u ON u.id = o.user_id
		INNER JOIN products p ON p.id = o.product_id
		WHERE u.id = ?
		`
	  if err := config.DB.Raw(query, userID).Scan(&rows).Error; err != nil {
		return nil, err 
	}

	if len(rows) == 0 {
		return nil, nil 
	}

	output := &schemas.ResponseOrderPerUserOutput{
		User:   rows[0].User,
		Orders: []schemas.ResponseOrderListOutput{},
	}

	for _, row := range rows {
		output.Orders = append(output.Orders, schemas.ResponseOrderListOutput{
			Product:  row.Product,
			Quantity: row.Quantity,
			Total:    row.Total,
		})
	}
	return output, nil
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
