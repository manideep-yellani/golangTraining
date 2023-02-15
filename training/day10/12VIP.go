// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"time"
)

func f1(c chan int) {

	go f2(c)
	fmt.Println(67, <-c)
}

func f2(c chan int) {
	c <- 5
}
func f3(c chan int) { //DONT BE FOOLED, ONLINE YOU GET 67 5 AS OUTPUT AS WELL

	//go f2(c)
	fmt.Println(68, <-c) ////DONT BE FOOLED, ONLINE YOU GET 67 5 AS OUTPUT AS WELL
}

func main() {
	//c := make(chan int)
	aa := make(chan int)
	go f1(aa)
	go f3(aa)
	time.Sleep(time.Second) //DONT BE FOOLED, ONLINE YOU GET 67 5 AS OUTPUT AS WELL
	//go f2(c)
	//a := <-c
	//fmt.Println(a)

}
