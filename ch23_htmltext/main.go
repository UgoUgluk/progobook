package main

import (
	"html/template"
	"os"
)

// Exec Exec Template
func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, &Kayak)
}

func main() {
	//ParseFiles
	allTemplates, err1 := template.ParseFiles("templates/template.html", "templates/extras.html")
	if err1 == nil {
		allTemplates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
		os.Stdout.WriteString("\n")
		allTemplates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
	} else {
		Printfln("Error: %v %v", err1.Error())
	}

	//ParseGlob
	os.Stdout.WriteString("\n")
	allTemplates, err := template.ParseGlob("templates/*.html")
	if err == nil {
		for _, t := range allTemplates.Templates() {
			Printfln("Template name: %v", t.Name())
		}
	} else {
		Printfln("Error: %v %v", err.Error())
	}

	//Lookup
	allTemplates, err2 := template.ParseGlob("templates/*.html")
	if err2 == nil {
		selectedTemplated := allTemplates.Lookup("template.html")
		err = Exec(selectedTemplated)
	}
	if err != nil {
		Printfln("Error: %v %v", err.Error())
	}
}
