package main

import "fmt"

func isPowerOfFour(num int) bool {
	if num < 1 {
		return false
	}
	for num%4 == 0 {
		num >>= 2
	}
	return num == 1
}

func main() {
	n := 5
	fmt.Println(isPowerOfFour(n))
	n = 8
	fmt.Println(isPowerOfFour(n))
	n = 64
	fmt.Println(isPowerOfFour(n))
}
