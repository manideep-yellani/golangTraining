package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

func second() {
	fmt.Println("SECOND ROUTINE")
}

func main() {
	go hello()
	go second()
	for {
	}
	fmt.Println("main function")
}
