package database

import (
	"database/sql"

	"example.com/auth/adapters/database/queries"
	// Driver
	_ "github.com/mattn/go-sqlite3"
)

func execute(callback func(_db *sql.DB)) {
	db, err := sql.Open("sqlite3", ":memory:")
	checkDBError(err)
	defer db.Close()

	initialize(db)
	callback(db)
}

func initialize(db *sql.DB) {
	_, err := db.Exec(queries.InitializeDB())
	checkDBError(err)

	stmt, err := db.Prepare(queries.InsertUser())
	checkDBError(err)

	_, err = stmt.Exec("John", "john@teste.com.br", "pass")
	checkDBError(err)
}

func checkDBError(err error) {
	if err != nil {
		panic(err)
	}
}
