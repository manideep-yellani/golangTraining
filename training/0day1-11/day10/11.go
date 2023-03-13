// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func f1(i chan int) {
	fmt.Println("f1")
	i <- 1
}
func f2(i chan int) {
	fmt.Println("f2") //DONT BE FOOLED, ONLINE YOU GET 67 5 AS OUTPUT AS WELL
	i <- 2
}
func f3(i chan int) {
	fmt.Println("f3") ////DONT BE FOOLED, ONLINE YOU GET DIFF ANS
	i <- 3
}

func main() {
	c := make(chan int, 1) // VVIP:IN ONLINE DIFFERENT ANS, DINT BE FOOLED
	go f1(c)
	go f2(c)
	go f3(c)

	a := <-c

	fmt.Println(a, "VVIP:DONT BE FOOLED,ONLINE DIFF ANS")

}
