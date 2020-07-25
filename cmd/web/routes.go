package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {

	mux := pat.New()
	//mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/api/allBooks", http.HandlerFunc(app.allBooks))
	mux.Get("/api/book/:id", http.HandlerFunc(app.getBook))
	mux.Post("/api/addbook", http.HandlerFunc(app.addBook))
	mux.Post("/api/user/signup", http.HandlerFunc(app.signupUser))

	return mux
}
