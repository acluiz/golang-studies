package db

import "database/sql"

func ConnectDatabase() *sql.DB {
	conn := "
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}

	return db
}