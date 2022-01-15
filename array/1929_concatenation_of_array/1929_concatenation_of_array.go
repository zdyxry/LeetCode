package main

func getConcatenation(nums []int) []int {
	result := append(nums, nums...)
	return result

}
