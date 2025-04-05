package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gin-boilerplate/utils/jwt"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	   if strings.HasPrefix(token, "Bearer ") {
			token = strings.TrimPrefix(token, "Bearer ")
		}
	   
	   claims, err := jwt.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	   c.Set("userID", claims.UserID)
	   c.Set("role", claims.Role)

		c.Next()
	}
}
