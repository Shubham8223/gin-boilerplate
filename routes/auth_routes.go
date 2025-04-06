package routes

import (
	"github.com/gin-gonic/gin"
	"gin-boilerplate/controllers"
)

func AuthUserRoutes(r *gin.RouterGroup) {
	userRoutes := r.Group("/auth")
	{
		userRoutes.POST("/login",controllers.UserLogin)
	}
}
