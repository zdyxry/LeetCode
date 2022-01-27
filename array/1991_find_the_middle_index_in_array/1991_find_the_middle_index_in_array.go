package main

func findMiddleIndex(nums []int) int {
	for idx, _ := range nums {
		if sum(nums[:idx]) == sum(nums[idx+1:]) {
			return idx
		}
	}
	return -1

}

func sum(nums []int) (result int) {
	for _, n := range nums {
		result += n
	}
	return
}
