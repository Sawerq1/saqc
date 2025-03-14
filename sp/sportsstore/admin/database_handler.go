package admin

import (
	"platform/http/actionresults"
	"platform/http/handling"
	"sportsstore/models"
)

type DatabaseHandler struct {
	models.Repository
	handling.URLGenerator
}

func (handler DatabaseHandler) GetData() actionresults.ActionResult {
	return actionresults.NewTemplateAction("admin_database.html", struct {
		InitURL, SeedURL string
	}{
		InitURL: mustGenerateURL(handler.URLGenerator, DatabaseHandler.PostDatabaseInit),
		SeedURL: mustGenerateURL(handler.URLGenerator, DatabaseHandler.PostDatabaseSeed),
	})
}

func (handler DatabaseHandler) PostDatabaseInit() actionresults.ActionResult {
	handler.Repository.Init()
	return actionresults.NewRedirectAction(mustGenerateURL(handler.URLGenerator, AdminHandler.GetSection, "Database"))
}

func (handler DatabaseHandler) PostDatabaseSeed() actionresults.ActionResult {
	handler.Repository.Seed()
	return actionresults.NewRedirectAction(mustGenerateURL(handler.URLGenerator, AdminHandler.GetSection, "Database"))
}
