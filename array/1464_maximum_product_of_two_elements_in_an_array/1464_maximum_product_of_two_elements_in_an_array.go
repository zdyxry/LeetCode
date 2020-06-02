package main

import "fmt"

func maxProduct(nums []int) int {
	max := 0
	secondMax := 0
	for i := 0; i < len(nums); i++ {
		e := nums[i]
		if e > max {
			secondMax = max
			max = e
		} else if e > secondMax {
			secondMax = e
		}
	}
	return (max - 1) * (secondMax - 1)
}

func main() {
	nums := []int{3, 4, 5, 2}
	res := maxProduct(nums)
	fmt.Println(res)
}
