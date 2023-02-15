package main

/*
import "fmt"

func per(a []int) {
	aa := []int{}
	m, mm := map[int]int{}, map[int]int{}
	for i := range a {
		mm[i] += 1
	}
	p(a, m, mm, 0, aa)
}
func p(a []int, m, mm map[int]int, n int, aa []int) {
	if n == len(a) {
		fmt.Println(aa)
	}
	for _, i := range a {
		if m[i] == 0 {
			m[i] += 1
			aa = append(aa, i)
			n += 1
			p(a, m, mm, n, aa)
			aa = aa[:len(aa)-1]
			n -= 1
			m[i] -= 1
		}
	}
}
func main() {

	a := []int{1, 2, 3, 4}
	per(a)

}
*/
