package main

import (
	"fmt"
	"net/http"
)

type login struct {
	name string `json:"name"`
	pass string `json:"pass"`
}

func middle(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		f(w, r)
	}
}

func f1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is f1 function")
}

func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is f2 function")
}

func main() {
	http.HandleFunc("/f1", middle(f1))
	http.HandleFunc("/f2", middle(f2))

	http.ListenAndServe(":8080", nil)
}
