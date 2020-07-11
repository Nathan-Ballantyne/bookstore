package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/all", http.HandlerFunc(app.allBooks))
	mux.Get("/book/:id", http.HandlerFunc(app.getBook))
	mux.Post("/addbook", http.HandlerFunc(app.addBook))

	return mux
}
