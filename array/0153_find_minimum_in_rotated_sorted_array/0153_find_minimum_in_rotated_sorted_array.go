package main 

import "fmt"

func findMin(nums []int) int {
	low, high := 0, len(nums) -1
	for low < high {
		if nums[low] < nums[high] {
			return nums[low]
		}

		mid := low + (high - low) >> 1
		if nums[mid] >= nums[low]{
			low = mid +1
		} else {
			high = mid
		}
	}

	return nums[low]
}

func findMin2(nums []int) int {
	Len := len(nums)

	i := 1
	for i < Len && nums[i-1] < nums[i] {
		i++
	}

	return nums[i%Len]
}

func main() {
	nums :=[]int{3,4,5,1,2}
	res := findMin(nums)
	fmt.Println(res)
	res1 := findMin2(nums)
	fmt.Println(res1)
}