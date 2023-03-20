package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	go http.ListenAndServe(":5000", nil)
	time.Sleep(time.Second)

	/*formData := map[string][]string{
		"name":     {"Kayak "},
		"category": {"Watersports"},
		"price":    {"279"},
	}

	response, err := http.PostForm("http://localhost:5000/echo", formData)
	if err == nil && response.StatusCode == http.StatusOK {
		io.Copy(os.Stdout, response.Body)
		defer response.Body.Close()
	} else {
		Printfln("Error: %v", err.Error())
	}*/

	var builder strings.Builder
	err := json.NewEncoder(&builder).Encode(Products[0])
	if err == nil {
		response, err := http.Post(
			"http://localhost:5000/echo",
			"application/json",
			strings.NewReader(builder.String()),
		)
		if err == nil && response.StatusCode == http.StatusOK {
			io.Copy(os.Stdout, response.Body)
			defer response.Body.Close()
		} else {
			Printfln("Error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}

}
