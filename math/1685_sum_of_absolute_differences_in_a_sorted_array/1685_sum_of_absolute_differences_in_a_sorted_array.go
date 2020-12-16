package main

import "fmt"

func getSumAbsoluteDifferences(nums []int) []int {
	sum, n := 0, len(nums)
	for _, num := range nums {
		sum += num
	}
	res := make([]int, n)
	left, right := 0, sum
	for i, num := range nums {
		res[i] = right - num*(n-i) + num*i - left
		left += num
		right -= num
	}
	return res
}

func main() {
	nums := []int{2, 3, 5}
	res := getSumAbsoluteDifferences(nums)
	fmt.Println(res)
}
