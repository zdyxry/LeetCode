package main

func countQuadruplets(nums []int) (ans int) {
	cnts := map[int]int{}
	for i := 1; i < len(nums)-2; i++ {
		for j := 0; j < i; j++ {
			cnts[nums[i]+nums[j]]++
		}
		for j := i + 2; j < len(nums); j++ {
			ans += cnts[nums[j]-nums[i+1]]
		}
	}
	return
}
