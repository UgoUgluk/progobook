package main

import (
	"io"
	"strings"
)

/*
	func processData(reader io.Reader, writer io.Writer) {
		count, err := io.Copy(writer, reader)
		if err == nil {
			Printfln("Read %v bytes", count)
		} else {
			Printfln("Error: %v", err.Error())
		}

}
*/
func main() {
	/*r := strings.NewReader("Kayak")
	var builder strings.Builder
	//Copy form r to builder
	processData(r, &builder)
	Printfln("String builder contents: %s", builder.String())*/

	//Pipe
	pipeReader, pipeWriter := io.Pipe()
	go GenerateData(pipeWriter)
	ConsumeData(pipeReader)

	//MultiReader
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	ConsumeData(concatReader)

	//MultiWriter
	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder
	combinedWriter := io.MultiWriter(&w1, &w2, &w3)
	GenerateData(combinedWriter)
	Printfln("Writer #1: %v", w1.String())
	Printfln("Writer #2: %v", w2.String())
	Printfln("Writer #3: %v", w3.String())

	//TeeReader
	r4 := strings.NewReader("Kayak")
	r5 := strings.NewReader("Lifejacket")
	r6 := strings.NewReader("Canoe")
	concatReader2 := io.MultiReader(r4, r5, r6)
	var writer strings.Builder
	teeReader := io.TeeReader(concatReader2, &writer)
	ConsumeData(teeReader)
	Printfln("Echo data: %v", writer.String())

	//LimitReader
	r7 := strings.NewReader("Kayak")
	r8 := strings.NewReader("Lifejacket")
	r9 := strings.NewReader("Canoe")
	concatReader3 := io.MultiReader(r7, r8, r9)
	limited := io.LimitReader(concatReader3, 5)
	ConsumeData(limited)

}
