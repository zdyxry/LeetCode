package main

import "fmt"

func fib(num int) int {
	if num <= 1 {
		return num
	}
	return fib(num-1) + fib(num-2)
}

func main() {
	fmt.Println(fib(6))
}
