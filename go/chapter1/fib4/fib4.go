package main

import "fmt"

var cache = map[int]int{
	0: 0,
	1: 1,
}

func fib4(n int) int {
	if value, exists := cache[n]; exists {
		return value
	}

	cache[n] = fib4(n-1) + fib4(n-2)
	return cache[n]
}

func main() {
	fmt.Println(fib4(5))
	fmt.Println(fib4(50))
}
