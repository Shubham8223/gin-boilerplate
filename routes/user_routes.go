package routes

import (
	"gin-boilerplate/controllers"
	"gin-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"admin"}), controllers.CreateUser)
		userRoutes.GET("/:id",middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"admin","user"}), controllers.GetUserByID)
		userRoutes.PUT("/:id", middlewares.AuthMiddleware(),middlewares.RbacMiddleware([]string{"admin"}), controllers.UpdateUser)
		userRoutes.DELETE("/:id", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"admin"}), controllers.DeleteUser)
	}
}
