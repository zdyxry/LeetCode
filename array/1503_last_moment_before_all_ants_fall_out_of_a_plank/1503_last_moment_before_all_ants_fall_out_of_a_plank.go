package main

import "fmt"

func getLastMoment(n int, left []int, right []int) int {
	max := -1
	for _, i := range left {
		if i > max {
			max = i
		}
	}
	for _, i := range right {
		if n-i > max { // 右边要走n-i
			max = n - i
		}
	}
	return max
}

func main() {
	n := 7
	left := []int{}
	right := []int{0, 1, 2, 3, 4, 5, 6, 7}
	res := getLastMoment(n, left, right)
	fmt.Println(res)
}
