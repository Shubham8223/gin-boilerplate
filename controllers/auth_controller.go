package controllers

import (
	"net/http"

	"gin-boilerplate/schemas"
	"gin-boilerplate/services"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var authUser schemas.AuthUser

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
