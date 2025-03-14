package admin

import (
	"platform/http/actionresults"
	"platform/http/handling"
	"sportsstore/models"
)

type OrdersHandler struct {
	models.Repository
	handling.URLGenerator
}

func (handler OrdersHandler) GetData() actionresults.ActionResult {
	return actionresults.NewTemplateAction("admin_orders.html", struct {
		Orders      []models.Order
		CallbackURL string
	}{
		Orders:      handler.Repository.GetOrders(),
		CallbackURL: mustGenerateURL(handler.URLGenerator, OrdersHandler.PostOrderToggle),
	})
}

func (handler OrdersHandler) PostOrderToggle(ref EditReference) actionresults.ActionResult {
	order := handler.Repository.GetOrder(ref.ID)
	order.Shipped = !order.Shipped
	handler.Repository.SetOrderShipped(&order)
	return actionresults.NewRedirectAction(mustGenerateURL(handler.URLGenerator, AdminHandler.GetSection, "Orders"))
}
