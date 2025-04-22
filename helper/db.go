package helper

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewDb() (*sql.DB, func(), error) {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306)/" + dbName

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, nil, err
	}
	db.SetMaxOpenConns(10)

	cleanup := func() {
		db.Close()
	}

	return db, cleanup, nil
}
