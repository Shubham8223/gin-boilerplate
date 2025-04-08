package routes

import (
	"gin-boilerplate/controllers"
	"gin-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup) {
	categoryRoutes := r.Group("/categories")
	{
		categoryRoutes.POST("/", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"admin"}), controllers.CreateCategory)
		categoryRoutes.GET("/", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"admin","user"}),controllers.GetAllCategories)
		categoryRoutes.GET("/:id", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"admin","user"}),controllers.GetCategoryByID)
		categoryRoutes.PUT("/:id", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"admin"}), controllers.UpdateCategory)
		categoryRoutes.DELETE("/:id", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"admin"}), controllers.DeleteCategory)
	}
}
