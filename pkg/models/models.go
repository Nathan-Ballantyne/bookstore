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
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Active         bool
}

type Book struct {
	ID          int
	Title       string
	Author      string
	ReleaseYear int
	PageCount   int
	Cover       string
	Series      string
	ReadStatus  string
	Rating      int
}

type List struct {
	ID     int
	BookID int
	UserID int
	Type   string
}
