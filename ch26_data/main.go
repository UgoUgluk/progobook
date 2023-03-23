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

func insertAndUseCategory(db *sql.DB, name string, productIDs ...int) (err error) {
	tx, err := db.Begin()
	updatedFailed := false
	if err == nil {
		catResult, err := tx.Stmt(insertNewCategory).Exec(name)
		if err == nil {
			newID, _ := catResult.LastInsertId()
			preparedStatement := tx.Stmt(changeProductCategory)
			for _, id := range productIDs {
				changeResult, err := preparedStatement.Exec(newID, id)
				if err == nil {
					changes, _ := changeResult.RowsAffected()
					if changes == 0 {
						updatedFailed = true
						break
					}
				}
			}

		}
	}
	if err != nil || updatedFailed {
		Printfln("Aborting transaction %v", err)
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}

func main() {
	//listDrivers()
	db, err := openDatabase()
	if err == nil {
		insertAndUseCategory(db, "Category_1", 2)
		p := queryDatabase(db, 2)
		Printfln("Product: %v", p)
		insertAndUseCategory(db, "Category_2", 100)

		db.Close()
	} else {
		panic(err)
	}
}
