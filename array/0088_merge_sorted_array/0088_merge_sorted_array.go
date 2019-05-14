package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] <= nums2[j] {
			nums1[i+j+1] = nums2[j]
			j--
		} else {
			nums1[i+j+1] = nums1[i]
			i--
		}

	}
	for j >= 0 {
		nums1[i+j+1] = nums2[j]
		j--
	}
	fmt.Println(i, j)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
