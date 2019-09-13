package main 

import (
	"fmt"
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res, delta := 0, math.MaxInt64

	for i := range nums{
		if i >0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i + 1 , len(nums) -1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < target:
				l++
				if delta > target -s {
					delta = target - s 
					res = s
				}
			case s > target:
				r--
				if delta > s - target {
					delta = s - target 
					res = s
				}
			default:
				return s
			}
		}
	}

	return res
}

func main() {
	nums := []int{-1,2,1,-4}
	target := 1 

	res := threeSumClosest(nums, target)
	fmt.Println(res)
}