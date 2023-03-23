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

func insertRow(db *sql.DB, p *Product) (id int64) {
	res, err := db.Exec(`
	INSERT INTO Products (Name, Category, Price)
	VALUES (?, ?, ?)`, p.Name, p.Category.ID, p.Price)
	if err == nil {
		id, err = res.LastInsertId()
		if err != nil {
			Printfln("Result error: %v", err.Error())
		}
	} else {
		Printfln("Exec error: %v", err.Error())
	}
	return
}

func main() {
	//listDrivers()
	db, err := openDatabase()
	if err == nil {
		//QueryRow
		for _, id := range []int{1, 3, 10} {
			p := queryDatabase(db, id)
			Printfln("Product: %v", p)
		}

		//Exec
		newProduct := Product{Name: "Stadium", Category: Category{ID: 2}, Price: 79500}
		newID := insertRow(db, &newProduct)
		p := queryDatabase(db, int(newID))
		Printfln("New Product: %v", p)

		db.Close()
	} else {
		panic(err)
	}
}
