package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(host, port, user, password, name string) *Database {
	db, err := sql.Open("mysql", fmt.Sprintf("%[3]s:%[4]s@tcp(%[1]s:%[2]s)/%[5]s", host, port, user, password, name))
	if err != nil {
		panic(err.Error())
	}

	return &Database{db}
}

func (d Database) Prepare(query string) (*sql.Stmt, error) {
	return d.db.Prepare(query)
}

func (d Database) Query(query string) (*sql.Rows, error) {
	return d.db.Query(query)
}

func (d Database) QueryRow(query string) *sql.Row {
	return d.db.QueryRow(query)
}

func (d Database) Exec(query string, args ...any) (sql.Result, error) {
	return d.db.Exec(query, args...)
}

func (d Database) Close() {
	_ = d.db.Close()
}
