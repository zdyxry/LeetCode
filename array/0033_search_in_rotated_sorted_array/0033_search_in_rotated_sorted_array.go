package main 

import (
	"fmt"
)

func search(nums []int, target int) int {
	rotated := indexOfMin(nums)
	// fmt.Println(rotated)
	size := len(nums)
	left, right := 0, size - 1

	for left <= right {
		mid := (left + right) / 2
		rotatedMid := (rotated+mid) % size
		switch {
		case nums[rotatedMid] < target:
			left = mid + 1
		case target < nums[rotatedMid]:
			right = mid -1
		default:
			return rotatedMid
		}
	}

	return -1
}

func indexOfMin(nums []int) int {
	size := len(nums)
	left, right := 0, size -1
	for left < right {
		mid := (left + right) /2
		// fmt.Println(nums[left], nums[mid], nums[right])
		if nums[right] < nums[mid]{
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	nums := []int{4,5,6,7,0,1,2}
	target := 0 
	res := search(nums, target)
	fmt.Println(res)
}