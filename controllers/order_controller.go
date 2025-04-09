package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-boilerplate/config"
	"gin-boilerplate/schemas"
	"gin-boilerplate/services"
)

func CreateOrder(c *gin.Context) {
	var order schemas.CreateOrderInput
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.UserID = userID.(uint)
	product,product_error := services.GetProductByID(order.ProductID)
	if product_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": product_error.Error()})
		return
	}
	order.Total = product.Price * float64(order.Quantity)
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		createdOrder, err := services.CreateOrder(tx, &order)
		if err != nil {
			return err
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": fmt.Sprintf("Order %d created successfully", createdOrder.ID),
		})
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func GetOrderByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := services.GetOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func GetOrderByUserId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := services.GetOrderByUserId(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found for this user"})
		return
	}
	c.JSON(http.StatusOK, order)
}


func UpdateOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updates schemas.UpdateOrderInput

	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateOrder(uint(id), &updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

func DeleteOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := services.DeleteOrder(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
