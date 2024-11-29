package shared

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := GetTokenFromRequest(c)
		token, err := jwt.ParseWithClaims(tokenStr, &Payload{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token")
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, _ := token.Claims.(*Payload)

		session, exists := Sessions[claims.Session]
		if !exists || session.ExpiryTime.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired session"})
			c.Abort()
			return
		}

		c.Set("user_id", uint(session.Uid))
		fmt.Println("User ID set in context:", session.Uid) // Depuración
		c.Next()
	}
}
