package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	list := make([]int, len(nums))
	copy(list, nums)
	sort.Slice(list, func(i int, j int) bool {
		return list[i] > list[j]
	})
	dict := make(map[int]int)
	for i, v := range list {
		dict[v] = i + 1
	}
	res := make([]string, len(nums))
	for i := 0; i < len(res); i++ {
		if dict[nums[i]] == 1 {
			res[i] = "Gold Medal"
		} else if dict[nums[i]] == 2 {
			res[i] = "Silver Medal"
		} else if dict[nums[i]] == 3 {
			res[i] = "Bronze Medal"
		} else {
			res[i] = strconv.Itoa(dict[nums[i]])
		}
	}
	return res
}

func main() {
	nums := []int{5, 4, 3, 2, 1}
	res := findRelativeRanks(nums)
	fmt.Println(res)
}
