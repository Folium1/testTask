package delivery

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"test/models"
	"time"

	"github.com/gin-gonic/gin"
)

// getImagesHandler retrieves all images from the database.
func getImagesHandler(c *gin.Context) {
	
	// Retrieve userID from context
	userID, ok := c.Get("userID")
	if !ok {
		handleError(c, http.StatusInternalServerError, "Failed to retrieve userID", nil)
		return
	}

	// Retrieve images from the database
	images, err := Storage.GetImagesByUserID(userID.(int))
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to retrieve images", err)
		return
	}
	// Return the images
	c.JSON(http.StatusOK, images)
}

// postImagesHandler saves the image to the local folder
func saveImageToFile(file *multipart.FileHeader, filePath string) error {
	fileData, err := file.Open()
	if err != nil {
		return err
	}
	defer fileData.Close()

	// Create a buffer to store the file content
	buffer := &bytes.Buffer{}

	// Copy the file content to the buffer
	_, err = io.Copy(buffer, fileData)
	if err != nil {
		return err
	}

	// Create the directory if it doesn't exist
	err = os.MkdirAll(filepath.Dir(filePath), 0777)
	if err != nil {
		return err
	}

	// Save the file to the local folder
	err = os.WriteFile(filePath, buffer.Bytes(), 0666)
	if err != nil {
		return err
	}
	// If everything went well, return nil
	return nil
}

// postImagesHandler saves the image to the local folder and saves the image's data to the database.
func postImagesHandler(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		handleError(c, http.StatusUnauthorized, "Couldn't get user's ID", nil)
		return
	}
	// Retrieve the file from the request
	file, err := c.FormFile("image")
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to get image from request", err)
		return
	}
	// Create the file path
	filePath := filepath.Join("uploads", fmt.Sprintf("%v/%v/%v/", time.Now().Year(), time.Now().Month().String(), time.Now().Day()), file.Filename)
	// Create the image data
	imageData := models.Image{UserID: userID.(int), Path: filePath, URL: c.Request.URL.String()}

	// Save image to the database
	if err := Storage.SaveImage(imageData); err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to save image", err)
		return
	}

	// Save the file to the local folder
	if err := saveImageToFile(file, filePath); err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to save image file", err)
		return
	}
	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
