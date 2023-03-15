package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	product := "Kayak"
	description := "A boat for one person" //"This  is  double  spaced"

	//strings
	fmt.Println("Contains:", strings.Contains(product, "yak"))
	fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
	fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
	fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
	fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
	fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))

	fmt.Println("Original:", description)
	fmt.Println("ToLower:", strings.ToLower(description))
	fmt.Println("ToUpper:", strings.ToUpper(description))
	fmt.Println("Title:", strings.Title(description))
	fmt.Println("ToTitle:", strings.ToTitle(description))

	//ToTitle
	specialChar := "\u01c9"
	fmt.Println("Original:", specialChar, []byte(specialChar))
	upperChar := strings.ToUpper(specialChar)
	fmt.Println("Upper:", upperChar, []byte(upperChar))
	titleChar := strings.ToTitle(specialChar)
	fmt.Println("Title:", titleChar, []byte(titleChar))

	//unicode - rune
	for _, char := range product {
		fmt.Println(
			"Original:", string(char),
			"IsLower:", unicode.IsLower(char),
			"ToLower:", string(unicode.ToLower(char)),
			"IsUpper:", unicode.IsUpper(char),
			"ToUpper:", string(unicode.ToUpper(char)),
			"IsTitle:", unicode.IsTitle(char),
			"ToTitle:", string(unicode.ToTitle(char)),
		)
	}

	//string validation "A boat for one person"
	fmt.Println("Count:", strings.Count(description, "o"))
	fmt.Println("Index:", strings.Index(description, "o"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	fmt.Println("IndexAny:", strings.IndexAny(description, "abcd"))
	fmt.Println("LastIndexAny:", strings.LastIndexAny(description, "abcd"))
	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}
	fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))

	//splitting "A boat for one person" or "This  is  double  spaced"
	isLetterO := func(r rune) bool {
		return r == 'O' || r == 'o'
	}
	fields := strings.Fields(description)
	for _, x := range fields {
		fmt.Println("Fields >>" + x + "<<")
	}
	fieldsFunc := strings.FieldsFunc(description, isLetterO)
	for _, x := range fieldsFunc {
		fmt.Println("FieldsFunc >>" + x + "<<")
	}
	splits := strings.Split(description, " ")
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}
	splitsAfter := strings.SplitAfter(description, " ")
	for _, x := range splitsAfter {
		fmt.Println("SplitAfter >>" + x + "<<")
	}
	splitsN := strings.SplitN(description, " ", 3)
	for _, x := range splitsN {
		fmt.Println("SplitN >>" + x + "<<")
	}

	//trimming
	username := " Alice"
	trimSpace := strings.TrimSpace(username)
	fmt.Println("Trimmed:", ">>"+trimSpace+"<<")
	trimmed := strings.Trim(description, "Asno ")
	fmt.Println("Trimmed:", trimmed)
	prefixTrimmed := strings.TrimPrefix(description, "A boat")
	wrongPrefix := strings.TrimPrefix(description, "A hat ")
	fmt.Println("Trimmed:", prefixTrimmed)
	fmt.Println("Not trimmed:", wrongPrefix)
	trimmer := func(r rune) bool {
		return r == 'A' || r == 'n'
	}
	trimFunc := strings.TrimFunc(description, trimmer)
	fmt.Println("Trimmed:", trimFunc)

	//change
	text := "It was a boat. A small boat."
	replace := strings.Replace(text, "boat", "canoe", 1)
	replaceAll := strings.ReplaceAll(text, "boat", "truck")
	fmt.Println("Replace:", replace)
	fmt.Println("Replace All:", replaceAll)
	mapper := func(r rune) rune {
		if r == 'b' {
			return 'c'
		}
		return r
	}
	mapped := strings.Map(mapper, text)
	fmt.Println("Mapped:", mapped)
	//replacer
	replacer := strings.NewReplacer(
		"boat", "kayak",
		"small", "huge",
	)
	replaced := replacer.Replace(text)
	fmt.Println("Replaced:", replaced)
	//Join and Repeat
	elements := strings.Fields(text)
	joined := strings.Join(elements, "--")
	fmt.Println("Joined:", joined)
	repeat := strings.Repeat(text, 3)
	fmt.Println("Repeat:", repeat)
	//Builder
	var builder strings.Builder
	for _, sub := range strings.Fields(text) {
		if sub == "small" {
			builder.WriteString("very ")
		}
		builder.WriteString(sub)
		builder.WriteRune(' ')
	}
	fmt.Println("Builder:", builder.String())

}
