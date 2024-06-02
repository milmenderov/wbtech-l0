package postgres

import "database/sql"

type Storage struct {
	Db *sql.DB
}
