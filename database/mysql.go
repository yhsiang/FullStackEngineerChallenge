package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func New() (DB, error) {
	// TODO: add os.LookupEnv here
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		return DB{}, err
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	err = db.Ping()
	if err != nil {
		return DB{}, err
	}

	return DB{db}, nil
}
