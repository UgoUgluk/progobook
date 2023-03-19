package main

import (
	// "fmt"
	// "time"
	"encoding/json"
	"os"
	"path/filepath"
)

func main() {
	path, err := os.UserHomeDir()
	if err == nil {
		path = filepath.Join(path, "MyApp", "MyTempFile.json")
	}
	Printfln("Full path: %v", path)
	Printfln("Volume name: %v", filepath.VolumeName(path))
	Printfln("Dir component: %v", filepath.Dir(path))
	Printfln("File component: %v", filepath.Base(path))
	Printfln("File extension: %v", filepath.Ext(path))

	err = os.MkdirAll(filepath.Dir(path), 0766)
	if err == nil {
		file, err := os.OpenFile(path, os.O_CREATE|
			os.O_WRONLY, 0666)
		if err == nil {
			defer file.Close()
			encoder := json.NewEncoder(file)
			encoder.Encode(Products)
		}
	}
	if err != nil {
		Printfln("Error %v", err.Error())
	}

}
