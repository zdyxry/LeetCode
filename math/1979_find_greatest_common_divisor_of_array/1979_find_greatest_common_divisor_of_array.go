package main

func findGCD(nums []int) int {
	sort.Ints(nums)
	min, max := nums[0], nums[len(nums)-1]

	result := 0
	for i := 1; i <= max; i++ {
		if min%i == 0 && max%i == 0 && i > result {
			result = i
		}
	}

	return result
}
