package main

import (
	//"bufio"
	"fmt"
	"io"
	"strings"
)

func scanFromReader(reader io.Reader, template string, vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}
func scanSingle(reader io.Reader, val interface{}) (int, error) {
	return fmt.Fscan(reader, val)
}

func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
	fmt.Fprintf(writer, template, vals...)
}

func writeReplaced(writer io.Writer, str string, subs ...string) {
	replacer := strings.NewReplacer(subs...)
	replacer.WriteString(writer, str)
}

func main() {

	//scanFromReader
	reader := strings.NewReader("Kayak Watersports $279.00")
	var name, category string
	var price float64
	scanTemplate := "%s %s $%f"
	_, err := scanFromReader(reader, scanTemplate, &name, &category, &price)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Name: %v", name)
		Printfln("Category: %v", category)
		Printfln("Price: %.2f", price)
	}

	//scanSingle
	reader2 := strings.NewReader("Kayak Watersports $279.00")
	for {
		var str string
		_, err := scanSingle(reader2, &str)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		Printfln("Value: %v", str)
	}

	//writeFormatted
	var writer strings.Builder
	template := "Name: %s, Category: %s, Price: $%.2f"
	writeFormatted(&writer, template, "Kayak", "Watersports", float64(279))
	fmt.Println(writer.String())

	//writeReplaced
	text := "It was a boat. A small boat."
	subs := []string{"boat", "kayak", "small", "huge"}
	var writer2 strings.Builder
	writeReplaced(&writer2, text, subs...)
	fmt.Println(writer2.String())
}
