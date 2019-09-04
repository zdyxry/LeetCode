package main 

import "fmt"

func intersect(nums1, nums2 []int) []int {
	m := map[int]int{}
	var res []int 

	for _, n := range nums1{
		m[n]++
	}

	for _, n := range nums2{
		if m[n] > 0 {
			res = append(res, n)
			m[n]--
		}
	}

	return res
}

func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	res := intersect(nums1, nums2)
	fmt.Println(res)
}