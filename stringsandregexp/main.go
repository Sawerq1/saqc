package main

import (
	"fmt"
	//"strings"
	//"unicode"
	"regexp"
)

// func getSubstring(s string, indexes []int) string {
// 	return (s[indexes[0]:indexes[1]])
// }

func main() {
	//product := "Kayak"
	// fmt.Println("Contains:", strings.Contains(product, "yak"))
	// fmt.Println("Contains any:", strings.ContainsAny(product, "abc"))
	// fmt.Println("Contains rune:", strings.ContainsRune(product, 'K'))
	// fmt.Println("EqualFold:", strings.EqualFold(product, "kAYAK"))
	// fmt.Println("HasSufix:", strings.HasSuffix(product, "ak"))
	// fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
	// description := "A boat for selling"
	// fmt.Println("To lower:", strings.ToLower(description))
	// fmt.Println("To upper:", strings.ToUpper(description))
	// fmt.Println("Original:", description)
	// fmt.Println("To title:", strings.Title(description))
	// specialChar := "\u01c9"
	// fmt.Println("Original:", specialChar, []byte(specialChar))
	// upperChar := strings.ToUpper(specialChar)
	// fmt.Println("To upper:", upperChar, []byte(upperChar))
	// titleChar := strings.ToTitle(specialChar)
	// fmt.Println("To title:", titleChar, []byte(titleChar))
	// for _, char := range product {
	// 	fmt.Println(string(char), "To upper:", unicode.IsUpper(char))
	// }
	// description := "A boat for one person"
	// fmt.Println("Count:", strings.Count(description, "o"))
	// fmt.Println("Index:", strings.Index(description, "o"))
	// fmt.Println("Last index:", strings.LastIndex(description, "o"))
	// fmt.Println("idex any:", strings.IndexAny(description, "abcd"))
	// fmt.Println("Last index:", strings.LastIndex(description, "o"))
	// fmt.Println("Last index any:", strings.LastIndexAny(description, "abcd"))
	// isLetterB := func(r rune) bool {
	// 	return r == 'B' || r == 'b'
	// }
	// isLetterO := func(r rune) bool {
	// 	return r == 'O' || r == 'o'
	// }
	// fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterO))
	// fmt.Println("Last index:", strings.LastIndexFunc(description, isLetterO))
	// splits := strings.Split(description, " ")
	// for _, msg := range splits {
	// 	fmt.Println("Split >>" + msg + "<<")
	// }
	// splitsAlter := strings.SplitAfter(description, " ")
	// for _, msg := range splitsAlter {
	// 	fmt.Println("Split alter >>" + msg + "<<")
	// }
	// splitNa := strings.SplitN(description, " ", 2)
	// for _, msg := range splitNa {
	// 	fmt.Println("Split N >>" + msg + "<<")
	// }
	// splitsAfterNa := strings.SplitAfterN(description, " ", 2)
	// for _, msg := range splitsAfterNa {
	// 	fmt.Println("Split after N >>" + msg + "<<")
	// }
	//description := "This  is  double spaced"
	// spliter := func(r rune) bool {
	// 	return r == ' '
	// }
	// splits := strings.FieldsFunc(description, spliter)
	// for _, msg := range splits {
	// 	fmt.Println(msg)
	// }
	// name := " Alice"
	// fmt.Println(">>" + strings.TrimSpace(name) + "<<")
	// fmt.Println(">>" + strings.Trim(description, "Asno ") + "<<")
	// fmt.Println(">>" + strings.TrimPrefix(description, "A boat ") + "<<")
	// fmt.Println(">>" + strings.TrimSuffix(description, " person") + "<<")
	// trimmer := func(r rune) bool {
	// 	return r == 'A' || r == 'n'
	// }
	// fmt.Println(">>" + strings.TrimFunc(description, trimmer) + "<<")
	// fmt.Println(">>" + strings.TrimRightFunc(description, trimmer) + "<<")
	// fmt.Println(">>" + strings.TrimLeftFunc(description, trimmer) + "<<")
	// text := "It was a boat. A small boat."
	// replace := strings.Replace(text, "boat", "kayak", 1)
	// replaceAll := strings.ReplaceAll(text, "boat", "truck")
	// fmt.Println(replace)
	// fmt.Println(replaceAll)
	// mapper := func(r rune) rune {
	// 	if r == 'b' {
	// 		return 'c'
	// 	}
	// 	return r
	// }
	// mapped := strings.Map(mapper, text)
	// fmt.Println(mapped)
	// replacer := strings.NewReplacer("boat", "kayak", "small", "huge")
	// replaced := replacer.Replace(text)
	// fmt.Println(text)
	// fmt.Println(replaced)
	// elments := strings.Fields(text)
	// fmt.Println(elments)
	// joined := strings.Join(elments, "--")
	// fmt.Println(joined)
	// var builder strings.Builder
	// for _, sub := range strings.Fields(text) {
	// 	if sub == "small" {
	// 		builder.WriteString("very ")
	// 	}
	// 	builder.WriteString(sub)
	// 	builder.WriteRune(' ')
	// }
	// fmt.Println(builder.String())
	// pattern, compileErr := regexp.Compile("[A-z]oat")
	// description := "A boat for one person"
	// question := "Is that a goat?"
	// perfomanse := "I like oats"
	// match, err := regexp.MatchString("[A-z]oat", description)
	// if err == nil {
	// 	fmt.Println(match)
	// } else {
	// 	fmt.Println(err)
	// }
	// if compileErr == nil {
	// 	fmt.Println("Desckription:", pattern.MatchString(description))
	// 	fmt.Println("Question:", pattern.MatchString(question))
	// 	fmt.Println("Pefomanse:", pattern.MatchString(perfomanse))
	// } else {
	// 	fmt.Println("Error:", compileErr)
	// }
	pattern := regexp.MustCompile("K[A-z]{4}|[A-z]oat")
	description := "Kayak. A boat for one person."
	// firstIndex := pattern.FindStringIndex(description)
	// allIndex := pattern.FindAllStringIndex(description, -1)
	// // fmt.Println("First insex:", firstIndex[0], "-", firstIndex[1], "=", getSubstring(description, firstIndex))
	// for i, idx := range allIndex {
	// 	fmt.Println("Index", i, "=", idx[0], "-", idx[1], "=", getSubstring(description, idx))
	// }
	firstMatch := pattern.FindString(description)
	allMatch := pattern.FindAllString(description, -1)
	fmt.Println("First match:", firstMatch)
	for i, m := range allMatch {
		fmt.Println("Match", i, "--", m)
	}
}
