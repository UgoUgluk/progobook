package main

import (
	"bufio"
	//"io"
	"strings"
)

func main() {

	text := "It was a boat. A small boat."
	/*var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder
	slice := make([]byte, 5)

	buffered := bufio.NewReader(reader)

	for {
		count, err := buffered.Read(slice)
		if count > 0 {
			Printfln("Buffer size: %v, buffered: %v", buffered.Size(), buffered.Buffered())
			writer.Write(slice[0:count])
		}
		if err != nil {
			break
		}
	}
	Printfln("Read data: %v", writer.String())*/
	var builder strings.Builder
	var writer = bufio.NewWriterSize(NewCustomWriter(&builder), 20)
	for i := 0; true; {
		end := i + 5
		if end >= len(text) {
			writer.Write([]byte(text[i:]))
			writer.Flush()
			break
		}
		writer.Write([]byte(text[i:end]))
		i = end
	}
	Printfln("Written data: %v", builder.String())

}
