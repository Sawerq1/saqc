package store

import (
	"math"
	"platform/http/actionresults"
	"platform/http/handling"
	"sportsstore/models"
)

const pageSize = 4

type ProductHabdler struct {
	Repository   models.Repository
	URLGenerator handling.URLGenerator
}
type ProductTemplateContext struct {
	Products         []models.Product
	Page             int
	PageCount        int
	PageNumbers      []int
	PageURLFunc      func(int) string
	SelectedCategory int
	AddToCartURL     string
}

func (handler ProductHabdler) GetProducts(category, page int) actionresults.ActionResult {
	prods, total := handler.Repository.GetProductPageCategory(category, page, pageSize)
	pageCount := int(math.Ceil(float64(total) / float64(pageSize)))
	return actionresults.NewTemplateAction("product_list.html",
		ProductTemplateContext{
			Products:         prods,
			Page:             page,
			PageCount:        pageCount,
			PageNumbers:      handler.generatePageNumbers(pageCount),
			PageURLFunc:      handler.createPageURLFunction(category),
			SelectedCategory: category,
			AddToCartURL:     mustGenerateURL(handler.URLGenerator, CartHandler.PostAddToCart),
		})
}

func (handler ProductHabdler) createPageURLFunction(category int) func(int) string {
	return func(page int) string {
		url, _ := handler.URLGenerator.GenerateUrl(ProductHabdler.GetProducts, category, page)
		return url
	}
}

func (handler ProductHabdler) generatePageNumbers(pageCount int) (pages []int) {
	pages = make([]int, pageCount)
	for i := 0; i < pageCount; i++ {
		pages[i] = i + 1
	}
	return
}

func mustGenerateURL(generator handling.URLGenerator, target interface{}) string {
	url, err := generator.GenerateUrl(target)
	if err != nil {
		panic(err)
	}
	return url
}
