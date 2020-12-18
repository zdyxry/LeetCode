package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := l + (r-l+1)/2
		if arr[mid-1] < arr[mid] {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {
	arr := []int{0, 1, 0}
	res := peakIndexInMountainArray(arr)
	fmt.Println(res)
}
