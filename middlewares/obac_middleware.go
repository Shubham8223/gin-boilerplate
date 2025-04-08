package middlewares

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"gin-boilerplate/config"
	"gin-boilerplate/enums"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ObacMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDVal, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		userID := userIDVal.(uint)

		pathParts := strings.Split(strings.Trim(c.Request.URL.Path, "/"), "/")
		if len(pathParts) < 2 {
			c.Abort()
			return
		}

		resource := pathParts[len(pathParts)-2]
		resourceIDStr := pathParts[len(pathParts)-1]
		resourceID, err := strconv.Atoi(resourceIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid resource ID"})
			c.Abort()
			return
		}

		configEntry, ok := enums.ResourceMap[resource]
		if !ok {
			c.Abort()
			return
		}

		modelPtr := reflect.New(reflect.TypeOf(configEntry.Model).Elem()).Interface()
		err = config.DB.First(modelPtr, resourceID).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": resource + " not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
			}
			c.Abort()
			return
		}

		val := reflect.ValueOf(modelPtr).Elem().FieldByName(configEntry.OwnerKey)
		if !val.IsValid() || val.Uint() != uint64(userID) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied to this " + resource})
			c.Abort()
			return
		}

		c.Next()
	}
}
