package main

import "fmt"

func main() {
	var m map[int]int
	var a chan int
	if a == nil {
		fmt.Println(a, len(a), cap(a), m, len(m))
		fmt.Printf("Type of a is %T \n", a)
		a = make(chan int)
		fmt.Println(len(a), cap(a))
		fmt.Printf("Type of a is %T", a)
	}
}
