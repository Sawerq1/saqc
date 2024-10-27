package main

import "strconv"

type Product struct {
	Name, Category string
	Price          float64
}

var ProductList = []*Product{
	{"Kayak", "Watersport", 275},
	{"Lifejacket", "Watersport", 49.95},
	{"Soccer Ball", "Soccer", 19.95},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-bling King", "Chess", 12000},
}

type ProductGroup []*Product
type ProductData map[string]ProductGroup

var Products = make(ProductData)

func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}

func init() {
	for _, p := range ProductList {
		if _, ok := Products[p.Category]; ok {
			Products[p.Category] = append(Products[p.Category], p)
		} else {
			Products[p.Category] = ProductGroup{p}
		}
	}
}
