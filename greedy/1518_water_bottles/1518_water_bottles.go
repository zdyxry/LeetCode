package main

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	s := 0
	left := 0
	for numBottles > 0 {
		s += numBottles
		left += numBottles
		numBottles = left / numExchange
		left %= numExchange
	}
	return s
}

func main() {
	numBottles := 9
	numExchange := 3
	res := numWaterBottles(numBottles, numExchange)
	fmt.Println(res)
}
