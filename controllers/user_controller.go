package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"

	"gin-boilerplate/config"
	"gin-boilerplate/models"
	"gin-boilerplate/services"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		createdUser, err := services.CreateUser(tx, &user)
		if err != nil {
			return err
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": fmt.Sprintf("User %d created successfully", createdUser.ID),
		})
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}


func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updates map[string]interface{}

	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateUser(uint(id), updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := services.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
