package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
	for {
		fmt.Printf("jkk")
	}
}

func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Printf(" \n main function")
}
