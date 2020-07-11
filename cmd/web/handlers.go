package main

import (
	"fmt"
	"net/http"
	"strconv"
)

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
