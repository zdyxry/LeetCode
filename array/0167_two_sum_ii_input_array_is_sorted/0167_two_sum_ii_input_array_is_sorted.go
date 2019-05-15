package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {

	start, end := 0, len(numbers)-1
	for start != end {
		sum := numbers[start] + numbers[end]
		if sum > target {
			end -= 1
		} else if sum < target {
			start += 1
		} else {
			break
		}
	}
	return []int{start + 1, end + 1}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}
