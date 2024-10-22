package store

const defaultTaxRate = 0.2
const minThersold = 10

type taxRate struct {
	rate, thersold float64
}

func newTaxRate(rate, thersold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if thersold < minThersold {
		thersold = minThersold
	}
	return &taxRate{rate, thersold}
}

func (taxRate *taxRate) calcTax(product *Product) float64 {
	if product.price > taxRate.thersold {
		return product.price + product.price*taxRate.rate
	}
	return product.price
}
