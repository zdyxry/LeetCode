package main

func targetIndices(nums []int, target int) []int {
	result := []int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = append(result, i)
		}
		if nums[i] > target {
			break
		}
	}
	return result

}
