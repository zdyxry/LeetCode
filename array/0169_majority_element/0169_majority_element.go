package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	x, t := nums[0], 1

	for i := 1; i < len(nums); i++ {
		switch {
		case x == nums[i]:
			t++
		case t > 0:
			t--
		default:
			x = nums[i]
			t = 1
		}
	}

	return x
}
func main() {
	a := []int{1, 1, 1, 2, 3}
	fmt.Println(majorityElement(a))

}
