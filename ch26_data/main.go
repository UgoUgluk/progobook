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

func queryDatabase(db *sql.DB, id int) (p Product) {
	row := db.QueryRow(`
	SELECT Products.Id, Products.Name, Products.Price,
	Categories.Id as Cat_Id, Categories.Name as CatName
	FROM Products, Categories
	WHERE Products.Category = Categories.Id
	AND Products.Id = ?`, id)
	if row.Err() == nil {
		scanErr := row.Scan(&p.ID, &p.Name, &p.Price,
			&p.Category.ID, &p.Category.Name)
		if scanErr != nil {
			Printfln("Scan error: %v", scanErr)
		}
	} else {
		Printfln("Row error: %v", row.Err().Error())
	}
	return
}

func insertAndUseCategory(name string, productIDs ...int) {
	result, err := insertNewCategory.Exec(name)
	if err == nil {
		newID, _ := result.LastInsertId()
		for _, id := range productIDs {
			changeProductCategory.Exec(int(newID), id)
		}
	} else {
		Printfln("Prepared statement error: %v", err)
	}
}

func main() {
	//listDrivers()
	db, err := openDatabase()
	if err == nil {
		insertAndUseCategory("Misc Products", 2)
		p := queryDatabase(db, 2)
		Printfln("Product: %v", p)

		db.Close()
	} else {
		panic(err)
	}
}
