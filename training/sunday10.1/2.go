package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	fmt.Println("RUN IT ONLINE DIFF ANS Q")
	select {

	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)

	case s3 := <-output1:
		fmt.Println("f3f", s3)
	case s3 := <-output2:
		fmt.Println("f4f", s3)
	case s3 := <-output1:
		fmt.Println("f5f", s3)

	}
	fmt.Println("RUN IT ONLINE DIFF ANS Q")
}
