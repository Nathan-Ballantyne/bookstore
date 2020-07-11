package mysql

import (
	"database/sql"
	"errors"

	"github.com/Nathan-Ballantyne/bookstore/pkg/models"
)

type ListTypeModel struct {
	DB *sql.DB
}

func (m *ListTypeModel) NewList(userID int, name string) (int, error) {
	stmt := `INSERT INTO list_type (user_id, name) VALUES (?, ?)`

	result, err := m.DB.Exec(stmt, userID, name)
	if err != nil {
		return 0, nil
	}

	// Use the LastInsertId() method on the result object to get the ID of our
	// newly inserted record in the list_type table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int type
	// before returning.
	return int(id), nil
}

func (m *ListTypeModel) GetUserList(id int) (*models.ListType, error) {
	l := &models.ListType{}
	stmt := `SELECT id, user_id, name FROM list_type WHERE id = ?`

	err := m.DB.QueryRow(stmt, id).Scan(&l.ID, &l.UserID, &l.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return l, nil
}

func (m *ListTypeModel) GetName(name string) (*models.ListType, error) {
	l := &models.ListType{}
	stmt := `SELECT id, user_id, name FROM list_type WHERE name LIKE '%` + name + `%'`

	err := m.DB.QueryRow(stmt, name).Scan(&l.ID, &l.UserID, &l.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return l, nil
}
