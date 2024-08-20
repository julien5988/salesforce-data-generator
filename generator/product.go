package generator

import (
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
)

// Product représente un produit avec des champs courants
type Product struct {
	Name        string  `faker:"word"`
	Description string  `faker:"sentence"`
	Price       float64 `faker:"-"` // On va générer le prix manuellement
	Category    string  `faker:"word"`
}

// GenerateProducts génère un nombre donné de produits fictifs
func GenerateProducts(count int) ([]Product, error) {
	rand.Seed(time.Now().UnixNano())
	var products []Product

	// Liste des catégories de produits
	categories := []string{"Electronics", "Furniture", "Clothing", "Books", "Toys", "Sports"}

	for i := 0; i < count; i++ {
		var product Product
		if err := faker.FakeData(&product); err != nil {
			return nil, err
		}

		// Attribuer un prix aléatoire entre 10 et 1000
		product.Price = 10 + rand.Float64()*(1000-10)

		// Attribuer une catégorie aléatoire
		product.Category = categories[rand.Intn(len(categories))]

		products = append(products, product) // Ajouter le produit à la liste
	}

	return products, nil
}
