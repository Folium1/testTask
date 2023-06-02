package delivery

import (
	"log"
	"net/http"
	"os"
	"test/entities"

	"github.com/gin-gonic/gin"
)

var (
	jwtSecret = os.Getenv("JWT_SECRET")
	Storage   *entities.Storage
)

// init is called before the main function is executed and is used to initialize Storage variable.
func init() {
	var err error
	Storage, err = entities.New()
	if err != nil {
		log.Fatal(err)
	}
}

// StartServer starts the server.
func StartServer() {
	r := gin.Default()
	r.POST("/login", loginHandler)
	r.Static("/uploads", "./uploads")
	tokenRequeried := r.Group("/", AuthMiddleware())
	{
		tokenRequeried.GET("/images", getImagesHandler)
		tokenRequeried.POST("/upload-picture", postImagesHandler)
		tokenRequeried.GET("/image/:url")
	}
	http.ListenAndServe(":8080", r)
}

// handleError is a helper function to handle errors.
func handleError(c *gin.Context, statusCode int, errMsg string, err error) {
	log.Println(err) // Log the error for debugging purposes
	c.JSON(statusCode, gin.H{"error": errMsg})
}
