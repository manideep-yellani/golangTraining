package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
}

/*
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

*/

//package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//)
//
//func main() {
//	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "Hello!,helo ra ungama ")
//		w.WriteHeader(200)
//		w.Write([]byte(`{}`))
//	})
//
//	fmt.Printf("Starting server at port 8080\n")
//	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
//		log.Fatal(err)
//	}
//}
