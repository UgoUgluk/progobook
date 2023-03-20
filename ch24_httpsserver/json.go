package main

import (
	"encoding/json"
	"net/http"
)

// HandleJSONRequest handle for request
func HandleJSONRequest(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(Products)
}
func init() {
	http.HandleFunc("/json", HandleJSONRequest)
}
