package mysql

import (
	"database/sql"
	//"errors"
)

type ListContentModel struct {
	DB *sql.DB
}

func (m *ListContentModel) Insert(typeID, bookID int) (int, error) {
	stmt := `INSERT INTO list_content (type_id, book_id) VALUES (?, ?)`

	result, err := m.DB.Exec(stmt, typeID, bookID)
	if err != nil {
		return 0, nil
	}

	// Use the LastInsertId() method on the result object to get the ID of our
	// newly inserted record in the list_content table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int type
	// before returning.
	return int(id), nil
}
