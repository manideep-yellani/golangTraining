package main

func fib(i int) int {
	if i == 0 || i == 1 {
		return i
	}
	return fib(i-1) + fib(i-2)
}
