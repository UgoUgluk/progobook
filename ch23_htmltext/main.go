package main

import (
	"html/template"
	"os"
	"strings"
)

// GetCategories get products categories
func GetCategories(products []Product) (categories []string) {
	catMap := map[string]string{}
	for _, p := range products {
		if catMap[p.Category] == "" {
			catMap[p.Category] = p.Category
			categories = append(categories, p.Category)
		}
	}
	return
}

// Exec Exec Template
func Exec(t *template.Template) error {
	//return t.Execute(os.Stdout, &Kayak)
	return t.Execute(os.Stdout, Products)
}

func main() {
	allTemplates := template.New("allTemplates")
	allTemplates.Funcs(map[string]interface{}{
		"getCats": GetCategories,
		"lower":   strings.ToLower,
	})
	allTemplates, err := allTemplates.ParseGlob("templates/*.html")

	if err == nil {
		//selectedTemplated := allTemplates.Lookup("template.html")
		selectedTemplated := allTemplates.Lookup("mainTemplate")
		err = Exec(selectedTemplated)
	}
}
