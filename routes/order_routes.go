package routes

import (
	"gin-boilerplate/controllers"
	"gin-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.RouterGroup) {
	orderRoutes := r.Group("/orders")
	{
		orderRoutes.POST("/", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"user"}),controllers.CreateOrder)
		orderRoutes.GET("/:id", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"user"}),middlewares.ObacMiddleware(), controllers.GetOrderByID)
		orderRoutes.PUT("/:id", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"user"}),middlewares.ObacMiddleware(), controllers.UpdateUser)
		orderRoutes.DELETE("/:id", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"user"}),middlewares.ObacMiddleware(), controllers.DeleteUser)
	}
}
