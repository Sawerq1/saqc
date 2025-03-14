package admin

import (
	"platform/http/actionresults"
	"platform/http/handling"
)

var sectionNames = []string{"Products", "Categories", "Orders", "Database"}

type AdminHandler struct {
	handling.URLGenerator
}

type AdminTemplateContext struct {
	Sections       []string
	ActiveSection  string
	SectionURLFunc func(string) string
}

func (handler AdminHandler) GetSection(section string) actionresults.ActionResult {
	return actionresults.NewTemplateAction("admin.html", AdminTemplateContext{
		Sections:      sectionNames,
		ActiveSection: section,
		SectionURLFunc: func(sec string) string {
			sectionURL, _ := handler.GenerateUrl(AdminHandler.GetSection, sec)
			return sectionURL
		},
	})
}
