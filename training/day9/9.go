package main

import (
	"fmt"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
	fmt.Println(len(squareop), cap(squareop), "squareop")
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
	fmt.Println(len(cubeop), cap(cubeop), "cubeop")
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	fmt.Println(len(sqrch), cap(sqrch)) //print
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println(squares, cubes, sqrch, cubech, &sqrch, &cubech)
	fmt.Println("Final output", squares+cubes, len(sqrch), cap(sqrch))
}
