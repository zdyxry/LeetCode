package main

import "fmt"

func runningSum(nums []int) []int {
	var sum int
	var res []int
	for _, n := range nums {
		sum += n
		res = append(res, sum)
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	res := runningSum(nums)
	fmt.Println(res)
}
