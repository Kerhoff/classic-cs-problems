package main

import "fmt"

var memo = map[int]int{
	0: 0,
	1: 1,
}

func fib3(n int) int {
	if _, exists := memo[n]; !exists {
		memo[n] = fib3(n-2) + fib3(n-1)
	}

	return memo[n]
}

func main() {
	fmt.Println(fib3(5))
	fmt.Println(memo)
	fmt.Println(fib3(50))
	fmt.Println(memo)
}
