package main

import (
	"fmt"
)

func countPrimes(n int) int {
	flagArray := make([]bool, n)
	result := 0
	for i := 2; i < n; i++ {
		if flagArray[i] == false {
			result++

			for j := 2; i*j < n; j++ {
				flagArray[i*j] = true
			}
		}
	}
	return result
}

func main() {
	fmt.Println(countPrimes(10))
}
