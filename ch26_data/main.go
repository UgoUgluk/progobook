package main

import (
	"database/sql"
)

func queryDatabase(db *sql.DB) {
	rows, err := db.Query("SELECT * from Products")
	if err == nil {
		for rows.Next() {
			var id, category int
			var name int
			var price float64
			scanErr := rows.Scan(&id, &name, &category, &price)
			if scanErr == nil {
				Printfln("Row: %v %v %v %v", id, name, category,
					price)
			} else {
				Printfln("Scan error: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error: %v", err)
	}
}

func main() {
	//listDrivers()
	db, err := openDatabase()
	if err == nil {
		queryDatabase(db)
		db.Close()
	} else {
		panic(err)
	}
}
