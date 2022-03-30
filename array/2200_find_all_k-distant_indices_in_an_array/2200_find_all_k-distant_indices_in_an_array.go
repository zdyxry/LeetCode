package main

func findKDistantIndices(nums []int, key int, k int) []int {
	res := []int{}
	for i := range nums {
		for j := range nums {
			if nums[j] == key && abs(i-j) <= k {
				res = append(res, i)
				break
			}
		}
	}
	return res

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
