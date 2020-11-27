package main

import "fmt"

func waysToMakeFair(nums []int) int {
	result, odd, even := 0, 0, 0
	for i, n := range nums {
		if i&1 == 0 {
			even += n
		} else {
			odd += n
		}
	}
	fmt.Println(even, odd)

	newOdd, newEven := 0, 0
	for i, n := range nums {
		if i&1 == 0 {
			even -= n
		} else {
			odd -= n
		}

		if odd+newOdd == even+newEven {
			result++
		}
		if i&1 == 0 {
			newOdd += n
		} else {
			newEven += n
		}
		fmt.Println(i, n, even, odd, newEven, newOdd, result)
	}
	return result
}

func main() {
	nums := []int{2, 1, 6, 4}
	res := waysToMakeFair(nums)
	fmt.Println(res)
}
