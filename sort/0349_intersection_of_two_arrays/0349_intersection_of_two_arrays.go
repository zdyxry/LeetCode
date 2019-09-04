package main 

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	var res []int
	for _, n := range nums1{
		m[n] = true
	}
	for _, n := range nums2{
		if m[n] {
			delete(m,n)
			res = append(res, n)
		}
	}

	return res
}

func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	res := intersection(nums1, nums2)
	fmt.Println(res)
}