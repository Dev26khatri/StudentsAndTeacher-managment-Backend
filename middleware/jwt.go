package middleware

import (
	"GOGIN/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//It see the Authorization section in client so first check if Your token is not into authorization then hiting the warning of unathorized
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization Header is missing",
			})
			c.Abort()
			return
		}

		//Fine you enter your token at perfect section but you forgot to adding the Bearer keyword at starting of the token then it will hiting the warning 
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization formate ex.(space in Token and Bearer key word )",
			})
			c.Abort()
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Set("emails", claims.Email )

		c.Next()

	}

}
