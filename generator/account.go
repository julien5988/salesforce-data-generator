package generator

import (
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
)

// Account représente l'objet Account dans Salesforce
type Account struct {
	Name        string `faker:"-"`
	Phone       string `faker:"phone_number"`
	Website     string `faker:"url"`
	Industry    string `faker:"oneof: Technology,Finance,Healthcare,Manufacturing,Education"`
	Description string `faker:"sentence"`
	ID          string `faker:"uuid_digit"`
}

// Liste de noms d'entreprises
var companyNames = []string{
	"Tech Innovations Inc.",
	"Future Finance Ltd.",
	"HealthPlus Solutions",
	"Manufacture Pro Co.",
	"Education Essentials",
}

// GenerateAccounts génère un nombre donné de comptes fictifs
func GenerateAccounts(count int) ([]Account, error) {
	rand.Seed(time.Now().UnixNano())
	var accounts []Account

	for i := 0; i < count; i++ {
		var account Account
		if err := faker.FakeData(&account); err != nil {
			return nil, err
		}
		// Assigner un nom d'entreprise aléatoire à partir de la liste
		account.Name = companyNames[rand.Intn(len(companyNames))]
		accounts = append(accounts, account)
	}

	return accounts, nil
}
