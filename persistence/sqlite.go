package persistence

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type RDB struct {
	*sql.DB
}

func NewRDB() *RDB {
	db, err := sql.Open("sqlite3", "db/semer.db")
	if err != nil {
		panic(err)
	}
	return &RDB{db}
}
