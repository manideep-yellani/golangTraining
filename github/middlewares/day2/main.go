package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Login struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type handlerStruct struct {
	inner http.Handler
}

var user Login = Login{
	Name: "mani",
	Pass: "abc",
}

func (authReceiver *handlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	fmt.Println(authHeader)
	authorisationCode := strings.Split(authHeader, " ")
	//fmt.Println(authorisationCode)
	//usr := strings.Split(authorisationCode[1], ":")
	namePass, _ := base64.StdEncoding.DecodeString(authorisationCode[1])
	fmt.Println(namePass)
	fmt.Println(string(namePass))
	namePassString := strings.Split(string(namePass), ":")
	fmt.Println(namePassString)
	fmt.Printf("%T", namePassString)

	if reflect.DeepEqual(string(name), user.Name) && reflect.DeepEqual(string(pass), user.Pass) {
		authReceiver.inner.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DoSomething(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {

	h := http.HandlerFunc(DoSomething)
	http.Handle("/auth", &handlerStruct{h})
	http.ListenAndServe(":8060", nil)
}
