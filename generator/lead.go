package generator

import (
	"math/rand"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
)

// Lead représente un lead avec des informations pertinentes
type Lead struct {
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
	Phone     string `faker:"phone_number"`
	Email     string
	Source    string
	CreatedAt time.Time `faker:"-"`
}

// Liste des sources possibles pour un lead
var sources = []string{
	"Web Form",
	"Social Media",
	"Referral",
	"Event",
	"Email Campaign",
	"Purchased List",
}

// GenerateLeads génère un nombre donné de leads fictifs
func GenerateLeads(count int) ([]Lead, error) {
	rand.Seed(time.Now().UnixNano())
	var leads []Lead

	// Liste des domaines d'email
	domains := []string{"gmail.com", "hotmail.com", "yahoo.com", "outlook.com"}

	for i := 0; i < count; i++ {
		var lead Lead
		if err := faker.FakeData(&lead); err != nil {
			return nil, err
		}

		// Construire l'adresse e-mail personnalisée
		emailDomain := domains[rand.Intn(len(domains))]
		lead.Email = strings.ToLower(lead.FirstName + "." + lead.LastName + "@" + emailDomain)

		// Assigner une source aléatoire
		lead.Source = sources[rand.Intn(len(sources))]
		// Assigner la date actuelle pour CreatedAt
		lead.CreatedAt = time.Now()

		leads = append(leads, lead) // Ajouter le lead à la liste
	}

	return leads, nil
}
