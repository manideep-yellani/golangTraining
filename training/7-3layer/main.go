package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	handler "training/7-3layer/internal/http/anime"
	service "training/7-3layer/internal/services/anime"
	"training/7-3layer/internal/stores/anime"
)

type Anime struct {
	name string
}

const (
	username = "root"
	password = "password"
	dbname   = "assignment1"
	dbaddr   = "127.0.0.1:3306"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/assignment1?parseTime=true")
	defer db.Close()
	if err != nil {
		fmt.Println("error is:", err)
	}

	store := anime.New(db)
	svc := service.New(store)
	h := handler.New(svc)

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/anime", h.GetAll).Methods(http.MethodGet)

	http.ListenAndServe("localhost:8000", r)
}
