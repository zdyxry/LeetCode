package main 

import "fmt"

func productExceptSelf(nums []int) []int {
	size := len(nums)

	res := make([]int, size)

	left := 1
	for i := 0; i < size; i++ {
		res[i] = left
		left *= nums[i]
	}

	right := 1
	for i := size - 1; i>= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}

func main() {
	nums := []int{1,2,3,4}
	res := productExceptSelf(nums)
	fmt.Println(res)
}