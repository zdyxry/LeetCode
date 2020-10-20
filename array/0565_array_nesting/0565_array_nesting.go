package main

import "fmt"

func arrayNesting(nums []int) int {
	result := 0
	visited := make([]bool, len(nums))
	for i := range nums {
		cur := i
		count := 0
		for !visited[cur] {
			count++
			visited[cur] = true
			cur = nums[cur]
		}
		result = max(result, count)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{5, 4, 0, 3, 1, 6, 2}
	res := arrayNesting(nums)
	fmt.Println(res)
}
