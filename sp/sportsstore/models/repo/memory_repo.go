package repo

import (
	"math"
	//"platform/services"
	"sportsstore/models"
)

// func RegisterMemoryRepoService() {
// 	services.AddSingleton(func() models.Repository {
// 		repo := &MemoryRepo{}
// 		repo.Seed()
// 		return repo
// 	})
// }

type MemoryRepo struct {
	products   []models.Product
	categories []models.Category
}

func (mr *MemoryRepo) GetProduct(id int) (product models.Product) {
	for _, p := range mr.products {
		if p.ID == id {
			product = p
			return product
		}
	}
	return
}

func (mr *MemoryRepo) GetProducts() (results []models.Product) {
	return mr.products
}

func (mr *MemoryRepo) GetProductPage(page, pageSize int) ([]models.Product, int) {
	return getPage(mr.products, page, pageSize), len(mr.products)
}

func getPage(src []models.Product, page, pageSize int) []models.Product {
	start := (page - 1) * pageSize
	if page > 0 && len(src) > start {
		end := (int)(math.Min((float64)(len(src)), (float64)(start+pageSize)))
		return src[start:end]
	}
	return []models.Product{}
}

func (mr *MemoryRepo) GetCategories() (result []models.Category) {
	return mr.categories
}

func (mr *MemoryRepo) GetProductPageCategory(category, page, pageSize int) (products []models.Product, totalAvailable int) {
	if category == 0 {
		return mr.GetProductPage(page, pageSize)
	} else {
		filteredProducts := make([]models.Product, 0, len(mr.products))
		for _, p := range mr.products {
			if p.Category.ID == category {
				filteredProducts = append(filteredProducts, p)
			}
		}
		return getPage(filteredProducts, page, pageSize), len(filteredProducts)
	}
}
