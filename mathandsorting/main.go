package main

// import (
// 	// "sort"
// 	// "time"
// 	// "math"
// 	// "math/rand"
// )

// func main() {
// 	// val1 := 279.00
// 	// val2 := 48.95
// 	// Printfln("Abs:%v", math.Abs(val1))
// 	// Printfln("Ceil:%v", math.Ceil(val2))
// 	// Printfln("Copysign:%v", math.Copysign(val1, -1))
// 	// Printfln("Floor:%v", math.Floor(val2))
// 	// Printfln("Max:%v", math.Max(val1, val2))
// 	// Printfln("Min:%v", math.Min(val1, val2))
// 	// Printfln("Mod:%.2f", math.Mod(val1, val2))
// 	// Printfln("Pov:%.2f", math.Pow(val2, 2))
// 	// Printfln("Round:%v", math.Round(val2))
// 	// Printfln("Round to even:%v", math.RoundToEven(val2))
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 5; i++ {
// 		Printfln("Value %v:%v", i, rand.Int())
// 	}
// }

// func intRange(max, min int) int {
// 	return (rand.Intn(max-min) + min)
// }

//	func main() {
//		rand.Seed(time.Now().UnixNano())
//		for i := 0; i < 5; i++ {
//			Printfln("Values %d:%v", i, intRange(20, 10))
//		}
//	}

// var names = []string{"Alice", "Bob", "Maxson", "Deimon"}

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	rand.Shuffle(len(names), func(first, second int) {
// 		names[first], names[second] = names[second], names[first]
// 	})
// 	for i, name := range names {
// 		Printfln("Value %d:%v", i, name)
// 	}
// }

// func main() {
// 	ints := []int{9, 2, 4, -1, 10}
// 	Printfln("Ints:%v", ints)
// 	sortInts := make([]int, len(ints))
// 	copy(sortInts, ints)
// 	sort.Ints(sortInts)
// 	Printfln("Ints sorted:%v", sortInts)
// 	indexOf4 := sort.SearchInts(sortInts, 4)
// 	indexOf3 := sort.SearchInts(sortInts, 3)
// 	Printfln("Index of 4: %v (present: %v)", indexOf4, sortInts[indexOf4] == 4)
// 	Printfln("Index of 3: %v (present: %v)", indexOf3, sortInts[indexOf3] == 3)
// floats := []float64{0.52, -45.67, 23.46, 52.25}
// Printfln("Float64:%v", floats)
// sort.Float64s(floats)
// Printfln("Sorted float64:%v", floats)

// strings := []string{"Kayak", "Lifejacket", "Soccer ball"}
// Printfln("Strings:%v", strings)
// if !sort.StringsAreSorted(strings) {
// 	sort.Strings(strings)
// 	Printfln("Sorted strings:%v", strings)
// } else {
// 	Printfln("Strings alredy sorted")
// }
// }

func main() {
	products := ProductSlice{
		{"Kayak", 275},
		{"Lifejacket", 48.98},
		{"Soccer ball", 19.50},
	}
	//ProductSlices(products)
	// ProductSlicesByName(products)
	// for _, p := range products {
	// 	Printfln("Name: %v Price: %.2f", p.Name, p.Price)
	// }
	SortWith(products, func(p1, p2 Product) bool {
		return p1.Name < p2.Name
	})
	for _, p := range products {
		Printfln("Name: %v Price: %.2f", p.Name, p.Price)
	}
}
