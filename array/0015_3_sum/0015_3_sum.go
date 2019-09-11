package main 

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums{
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i + 1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s < 0 :
				l++
			case s > 0:
				r--
			default:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l, r = next(nums, l , r)
			}
		}
	}

	return res
}

func next(nums []int, l, r int) (int, int){
	for l < r {
		switch {
		case nums[l] == nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l, r
		}
	}

	return l, r
}

func main() {
	nums := []int{-1,0,1,2,-1,-4}
	res := threeSum(nums)
	fmt.Println(res)
}