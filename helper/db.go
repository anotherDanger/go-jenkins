package helper

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDb() (*sql.DB, func(), error) {
	db, err := sql.Open("mysql", "root:andhikad@/library_new")
	if err != nil {
		return nil, nil, err
	}
	db.SetMaxOpenConns(10)

	cleanup := func() {
		db.Close()
	}

	return db, cleanup, nil
}
