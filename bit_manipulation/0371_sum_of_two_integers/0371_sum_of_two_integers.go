package main

import "fmt"

func getSum(a int, b int) int {
	return ((a & b) << 1) + (a ^ b)
}

func main() {
	a, b := 1, 2
	fmt.Println(getSum(a, b))
}
