package main

import "fmt"

func findMinFibonacciNumbers(k int) int {
	nums := []int{1, 1}
	for i := 2; ; i++ {
		v := nums[i-1] + nums[i-2]
		if v > k {
			break
		}
		nums = append(nums, v)
	}

	count := 0
	for k > 0 {
		k -= nums[len(nums)-1]
		count++
		if k == 0 {
			break
		}
		i := len(nums) - 1
		for ; i >= 0 && k < nums[i]; i-- {
		}
		nums = nums[:i+1]
	}
	return count
}

func main() {
	k := 7
	res := findMinFibonacciNumbers(k)
	fmt.Println(res)
}
