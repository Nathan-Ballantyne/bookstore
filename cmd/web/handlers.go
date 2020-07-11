package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

func (app *application) addBook(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
	w.Write([]byte("All books in the system"))
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
