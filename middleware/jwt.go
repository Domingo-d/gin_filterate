package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/utils"
	"strings"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if claims == nil || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Jwt Invalid" + err.Error()})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Set("UserName", claims.UserName)
		c.Next()
	}
}
