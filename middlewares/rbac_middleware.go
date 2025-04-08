package middlewares

import (
	"net/http"

	"gin-boilerplate/utils"

	"github.com/gin-gonic/gin"
)

func RbacMiddleware(roleAllowed []string) gin.HandlerFunc {

	return func(c *gin.Context) {

		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found"})
			c.Abort()
			return
		}
		roleStr, ok := role.(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role"})
			c.Abort()
			return
		}

		hasRolePermission := utils.SearchQuerySlice(roleStr, roleAllowed)
		if !hasRolePermission {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
