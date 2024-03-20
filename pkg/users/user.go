package users

import "database/sql"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	Token     sql.NullString
	Balance   float64
	Status    string
}
