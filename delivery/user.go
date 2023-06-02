package delivery

import (
	"log"
	"net/http"
	"test/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// loginHandler handles login requests.
func loginHandler(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Retrieve JSON body from request
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if user exists
	user, err := Storage.GetUserByUsername(request.Username)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if password is correct
	if request.Password != user.Password {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}
	// hash password
	hashedPassword, err := hashPassword(request.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	userLogin := models.User{
		Username: request.Username,
		Password: string(hashedPassword),
	}
	// Save user log
	err = Storage.SaveUserLog(userLogin)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user log"})
		return
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(12 * time.Hour).Unix(),
	})

	// Generate encoded token
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// Set token in header
	c.Writer.Header().Set("Authorization", "Bearer "+tokenString)
	// Return token
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// hashPassword hashes the given password
func hashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Println("Couldn't hash user's password", err)
		return []byte(""), err
	}
	return hashedPassword, nil
}
