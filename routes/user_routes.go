package routes

import (
	"github.com/gin-gonic/gin"
	"gin-boilerplate/controllers"
	"gin-boilerplate/middlewares"
)

func RegisterUserRoutes(r *gin.RouterGroup) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", middlewares.AuthMiddleware(), middlewares.RoleMiddleware([]string{"admin","user"}), controllers.CreateUser)
		userRoutes.GET("/:id",middlewares.AuthMiddleware(), middlewares.RoleMiddleware([]string{"admin","user"}), controllers.GetUserByID)
		userRoutes.PUT("/:id", middlewares.AuthMiddleware(),middlewares.RoleMiddleware([]string{"admin","user"}), controllers.UpdateUser)
		userRoutes.DELETE("/:id", middlewares.AuthMiddleware(), middlewares.RoleMiddleware([]string{"admin","user"}), controllers.DeleteUser)
	}
}
