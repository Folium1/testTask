package delivery

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware is a middleware to check if user is authorized. 
// It checks if the token is valid and passes userID to context.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// retrieve token from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		// retrieve token string
		tokenString := strings.Split(authHeader, " ")[1]
		// parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// pass userID to context
		claims := token.Claims.(jwt.MapClaims)
		userID := int(claims["userID"].(float64))
		c.Set("userID", userID)
		c.Next()
	}
}
