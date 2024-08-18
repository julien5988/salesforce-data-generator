package routes

import (
	"net/http"
	"salesforce-data-generator/generator"
	"salesforce-data-generator/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleGenerateContacts(c *gin.Context) {
	countStr := c.PostForm("count")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid count parameter"})
		return
	}

	contacts, err := generator.GenerateContact(count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating contacts"})
		return
	}

	filePath := "./static/contacts.csv"
	err = utils.WriteCSV(filePath, contacts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing CSV file"})
		return
	}
	c.Header("Content-Disposition", "attachment; filename=contacts.csv")
	c.Header("Content-Type", "text/csv")
	c.File(filePath)
}
