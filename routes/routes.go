package routes

import (
	"github.com/gin-gonic/gin"
)

func serveIndex(c *gin.Context) {
	c.File("./static/index.html")
}

// SetupRoutes configures all the routes for the application
func SetupRoutes(r *gin.Engine) {
	// Serve the index.html file
	r.GET("/", serveIndex)

	// Handle the form submission and generate the CSV file
	r.POST("/generate-account", handleGenerateAccounts)
	r.POST("/generate-contact", handleGenerateContacts)
}
