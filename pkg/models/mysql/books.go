package mysql

import (
	"database/sql"

	"errors"

	"github.com/Nathan-Ballantyne/bookstore/pkg/models"
)

type BookModel struct {
	DB *sql.DB
}

// this will Insert a book into the DB with the following info
func (m *BookModel) Insert(title, author, cover, series, readStatus string, releaseYear, pageCount, rating int) (int, error) {
	stmt := `INSERT INTO book (title, author, release_year, page_count, cover, series, read_status, rating)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, title, author, releaseYear, pageCount, cover, series, readStatus, rating)
	if err != nil {
		return 0, nil
	}

	// Use the LastInsertId() method on the result object to get the ID of our
	// newly inserted record in the book table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int type
	// before returning.
	return int(id), nil
}

// This will return/ Get a specific book based on its id.
func (m *BookModel) Get(id int) (*models.Book, error) {
	b := &models.Book{}
	stmt := `SELECT id, 
					title, 
					author, 
					release_year, 
					page_count, 
					cover, series, 
					read_status, 
					rating 
			FROM book    
			WHERE id = ?`
	err := m.DB.QueryRow(stmt, id).Scan(&b.ID, &b.Title, &b.Author, &b.ReleaseYear, &b.PageCount, &b.Cover, &b.Series, &b.ReadStatus, &b.Rating)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return b, nil
}
