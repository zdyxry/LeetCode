package main 

import "fmt"

func getMinDistance(nums []int, target int, start int) int {
	var n= len(nums)
	for left,right := start,start;left>=0 || right<n;{
		if left>=0 && nums[left]==target{
			return start-left
		}
		if right<n && nums[right] == target{
			return right-start
		}
		left--
		right++
	}
	return 0
}

func main() {
	nums := []int{1,2,3,4,5}
	target := 5
	start := 3
	res := getMinDistance(nums, target, start)
	fmt.Println(res)
}