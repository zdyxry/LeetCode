package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	for n > 1 {
		if n%2 == 1 {
			return false
		}
		n >>= 1
	}

	return true
}

func main() {
	n := 5
	fmt.Println(isPowerOfTwo(n))
	n = 8
	fmt.Println(isPowerOfTwo(n))
}
