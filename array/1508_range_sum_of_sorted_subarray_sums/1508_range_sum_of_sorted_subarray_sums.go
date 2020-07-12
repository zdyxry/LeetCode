package main

import (
	"fmt"
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	l := []int{}
	for i, n := range nums {
		l = append(l, n)
		total := n
		for _, m := range nums[i+1:] {
			total += m
			l = append(l, total)
		}
	}
	sort.Ints(l)

	res := 0
	for _, n := range l[left-1 : right] {
		res += n
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	n := 4
	left := 1
	right := 5
	res := rangeSum(nums, n, left, right)
	fmt.Println(res)
}
