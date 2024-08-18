package generator

import (
	"math/rand"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
)

// Contact représente un contact fictif
type Contact struct {
	FirstName   string `faker:"first_name"`
	LastName    string `faker:"last_name"`
	Phone       string `faker:"phone_number"`
	Email       string
	RealAddress faker.RealAddress `faker:"real_address"`
}

// GenerateContact génère un nombre donné de contacts fictifs
func GenerateContact(count int) ([]Contact, error) {
	rand.Seed(time.Now().UnixNano())
	var contacts []Contact

	// Liste des domaines d'email
	domains := []string{"gmail.com", "hotmail.com", "yahoo.com", "outlook.com"}

	for i := 0; i < count; i++ {
		var contact Contact
		if err := faker.FakeData(&contact); err != nil {
			return nil, err // Retourner l'erreur si la génération échoue
		}

		// Construire l'adresse e-mail personnalisée
		emailDomain := domains[rand.Intn(len(domains))]
		contact.Email = strings.ToLower(contact.FirstName + "." + contact.LastName + "@" + emailDomain)

		contacts = append(contacts, contact) // Ajouter le contact à la liste
	}

	return contacts, nil
}
