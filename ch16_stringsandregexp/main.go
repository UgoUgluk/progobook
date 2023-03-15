package main

import (
	"fmt"
	//"strings"
	"regexp"
)

func getSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}

func main() {
	pattern, compileErr := regexp.Compile("[A-z]oat")

	description := "A boat for one person"
	question := "Is that a goat?"
	preference := "I like oats"

	if compileErr == nil {
		fmt.Println("Description:", pattern.MatchString(description))
		fmt.Println("Question:", pattern.MatchString(question))
		fmt.Println("Preference:", pattern.MatchString(preference))
	} else {
		fmt.Println("Error:", compileErr)
	}

	pattern2 := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description2 := "Kayak. A boat for one person."
	firstIndex := pattern2.FindStringIndex(description2)
	allIndices := pattern2.FindAllStringIndex(description2, -1)
	fmt.Println(
		"First index", firstIndex[0],
		"-", firstIndex[1],
		"=", getSubstring(description2, firstIndex))
	for i, idx := range allIndices {
		fmt.Println("Index", i, "=", idx[0], "-",
			idx[1], "=", getSubstring(description2, idx))
	}

	firstMatch := pattern2.FindString(description2)
	allMatches := pattern2.FindAllString(description2, -1)
	fmt.Println("First match:", firstMatch)
	for i, m := range allMatches {
		fmt.Println("Match", i, "=", m)
	}
	//subs
	pattern3 := regexp.MustCompile("A ([A-z]*) for ([A-z]*) person")
	subs := pattern3.FindStringSubmatch(description2)
	for _, s := range subs {
		fmt.Println("Subs Match:", s)
	}
	//named subs
	pattern4 := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	subs2 := pattern4.FindStringSubmatch(description2)
	for _, name := range []string{"type", "capacity"} {
		fmt.Println(
			name, "=",
			subs2[pattern4.SubexpIndex(name)],
		)
	}
	//replace subs
	template := "(type: ${type}, capacity: ${capacity})"
	replaced := pattern4.ReplaceAllString(description2, template)
	fmt.Println(replaced)
	//replace by func
	replaced2 := pattern4.ReplaceAllStringFunc(
		description2,
		func(s string) string {
			return "This is the replacement content"
		},
	)
	fmt.Println(replaced2)

}
