package main

import (
	"database/sql"
	"net/http"

	"github.com/awesomeProject/internal/http/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db, _ := sql.Open("mysql", "root:1@tcp(3306)")
	store := New(db)
	svc := New(store)
	handle := New(svc)

	r := mux.NewRouter()
	r.HandleFunc("/student", handlers.Post).Methods("POST")

	http.ListenAndServe("localhost:8080", r)

}
