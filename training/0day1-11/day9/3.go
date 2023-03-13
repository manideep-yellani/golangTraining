package main

import (
	"fmt"
)

func hello(i int) {
	fmt.Println(i)
}
func main() {
	for i := 0; ; {
		go hello(i)
		i++
	}

}
