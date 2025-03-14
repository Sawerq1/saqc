package admin

import (
	"platform/http/actionresults"
	"platform/http/handling"
	"platform/sessions"
	"sportsstore/models"
)

type CategoriesHandler struct {
	models.Repository
	handling.URLGenerator
	sessions.Session
}

type CategoryTemplateContext struct {
	Categories []models.Category
	EditID     int
	EditURL    string
	SaveURL    string
}

const CATEGORY_EDIT_KEY string = "category_edit"

func (handler CategoriesHandler) GetData() actionresults.ActionResult {
	return actionresults.NewTemplateAction("admin_categories.html", CategoryTemplateContext{
		Categories: handler.Repository.GetCategories(),
		EditID:     handler.Session.GetValueDefault(CATEGORY_EDIT_KEY, 0).(int),
		EditURL:    mustGenerateURL(handler.URLGenerator, CategoriesHandler.PostCategoryEdit),
		SaveURL:    mustGenerateURL(handler.URLGenerator, CategoriesHandler.PostCategorySave),
	})
}

func (handler CategoriesHandler) PostCategoryEdit(ref EditReference) actionresults.ActionResult {
	handler.Session.SetValue(CATEGORY_EDIT_KEY, ref.ID)
	return actionresults.NewRedirectAction(mustGenerateURL(handler.URLGenerator,
		AdminHandler.GetSection, "Categories"))
}

func (handler CategoriesHandler) PostCategorySave(c models.Category) actionresults.ActionResult {
	handler.Repository.SaveCategory(&c)
	handler.Session.SetValue(CATEGORY_EDIT_KEY, 0)
	return actionresults.NewRedirectAction(mustGenerateURL(handler.URLGenerator, AdminHandler.GetSection, "Categories"))
}

func (handler CategoriesHandler) GetSelect(current int) actionresults.ActionResult {
	return actionresults.NewTemplateAction("select_category.html", struct {
		Current    int
		Categories []models.Category
	}{Current: current, Categories: handler.GetCategories()})
}
