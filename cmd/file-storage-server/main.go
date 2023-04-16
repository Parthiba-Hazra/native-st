package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Parthiba-Hazra/restapi"
	"github.com/Parthiba-Hazra/restapi/operations"
	"github.com/Parthiba-Hazra/storage"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/loads"
	"github.com/go-swagger/go-swagger/examples/oauth2/restapi"
)

func main() {

	ceph, err := storage.NewCephStorage()
	if err != nil {
		log.Fatalf("Error creating ceph storage: %s", err)
	}
	defer ceph.Close()

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalf("Error loading Swagger spec: %s, err")
	}

	apiHandler := operations.NewFileStorageAPI(swaggerSpec)

	apiHan
}

func uploadHandler(c *gin.Context) {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to get file: %s", err),
		})
		return
	}
	defer file.Close()

	// UPload file to Ceph
	err = models.UploadFile(file, header.Filename)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to upload file: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded sucessfully",
	})
}

func downloadHandler(c *gin.Context) {
	filename := c.Query("filename")
	if filename == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Missing 'filename' query parameter",
		})
		return
	}

	fileCounts, err := models.DownloadFile(filename)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to download file: %s", err),
		})
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	c.Data(http.StatusOK, "application/delete-stream", fileCounts)
}

func deleteHandler(c *gin.Context) {
	filename := c.PostForm("filename")
	if filename == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Missing 'filename' from field",
		})
		return
	}

	err := models.DeleteFile(filename)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to delete file: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "File deleted sucessfully",
	})
}
