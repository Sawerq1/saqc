package cart

import (
	"encoding/json"
	"platform/services"
	"platform/sessions"
	"sportsstore/models"
	"strings"
)

const CART_KEY string = "cart"

func RegisterCartService() {
	services.AddScoped(func(session sessions.Session) Cart {
		lines := []*CartLine{}
		sessionVal := session.GetValue(CART_KEY)
		if strVal, ok := sessionVal.(string); ok {
			json.NewDecoder(strings.NewReader(strVal)).Decode(&lines)
		}
		return &SessionCart{
			BasicCart: &BasicCart{lines: lines},
			Session:   session,
		}
	})
}

type SessionCart struct {
	*BasicCart
	sessions.Session
}

func (sc *SessionCart) AddProduct(p models.Product) {
	sc.BasicCart.AddProduct(p)
	sc.SaveToSession()
}

func (sc *SessionCart) RemoveLineForProduct(id int) {
	sc.BasicCart.RemoveLineForProduct(id)
	sc.SaveToSession()
}

func (sc *SessionCart) SaveToSession() {
	builder := strings.Builder{}
	json.NewEncoder(&builder).Encode(sc.lines)
	sc.Session.SetValue(CART_KEY, builder.String())
}

func (sc *SessionCart) Reset() {
	sc.lines = []*CartLine{}
	sc.SaveToSession()
}
