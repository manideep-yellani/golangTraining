package main

import "fmt"

func main() {

	type data1 struct {
		name1, name2 int
	}

	var data = struct {
		a, b, c int
		gg      string
	}{
		12, 13, 2, "hello",
	}

	data.a = 12

	aa := data1{}
	fmt.Print(data, aa)
}
