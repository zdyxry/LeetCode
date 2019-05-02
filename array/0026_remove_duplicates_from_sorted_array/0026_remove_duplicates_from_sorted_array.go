package main

import (
	"fmt"
)

func removeDuplicates(a []int) int {
	left, right, size := 0, 1, len(a)
	for ; right < size; right++ {
		if a[left] == a[right] {
			continue
		}
		left++
		a[left] = a[right]

	}
	return left + 1
}

func main() {
	a := []int{1, 1, 2}
	fmt.Println(removeDuplicates(a))
}
