package store

import (
	"encoding/json"
	"platform/http/actionresults"
	"platform/http/handling"
	"platform/sessions"
	"platform/validation"
	"sportsstore/models"
	"sportsstore/store/cart"
	"strings"
)

type OrderHandler struct {
	cart.Cart
	sessions.Session
	Repository   models.Repository
	URLGenerator handling.URLGenerator
	validation.Validator
}

type OrderTemlateContext struct {
	models.ShippingDetails
	ValidationErrors [][]string
	CancelURL        string
}

func (handler OrderHandler) GetCheckout() actionresults.ActionResult {
	context := OrderTemlateContext{}
	jsonData := handler.Session.GetValueDefault("checkout_details", "")
	if jsonData != nil {
		json.NewDecoder(strings.NewReader(jsonData.(string))).Decode(&context)
	}
	context.CancelURL = mustGenerateURL(handler.URLGenerator, CartHandler.GetCart)
	return actionresults.NewTemplateAction("checkout.html", context)
}

func (handler OrderHandler) PostCheckout(details models.ShippingDetails) actionresults.ActionResult {
	valid, errors := handler.Validator.Validate(details)
	if !valid {
		ctx := OrderTemlateContext{
			ShippingDetails:  details,
			ValidationErrors: [][]string{},
		}
		for _, err := range errors {
			ctx.ValidationErrors = append(ctx.ValidationErrors, []string{err.FieldName, err.Error.Error()})
		}
		builder := strings.Builder{}
		json.NewEncoder(&builder).Encode(ctx)
		handler.Session.SetValue("checkout_details", builder.String())
		redirectURL := mustGenerateURL(handler.URLGenerator, OrderHandler.GetCheckout)
		return actionresults.NewRedirectAction(redirectURL)
	} else {
		handler.Session.SetValue("checkout_details", "")
	}
	order := models.Order{
		ShippingDetails: details,
		Products:        []models.ProductSelection{},
	}
	for _, cl := range handler.Cart.GetLines() {
		order.Products = append(order.Products, models.ProductSelection{
			Quantity: cl.Quantity,
			Product:  cl.Product,
		})
	}
	handler.Repository.SaveOrder(&order)
	handler.Cart.Reset()
	targetURL, _ := handler.URLGenerator.GenerateUrl(OrderHandler.GetSummary, order.ID)
	return actionresults.NewRedirectAction(targetURL)
}

func (handler OrderHandler) GetSummary(id int) actionresults.ActionResult {
	targetURL, _ := handler.URLGenerator.GenerateUrl(ProductHabdler.GetProducts, 0, 1)
	return actionresults.NewTemplateAction("checkout_summary.html", struct {
		ID        int
		TargetURL string
	}{ID: id, TargetURL: targetURL})
}
