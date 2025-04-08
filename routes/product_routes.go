package routes

import (
	"gin-boilerplate/controllers"
	"gin-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup) {
	productRoutes := r.Group("/products")
	{
		productRoutes.POST("/", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"admin"}), controllers.CreateProduct)
		productRoutes.GET("/", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"user","admin"}),controllers.GetAllProducts)
		productRoutes.GET("/:id", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"user","admin"}),controllers.GetProductByID)
		productRoutes.PUT("/:id", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"admin"}), controllers.UpdateProduct)
		productRoutes.DELETE("/:id", middlewares.AuthMiddleware(), middlewares.RbacMiddleware([]string{"admin"}), controllers.DeleteProduct)
	}
}
