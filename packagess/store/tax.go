package store

const defaultTaxRate = 0.2
const minThersold = 10

type taxRate struct {
	rate, thersold float64
}

var categoryMaxPrices = map[string]float64{
	"Watersport": 250,
	"Chess":      50,
	"Soccer":     150,
}

func init() {
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price + (price * defaultTaxRate)
	}
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

func (taxRate *taxRate) calcTax(product *Product) (price float64) {
	if product.price > taxRate.thersold {
		price = product.price + (product.price * taxRate.rate)
	} else {
		price = product.price
	}
	if max, ok := categoryMaxPrices[product.Category]; ok && price > max {
		price = max
	}
	return
}
