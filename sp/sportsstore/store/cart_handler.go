package store

import (
	"platform/http/actionresults"
	"platform/http/handling"
	"sportsstore/models"
	"sportsstore/store/cart"
)

type CartHandler struct {
	models.Repository
	cart.Cart
	handling.URLGenerator
}

type CartTemplateContext struct {
	cart.Cart
	ProductListURL string
	CartURL        string
	CheckoutURL    string
	RemoveURL      string
}

func (handler CartHandler) GetCart() actionresults.ActionResult {
	return actionresults.NewTemplateAction("cart.html", CartTemplateContext{
		Cart:           handler.Cart,
		ProductListURL: handler.mustGenerateURL(ProductHabdler.GetProducts, 0, 1),
		RemoveURL:      handler.mustGenerateURL(CartHandler.PostRemoveFromCart),
		CheckoutURL:    handler.mustGenerateURL(OrderHandler.GetCheckout),
	})
}

func (handler CartHandler) GetWidget() actionresults.ActionResult {
	return actionresults.NewTemplateAction("cart_widget.html", CartTemplateContext{
		Cart:    handler.Cart,
		CartURL: handler.mustGenerateURL(CartHandler.GetCart),
	})
}

type CartProductReference struct {
	ID int
}

func (handler CartHandler) PostAddToCart(ref CartProductReference) actionresults.ActionResult {
	p := handler.Repository.GetProduct(ref.ID)
	handler.Cart.AddProduct(p)
	return actionresults.NewRedirectAction(handler.mustGenerateURL(CartHandler.GetCart))
}

func (handler CartHandler) PostRemoveFromCart(ref CartProductReference) actionresults.ActionResult {
	handler.RemoveLineForProduct(ref.ID)
	return actionresults.NewRedirectAction(handler.mustGenerateURL(CartHandler.GetCart))
}

func (handler CartHandler) mustGenerateURL(method interface{}, data ...interface{}) string {
	url, err := handler.URLGenerator.GenerateUrl(method, data...)
	if err != nil {
		panic(err)
	}
	return url
}
