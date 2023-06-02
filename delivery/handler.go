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
	r.Static("/uploads", "./uploads")

	r.POST("/login", loginHandler)
	r.GET("/logs", getLogsHandler)

	image := r.Group("/", AuthMiddleware())
	{
		image.GET("/images", getImagesHandler)
		image.POST("/upload-picture", postImagesHandler)
		image.GET("/image/:url")
	}
	http.ListenAndServe(":8080", r)
}

// handleError is a helper function to handle errors.
func handleError(c *gin.Context, statusCode int, errMsg string, err error) {
	log.Println(err) // Log the error for debugging purposes
	c.JSON(statusCode, gin.H{"error": errMsg})
}
