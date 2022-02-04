package main

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	ans := []int{}
	tmp := make([][]int, 3)
	for i := range tmp {
		tmp[i] = make([]int, 101)
	}
	for _, v := range nums1 {
		tmp[0][v] = 1
	}
	for _, v := range nums2 {
		tmp[1][v] = 1
	}
	for _, v := range nums3 {
		tmp[2][v] = 1
	}
	for i := 1; i < 101; i++ {
		if tmp[0][i]+tmp[1][i]+tmp[2][i] >= 2 {
			ans = append(ans, i)
		}
	}
	return ans
}
