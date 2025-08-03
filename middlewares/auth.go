package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
// 			c.Abort()
// 			return
// 		}

// 		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			return []byte(os.Getenv("JWT_SECRET")), nil
// 		})

// 		if err != nil || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		fmt.Println("DEBUG Authorization Header:", authHeader)
// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
// 			c.Abort()
// 			return
// 		}

// 		if !strings.HasPrefix(authHeader, "Bearer ") {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization format must be Bearer <token>"})
// 			c.Abort()
// 			return
// 		}

// 		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			return []byte(os.Getenv("JWT_SECRET")), nil
// 		})
// 		if err != nil || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		fmt.Println("DEBUG Authorization Header:", authHeader)

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := authHeader
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
