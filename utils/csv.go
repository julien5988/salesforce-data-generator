package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"salesforce-data-generator/generator"
)

// WriteCSV écrit les données dans un fichier CSV avec des en-têtes
func WriteCSV(filepath string, data interface{}) error {
	// Ouvrir le fichier en écriture
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Créer un writer CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Définir les en-têtes
	var headers []string
	switch data.(type) {
	case []generator.Account:
		headers = []string{"ID", "Name", "Phone", "Website", "Industry", "Description"}
	case []generator.Contact:
		headers = []string{"FirstName", "LastName", "Phone", "Email", "Address", "City", "State", "PostalCode"}
	case []generator.Lead:
		headers = []string{"FirstName", "LastName", "Phone", "Email", "Source"}
	default:
		return fmt.Errorf("unsupported data type")
	}

	// Écrire les en-têtes
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Convertir les données en lignes CSV
	var records [][]string
	switch v := data.(type) {
	case []generator.Account:
		for _, account := range v {
			records = append(records, []string{
				account.ID, // Inclure l'ID ici
				account.Name,
				account.Phone,
				account.Website,
				account.Industry,
				account.Description,
			})
		}
	case []generator.Contact:
		for _, contact := range v {
			address := contact.RealAddress.Address
			city := contact.RealAddress.City
			state := contact.RealAddress.State
			postalCode := contact.RealAddress.PostalCode

			records = append(records, []string{
				contact.FirstName,
				contact.LastName,
				contact.Phone,
				contact.Email,
				address,
				city,
				state,
				postalCode,
			})
		}
	case []generator.Lead:
		for _, lead := range v {
			records = append(records, []string{
				lead.FirstName,
				lead.LastName,
				lead.Email,
				lead.Phone,
				lead.Source,
			})

		}

	default:
		return fmt.Errorf("unsupported data type")
	}

	// Écrire les lignes de données
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
