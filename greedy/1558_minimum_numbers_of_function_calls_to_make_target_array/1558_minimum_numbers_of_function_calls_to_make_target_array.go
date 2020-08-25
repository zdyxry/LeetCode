package main

import "fmt"

func minOperations(nums []int) int {
	tot := 0
	maxCounter := 0
	for _, v := range nums {
		counter := 0
		for v > 1 {
			if v%2 == 1 {
				v--
				tot++
			}
			v /= 2
			counter++
		}

		if v == 1 {
			tot++
		}

		if maxCounter < counter {
			maxCounter = counter
		}
	}

	return tot + maxCounter
}

func main() {
	nums := []int{1, 5}
	res := minOperations(nums)
	fmt.Println(res)
}
