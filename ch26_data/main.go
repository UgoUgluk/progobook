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

func queryDatabase(db *sql.DB, categoryName string) []Product {
	products := []Product{}
	rows, err := db.Query(`
		SELECT Products.Id, Products.Name, Products.Price,
			Categories.Id as Cat_Id, Categories.Name as CatName
			FROM Products, Categories
		WHERE Products.Category = Categories.Id 
			AND Categories.Name = ?`, categoryName)
	if err == nil {
		for rows.Next() {
			p := Product{}
			scanErr := rows.Scan(
				&p.ID,
				&p.Name,
				&p.Price,
				&p.Category.ID,
				&p.Category.Name,
			)
			if scanErr == nil {
				products = append(products, p)
			} else {
				Printfln("Scan error: %v", scanErr)
				break
			}

		}
	} else {
		Printfln("Error: %v", err)
	}
	return products
}

func main() {
	//listDrivers()
	db, err := openDatabase()
	if err == nil {
		for _, cat := range []string{"Soccer", "Watersports"} {
			Printfln("--- %v Results ---", cat)
			products := queryDatabase(db, cat)
			for i, p := range products {
				Printfln("#%v: %v", i, p)
			}
		}
		db.Close()
	} else {
		panic(err)
	}
}
