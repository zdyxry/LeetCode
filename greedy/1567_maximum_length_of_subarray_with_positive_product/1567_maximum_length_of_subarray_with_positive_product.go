package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxLen(nums []int) int {
	odd, even, cnt, res := make([]int, 0), []int{-1}, 0, 0
	for i := range nums {
		if nums[i] == 0 {
			even = []int{i}
			odd = []int{}
			cnt = 0
		}
		if nums[i] < 0 {
			cnt++
			if len(odd) == 0 {
				odd = append(odd, i)
			}
		}
		if cnt%2 == 1 {
			res = max(res, i-odd[0])
		} else {
			res = max(res, i-even[0])
		}
	}
	return res
}

func main() {
	nums := []int{1, -2, -3, 4}
	res := getMaxLen(nums)
	fmt.Println(res)
}
