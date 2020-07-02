package main

import (
	"fmt"
	"sort"
)

const mod = int(1e9 + 7)

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	mem := make([]int, len(nums))
	mem[0] = 1
	for i := 1; i < len(mem); i++ {
		mem[i] = (mem[i-1] * 2) % mod
	}

	res := 0
	for i, j := 0, len(nums)-1; i <= j; {
		if nums[i]+nums[j] > target {
			j--
		} else {
			res += mem[j-i]
			res %= mod
			i++
		}
	}
	return res
}

func main() {
	nums := []int{3, 5, 6, 7}
	target := 9
	res := numSubseq(nums, target)
	fmt.Println(res)							
}
