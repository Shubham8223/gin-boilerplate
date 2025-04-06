package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-boilerplate/models"
	"gin-boilerplate/services"
)

func UserLogin(c *gin.Context) {
	var authUser models.AuthUser

	if err := c.ShouldBindJSON(&authUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := services.LoginUser(&authUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
	})
}
