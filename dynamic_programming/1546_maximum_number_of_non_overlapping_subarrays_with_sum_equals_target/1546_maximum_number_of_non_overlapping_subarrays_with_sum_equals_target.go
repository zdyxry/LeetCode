package main 

import "fmt"

func maxNonOverlapping(nums []int, target int) int {
	count, sum := 0, 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		_, hasKey := m[sum-target]
		if hasKey {
			m = make(map[int]int)
			count++
			sum = 0
		}
		m[sum]++
	}
	return count
}

func main() {
	nums := []int{1,1,1,1,1,}
	res := maxNonOverlapping(nums, 2)
	fmt.Println(res)
}