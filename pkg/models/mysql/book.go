package mysql

import (
	"database/sql"
	//"errors"
)

type BookModel struct {
	DB *sql.DB
}
