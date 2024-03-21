package controller

import "database/sql"

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/uaspbo")
	if err != nil {
		panic(err)
	}
	return db
}
