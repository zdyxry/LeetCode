package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	n := len(nums)

	if k > n {
		k %= n
	}
	if k == 0 || k == n {
		return
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
func main() {
	a := []int{1, 1, 1, 2, 3}
	rotate(a, 1)
	fmt.Println(a)

}
