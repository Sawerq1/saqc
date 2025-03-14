package repo

import (
	"fmt"
	"math/rand"
	"sportsstore/models"
)

func (mr *MemoryRepo) Seed() {
	mr.categories = make([]models.Category, 3)
	for i := 0; i < 3; i++ {
		catName := fmt.Sprintf("Category_%v", i+1)
		mr.categories[i] = models.Category{ID: i + 1, CategoryName: catName}
	}
	for i := 0; i < 20; i++ {
		name := fmt.Sprintf("Product_%v", i+1)
		price := rand.Float64() * float64(rand.Intn(500))
		cat := &mr.categories[rand.Intn(len(mr.categories))]
		mr.products = append(mr.products, models.Product{
			ID:   i + 1,
			Name: name, Price: price,
			Category:    cat,
			Description: fmt.Sprintf("%v (%v)", name, cat.CategoryName),
		})
	}
}
