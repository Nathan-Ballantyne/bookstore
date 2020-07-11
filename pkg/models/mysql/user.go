package mysql

import (
	"database/sql"
	//"errors"
)

type UserModel struct {
	DB *sql.DB
}
