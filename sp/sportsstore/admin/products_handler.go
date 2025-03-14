package admin

import (
	"platform/http/actionresults"
	"platform/http/handling"
	"platform/sessions"
	"sportsstore/models"
)

type ProductsHandler struct {
	models.Repository
	handling.URLGenerator
	sessions.Session
}

type ProductTemplateContext struct {
	Products []models.Product
	EditID   int
	EditURL  string
	SaveURL  string
}

const PRODUCT_EDIT_KEY string = "product_edit"

func (handler ProductsHandler) GetData() actionresults.ActionResult {
	return actionresults.NewTemplateAction("admin_products.html", ProductTemplateContext{
		Products: handler.GetProducts(),
		EditID:   handler.Session.GetValueDefault(PRODUCT_EDIT_KEY, 0).(int),
		EditURL:  mustGenerateURL(handler.URLGenerator, ProductsHandler.PostProductEdit),
		SaveURL:  mustGenerateURL(handler.URLGenerator, ProductsHandler.PostProductSave),
	})
}

type EditReference struct {
	ID int
}

func (handler ProductsHandler) PostProductEdit(ref EditReference) actionresults.ActionResult {
	handler.Session.SetValue(PRODUCT_EDIT_KEY, ref.ID)
	return actionresults.NewRedirectAction(mustGenerateURL(handler.URLGenerator, AdminHandler.GetSection, "Products"))
}

type ProductSaveReference struct {
	ID                int
	Name, Description string
	Category          int
	Price             float64
}

func (handler ProductsHandler) PostProductSave(p ProductSaveReference) actionresults.ActionResult {
	handler.Repository.SaveProduct(&models.Product{
		ID: p.ID, Name: p.Name, Description: p.Description,
		Category: &models.Category{ID: p.Category},
		Price:    p.Price,
	})
	handler.Session.SetValue(PRODUCT_EDIT_KEY, 0)
	return actionresults.NewRedirectAction(
		mustGenerateURL(handler.URLGenerator, AdminHandler.GetSection, "Products"))
}

func mustGenerateURL(gen handling.URLGenerator, target interface{}, data ...interface{}) string {
	URL, err := gen.GenerateUrl(target, data...)
	if err != nil {
		panic(err)
	}
	return URL
}
