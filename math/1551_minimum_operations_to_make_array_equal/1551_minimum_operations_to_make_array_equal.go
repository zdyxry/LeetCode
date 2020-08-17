package main

import "fmt"

func minOperations(n int) int {
	return n * n / 4
}

func main() {
	n := 4
	res := minOperations(n)
	fmt.Println(res)
}
