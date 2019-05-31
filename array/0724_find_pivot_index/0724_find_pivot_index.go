package main

import (
	"fmt"
)

func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	tsum := 0
	for i, num := range nums {
		if (sum - num) == 2*tsum {
			return i
		}
		tsum += num

	}
	return -1

}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}
