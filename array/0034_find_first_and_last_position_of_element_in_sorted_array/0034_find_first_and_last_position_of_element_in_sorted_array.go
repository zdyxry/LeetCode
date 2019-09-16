package main 

import "fmt"

func searchRange(nums []int, target int) []int {
	index := search(nums, target)
	if index == -1 {
		return []int{-1, -1}
	}

	first := index
	for {
		f := search(nums[:first], target)
		if f == -1{
			break
		}
		first = f
	}

	last := index
	for {
		l := search(nums[last+1:], target)
		if l == -1{
			break
		}
		last += l + 1
	}

	return []int{first, last}
}

func search(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	var median int

	for low < high {
		median = (low + high) / 2
		switch{
		case nums[median] < target :
			low = median + 1
		case nums[median] > target:
			high = median - 1
		default:
			return median
		}
	}

	return -1
}

func main() {
	nums := []int{5,7,7,8,8,10}
	target := 8

	res := searchRange(nums, target)
	fmt.Println(res)
}