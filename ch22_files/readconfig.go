package main

import (
	"encoding/json"
	"os"
	"strings"
)

// ConfigData presente config data
type ConfigData struct {
	UserName           string
	AdditionalProducts []Product
}

// Config config data
var Config ConfigData

// LoadConfig load config by ReadFile
func LoadConfig() (err error) {
	data, err := os.ReadFile("config.json")
	if err == nil {
		decoder := json.NewDecoder(strings.NewReader(string(data)))
		err = decoder.Decode(&Config)
	}
	return
}

// LoadConfigOpen load config by Open
func LoadConfigOpen() (err error) {
	file, err := os.Open("config.json")
	if err == nil {
		defer file.Close()
		nameSlice := make([]byte, 5)
		file.ReadAt(nameSlice, 20)
		Config.UserName = string(nameSlice)
		file.Seek(55, 0)
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Config.AdditionalProducts)

	}
	return
}

func init() {
	err := LoadConfig()
	if err != nil {
		Printfln("Error Loading Config: %v", err.Error())
	} else {
		Printfln("Username: %v", Config.UserName)
		Products = append(Products, Config.AdditionalProducts...)
	}

	err2 := LoadConfigOpen()
	if err2 != nil {
		Printfln("Error Loading Config: %v", err2.Error())
	} else {
		Printfln("Username: %v", Config.UserName)
		Products = append(Products, Config.AdditionalProducts...)
	}
}
