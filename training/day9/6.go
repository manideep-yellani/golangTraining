package main

import "fmt"

func routine(ch chan int) {
	fmt.Println("routine")
	ch <- 100
}

func main() {
	ch := make(chan int)
	go routine(ch)
	fmt.Println("main function")

}
