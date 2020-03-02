package main

func smallerNumbersThanCurrent(nums []int) []int {
	var sli = []int{}
	count := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i == j {
				continue
			} else if nums[i] > nums[j] {
				count++
			}
		}
		sli = append(sli, count)
		count = 0
	}
	return sli
}
