package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	//"github.com/Nathan-Ballantyne/bookstore/pkg/forms"
	"github.com/Nathan-Ballantyne/bookstore/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
	w.Write([]byte("Home handler"))
}

func (app *application) allBooks(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
	w.Write([]byte("All books in the system"))
}

func (app *application) addBook(w http.ResponseWriter, r *http.Request) int {
	app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

	// Parse the form data.
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return 0
	}
	// Validate the form contents using the form helper we made earlier.
	form := r.PostForm
	// Try to create a new user record in the database. If the email already exists
	// add an error message to the form and re-display it.
	relYear, err := strconv.Atoi(form.Get("release_year"))
	pageCount, err := strconv.Atoi(form.Get("page_count"))
	rating, err := strconv.Atoi(form.Get("rating"))
	id, err := app.books.Insert(form.Get("title"), form.Get("author"), form.Get("cover"),
		form.Get("series"), form.Get("read_status"), relYear, pageCount, rating)

	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			fmt.Fprint(w, "An error occured")
		} else {
			app.serverError(w, err)
		}
		return 0
	}
	return id
}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		fmt.Fprintf(w, "Book with id %d not found", id)
		return
	}

	app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

	b, err := app.books.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	bookJSON, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(bookJSON))
}
