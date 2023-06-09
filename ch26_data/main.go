package main

import (
	"database/sql"
)

// Category present a category
type Category struct {
	ID   int
	Name string
}

// Product present a product
type Product struct {
	ID   int
	Name string
	Category
	Price float64
}

func queryDatabase(db *sql.DB) (products []Product, err error) {
	rows, err := db.Query(`SELECT Products.Id, Products.Name,
		Products.Price,
		Categories.Id as "Category.Id", Categories.Name as
		"Category.Name"
	FROM Products, Categories
	WHERE Products.Category = Categories.Id`)
	if err == nil {
		results, err := scanIntoStruct(rows, &Product{})
		if err == nil {
			products = (results).([]Product)
		} else {
			Printfln("Scanning error: %v", err)
		}
	}
	return
}

func main() {
	//listDrivers()
	db, err := openDatabase()
	if err == nil {
		products, _ := queryDatabase(db)
		for _, p := range products {
			Printfln("Product: %v", p)
		}

		db.Close()
	} else {
		panic(err)
	}
}
