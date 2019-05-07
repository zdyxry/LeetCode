package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	max := 0
	temp := 0
	for i := 1; i < len(prices); i++ {
		temp += prices[i] - prices[i-1]
		if temp < 0 {
			temp = 0
		}
		if max < temp {
			max = temp
		}
	}

	return max
}

func main() {
	a := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(a))

}
