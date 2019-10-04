package main 

import "fmt"

func search(nums []int, target int) bool {
	length := len(nums)

	if length == 0 {
		return false
	}

	k:= 1
	for k < len(nums) && nums[k-1] <= nums[k] {
		k++
	}
	fmt.Println("k is", k)
	i, j := 0, length -1
	for i <= j{
		m := (i + j)/ 2
		med := (m+k) % length
		fmt.Println("m is", m, "med is", med)

		switch {
		case nums[med] < target:
			i = m + 1
		case target < nums[med]:
			j = m - 1
		default:
			return true
		}
	}

	return false
}

func main() {
	nums := []int{2,5,6,0,0,1,2}
	target := 0
	res := search(nums, target)
	fmt.Println(res)
}