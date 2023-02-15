package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {

	http.HandleFunc("/profile1", http.HandlerFunc(ordinaryfunc))
	http.ListenAndServe("localhost:8081", nil)
}

func ordinaryfunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`"this is string",this is not a string`))
}

//  curl http://localhost:8081/debug/pprof/heap --output heap.tar.gz
// $ go tool pprof -web http://localhost:8081/debug/pprof/heap --output heap.tar.gz
