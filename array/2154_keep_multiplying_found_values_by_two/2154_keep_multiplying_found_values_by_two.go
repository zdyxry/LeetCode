package main

func findFinalValue(nums []int, original int) int {
	has := map[int]bool{}
	for _, num := range nums {
		has[num] = true
	}
	for has[original] {
		original *= 2
	}
	return original
}
