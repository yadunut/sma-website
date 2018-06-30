package postgresql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sqlx.DB
}

// Open opens a new connection to the database
func Open(connstr string) (*Database, error) {
	db, err := sqlx.Connect("postgres", connstr)
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

// Close closes connection to the database
func (db *Database) Close() {
	db.db.Close()
}
