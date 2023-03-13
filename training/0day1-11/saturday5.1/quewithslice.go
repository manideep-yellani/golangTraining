package main

import "fmt"

func pushq(a []int, b ...int) []int {

	for _, i := range b {
		a = append(a, i)
	}
	return a
}

func popq(a []int) []int {
	a = a[1:]
	return a

}

func main() {
	a := []int{}
	a = pushq(a, 3, 4, 23, 54, 6)
	a = popq(a)

	fmt.Println(a)
}
