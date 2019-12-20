package main 

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	n := len(nums)
	vector := make([]int, n)
	taken := make([]bool, n)
	var ans [][]int

	makePermutation(0, n, nums, vector, taken, &ans)
	return ans
}

func makePermutation(cur, n int, nums, vector []int, taken []bool, ans *[][]int) {
	if cur == n {
		tmp := make([]int, n)
		copy(tmp, vector)
		*ans = append(*ans, tmp)
		return
	}

	used := make(map[int]bool, n - cur)
	for i := 0; i < n; i++ {
		if !taken[i] && !used[nums[i]] {
			used[nums[i]] = true

			taken[i] = true
			vector[cur] = nums[i]
			makePermutation(cur+1, n, nums, vector, taken, ans)
			taken[i] = false
		}
	}
}

func main() {
	nums := []int{1,1,2}
	res := permuteUnique(nums)
	fmt.Println(res)
}