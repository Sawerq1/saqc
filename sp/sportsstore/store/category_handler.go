package store

import (
	"platform/http/actionresults"
	"platform/http/handling"
	"sportsstore/models"
)

type CategoryHandler struct {
	Repository   models.Repository
	URLGenerator handling.URLGenerator
}

type categoryTemplateContext struct {
	Categories       []models.Category
	SelectedCategory int
	CategoryURLFunc  func(int) string
}

func (handler CategoryHandler) GetButtons(selected int) actionresults.ActionResult {
	return actionresults.NewTemplateAction("category_buttons.html", categoryTemplateContext{
		Categories:       handler.Repository.GetCategories(),
		SelectedCategory: selected,
		CategoryURLFunc:  handler.createCategoryFilterFunction(),
	})
}

func (handler CategoryHandler) createCategoryFilterFunction() func(int) string {
	return func(category int) string {
		url, _ := handler.URLGenerator.GenerateUrl(ProductHabdler.GetProducts, category, 1)
		return url
	}
}
