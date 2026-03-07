package main

import "fmt"

func fib1(n int) int {
	return fib1(n-1) + fib1(n-2)
}

func main() {
	fmt.Println(fib1(5))
}
