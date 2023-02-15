package main

import (
	"fmt"
	"sort"
)

func compa(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)
	l := 0
	if len(a) < len(b) {
		l = len(a)
	} else {
		l = len(b)
	}
	i := 0
	for ; i < l; i++ {
		if a[i] != b[i] {
			break
		}
	}
	if i == l {
		return true
	}
	return false
}

func main() {
	p := []int{1, 7, 7, 7, 7, 2, 3, 9}
	q := []int{3, 2, 1}

	n := compa(p, q)
	fmt.Println(n, p, q)

}
