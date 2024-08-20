package routes

import (
	"net/http"
	"salesforce-data-generator/generator"
	"salesforce-data-generator/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// handleGenerate handles form submission and generates the CSV file
func handleGenerateLeads(c *gin.Context) {
	countStr := c.PostForm("count")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid count parameter"})
		return
	}

	// Generate fake accounts
	accounts, err := generator.GenerateLeads(count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating accounts"})
		return
	}

	// Define the path for the CSV file
	filePath := "./static/lead.csv"
	err = utils.WriteCSV(filePath, accounts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing CSV file"})
		return
	}

	// Serve the CSV file as a downloadable attachment
	c.Header("Content-Disposition", "attachment; filename=lead.csv")
	c.Header("Content-Type", "text/csv")
	c.File(filePath)
}
