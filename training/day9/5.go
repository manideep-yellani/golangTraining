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

	fmt.Printf("\n")
	for {
		fmt.Printf("7")
	}
}
func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	for {
		fmt.Println("kk")
	}
	fmt.Printf(" \n main function")
}
