package user

import "database/sql"

type dbInterface interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}
