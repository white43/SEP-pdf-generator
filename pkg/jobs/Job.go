package jobs

import "database/sql"

type Job struct {
	ID      string
	Payload string
	Result  sql.NullString
	Status  string
	Type    string
	UserId  int
}
