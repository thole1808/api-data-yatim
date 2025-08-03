package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// func APIKeyMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		apiKey := c.GetHeader("API_KEY")
// 		if apiKey != os.Getenv("API_KEY") {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }

func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("API_KEY")
		fmt.Println("DEBUG Header API_KEY:", apiKey)
		fmt.Println("DEBUG ENV API_KEY:", os.Getenv("API_KEY"))

		if apiKey == "" || apiKey != os.Getenv("API_KEY") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API Key"})
			c.Abort()
			return
		}
		c.Next()
	}
}
