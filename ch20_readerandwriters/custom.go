package main

import "io"

//CustomReader reader with count
type CustomReader struct {
	reader    io.Reader
	readCount int
}

//NewCustomReader create custom reader
func NewCustomReader(reader io.Reader) *CustomReader {
	return &CustomReader{reader, 0}
}

func (cr *CustomReader) Read(slice []byte) (count int, err error) {
	count, err = cr.reader.Read(slice)
	cr.readCount++
	Printfln("Custom Reader: %v bytes", count)
	if err == io.EOF {
		Printfln("Total Reads: %v", cr.readCount)
	}
	return
}

//CustomWriter writer with count
type CustomWriter struct {
	writer     io.Writer
	writeCount int
}

//NewCustomWriter  create custom writer
func NewCustomWriter(writer io.Writer) *CustomWriter {
	return &CustomWriter{writer, 0}
}
func (cw *CustomWriter) Write(slice []byte) (count int, err error) {
	count, err = cw.writer.Write(slice)
	cw.writeCount++
	Printfln("Custom Writer: %v bytes", count)
	return
}

//Close for custom writer
func (cw *CustomWriter) Close() (err error) {
	if closer, ok := cw.writer.(io.Closer); ok {
		closer.Close()
	}
	Printfln("Total Writes: %v", cw.writeCount)
	return
}
