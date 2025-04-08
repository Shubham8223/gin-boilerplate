package controllers

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"gin-boilerplate/config"
	"gin-boilerplate/schemas"
	"gin-boilerplate/services"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category schemas.CreateCategoryInput

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		createdCategory, err := services.CreateCategory(tx, &category)
		if err != nil {
			return err
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":       "Category created successfully",
			"category_id":   createdCategory.ID,
			"category_name": createdCategory.Name,
		})
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func GetAllCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No categories found"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := services.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var category schemas.UpdateCategoryInput

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateCategory(uint(id), &category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := services.DeleteCategory(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
