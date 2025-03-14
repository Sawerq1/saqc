package admin

import (
	"platform/authorization/identity"
	"platform/http/actionresults"
	"platform/http/handling"
)

type SignOutHandler struct {
	identity.User
	handling.URLGenerator
}

func (handler SignOutHandler) GetUserWidget() actionresults.ActionResult {
	return actionresults.NewTemplateAction("user_widget.html", struct {
		identity.User
		SignoutURL string
	}{
		handler.User,
		mustGenerateURL(handler.URLGenerator, AuthenticationHandler.PostSignOut),
	})
}
