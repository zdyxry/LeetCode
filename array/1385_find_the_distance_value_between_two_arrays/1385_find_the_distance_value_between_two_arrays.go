package main

import "fmt"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var (
		i, j, rst int
		length1   = len(arr1)
		length2   = len(arr2)
	)
	for i = 0; i < length1; i++ {
		fail := true
		for j = 0; j < length2; j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				fail = false
			}
		}
		if fail {
			rst++
		}
	}
	return rst
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	arr1 := []int{4, 5, 8}
	arr2 := []int{10, 9, 1, 8}
	d := 2
	res := findTheDistanceValue(arr1, arr2, d)
	fmt.Println(res)
}
