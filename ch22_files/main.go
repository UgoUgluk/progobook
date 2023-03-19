package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	total := 0.0
	for _, p := range Products {
		total += p.Price
	}
	dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n", time.Now().Format("Mon 15:04:05"), total)

	//WriteFile
	err := os.WriteFile("output.txt", []byte(dataStr), 0666)
	if err == nil {
		fmt.Println("Output file created")
	} else {
		Printfln("Error: %v", err.Error())
	}

	//OpenFile
	file, err := os.OpenFile("output2.txt",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		defer file.Close()
		file.WriteString(dataStr)
	} else {
		Printfln("Error: %v", err.Error())
	}

	//json
	cheapProducts := []Product{}
	for _, p := range Products {
		if p.Price < 100 {
			cheapProducts = append(cheapProducts, p)
		}
	}
	file2, err := os.OpenFile("cheap.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err == nil {
		defer file2.Close()
		encoder := json.NewEncoder(file2)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err.Error())
	}

}
