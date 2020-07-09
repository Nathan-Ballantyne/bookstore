package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bmizerany/pat"
)

func main() {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(home))
	mux.Get("/all", http.HandlerFunc(allBooks))
	mux.Get("/book/:id", http.HandlerFunc(book))
	mux.Post("/addbook", http.HandlerFunc(addBook))
	http.ListenAndServe(":4000", mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home handler"))
}

func allBooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All books in the system"))
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All books in the system"))
}

func book(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		w.Write([]byte("Book with not found"))
		return
	}
	fmt.Fprintf(w, "Book with id: %d", id)
}
