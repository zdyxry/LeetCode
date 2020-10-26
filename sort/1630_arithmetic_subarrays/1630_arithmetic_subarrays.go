package main

import (
	"fmt"
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {

	res := make([]bool, len(l))

	for i := range l {
		temp := append([]int{}, nums[l[i]:r[i]+1]...)
		sort.Ints(temp)
		flag := true
		for j := 1; j < len(temp)-1; j++ {
			if temp[j+1]-temp[j] != temp[j]-temp[j-1] {
				flag = false
				break
			}
		}
		if flag {
			res[i] = true
		}
	}

	return res
}

func main() {
	nums := []int{4, 6, 5, 9, 3, 7}
	l := []int{0, 0, 2}
	r := []int{2, 3, 5}
	res := checkArithmeticSubarrays(nums, l, r)
	fmt.Println(res)
}
