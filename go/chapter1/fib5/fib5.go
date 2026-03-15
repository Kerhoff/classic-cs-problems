package main

import "fmt"

func fib5(n int) int {
	if n == 0 {
		return n
	}

	last := 0
	next := 1

	for i := 1; i < n; i++ {
		last, next = next, last+next
	}

	return next
}

func main() {
	fmt.Println(fib5(5))
	fmt.Println(fib5(50))
}
