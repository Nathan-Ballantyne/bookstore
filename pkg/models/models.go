package models

import (
	"errors"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
	// Add a new ErrInvalidCredentials error. We'll use this later if a user
	// tries to login with an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// Add a new ErrDuplicateEmail error. We'll use this later if a user
	// tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword []byte `json:"hashed_password"`
	Active         bool   `json:"active"`
}

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	ReleaseYear int    `json:"release_year"`
	PageCount   int    `json:"page_count"`
	Cover       string `json:"cover"`
	Series      string `json:"series"`
	ReadStatus  string `json:"read_status"`
	Rating      int    `json:"rating"`
}

type ListContent struct {
	TypeID int `json:"type_id"`
	BookID int `json:"book_id"`
}

type ListType struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}
