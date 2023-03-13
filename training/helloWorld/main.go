package main

import (
	"fmt"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	fmt.Println("HELLO WORLD")
	h := handler{}
	app := gofr.New()
	app.Server.ValidateHeaders = false
	app.REST("hello", &h)
	app.Start()
}

type handler struct {
}

func (h *handler) Index(ctx *gofr.Context) (interface{}, error) {
	return "hello-world", nil
}
