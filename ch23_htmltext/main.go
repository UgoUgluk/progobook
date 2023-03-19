package main

import (
	"html/template"
	"os"
)

// Exec Exec Template
func Exec(t *template.Template) error {
	//return t.Execute(os.Stdout, &Kayak)
	return t.Execute(os.Stdout, Products)
}

func main() {
	//Lookup
	allTemplates, err := template.ParseGlob("templates/*.html")
	if err == nil {
		//selectedTemplated := allTemplates.Lookup("template.html")
		selectedTemplated := allTemplates.Lookup("mainTemplate")
		err = Exec(selectedTemplated)
	}
}
