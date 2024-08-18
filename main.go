package main

import (
	"salesforce-data-generator/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static files from the "static" directory
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")

	// Setup routes
	routes.SetupRoutes(r)

	// Run the server
	r.Run(":8080")
}
